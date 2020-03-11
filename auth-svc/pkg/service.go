package pkg

import (
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/transport"
	"github.com/metamatex/metamate/gen/v0/sdk/utils/ptr"
	"github.com/square/go-jose"
)

type Service struct {
	signer jose.Signer
	opts   ServiceOpts
}

type ServiceOpts struct {
	Client     transport.Client
	Salt       string
	PrivateKey *rsa.PrivateKey
}

func NewService(opts ServiceOpts) (svc Service, err error) {
	svc.opts = opts

	svc.signer, err = jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: svc.opts.PrivateKey}, &jose.SignerOptions{})
	if err != nil {
		return
	}

	return
}

func (s Service) Name() (string) {
	return "auth-svc"
}

func (s Service) GetPipeClientAccountsEndpoint() (sdk.PipeClientAccountsEndpoint) {
	return sdk.PipeClientAccountsEndpoint{
		Filter: &sdk.PipeClientAccountsRequestFilter{
			Mode: &sdk.PipeModeFilter{
				Kind: &sdk.EnumFilter{
					Is: &sdk.PipeModeKind.Context,
				},
				Context: &sdk.ContextPipeModeFilter{
					Requester: &sdk.EnumFilter{
						In: []string{sdk.BusActor.Client},
					},
					Method: &sdk.EnumFilter{
						In: []string{sdk.Methods.Post},
					},
					Stage: &sdk.EnumFilter{
						In: []string{sdk.RequestStage.Request},
					},
				},
			},
		},
	}
}

func (s Service) PipeClientAccounts(ctx context.Context, req sdk.PipeClientAccountsRequest) (rsp sdk.PipeClientAccountsResponse) {
	accounts := req.Context.Post.ClientRequest.ClientAccounts

	for i, _ := range accounts {
		hashedPassword := hashPassword(s.opts.Salt, *accounts[i].Password.Value)

		accounts[i].Password = &sdk.Password{
			IsHashed:     ptr.Bool(true),
			HashFunction: &sdk.HashFunction.Sha256,
			Value:        &hashedPassword,
		}
	}

	req.Context.Post.ClientRequest.ClientAccounts = accounts

	rsp.Context = req.Context

	return
}

func hashPassword(salt, password string) (string) {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(salt+password)))
}

func (s Service) GetVerifyTokenEndpoint() (sdk.VerifyTokenEndpoint) {
	return sdk.VerifyTokenEndpoint{}
}

func (s Service) VerifyToken(ctx context.Context, req sdk.VerifyTokenRequest) (rsp sdk.VerifyTokenResponse) {
	err := func() (err error) {
		object, err := jose.ParseSigned(*req.Input.Token.Value)
		if err != nil {
			return
		}

		output, err := object.Verify(&s.opts.PrivateKey.PublicKey)
		if err != nil {
			return
		}

		jwtPayload := JwtPayload{}
		err = json.Unmarshal(output, &jwtPayload)
		if err != nil {
			return
		}

		rsp.Output = &sdk.VerifyTokenOutput{
			IsValid:         ptr.Bool(true),
			ClientAccountId: &jwtPayload.ClientAccountId,
		}

		return
	}()
	if err != nil {
		return
	}

	return
}

type JwtPayload struct {
	ClientAccountId sdk.ServiceId
}

func (s Service) GetAuthenticateClientAccountEndpoint() (sdk.AuthenticateClientAccountEndpoint) {
	return sdk.AuthenticateClientAccountEndpoint{}
}

func (s Service) AuthenticateClientAccount(ctx context.Context, req sdk.AuthenticateClientAccountRequest) (rsp sdk.AuthenticateClientAccountResponse) {
	err := func() (err error) {
		getReq := sdk.GetClientAccountsRequest{
			Mode: &sdk.GetMode{
				Kind: &sdk.GetModeKind.Id,
				Id:   req.Input.Id,
			},
			Select: &sdk.GetClientAccountsResponseSelect{
				// todo response meta select
				ClientAccounts: &sdk.ClientAccountSelect{
					Id: &sdk.ServiceIdSelect{
						Value:       ptr.Bool(true),
						ServiceName: ptr.Bool(true),
					},
					Password: &sdk.PasswordSelect{
						Value: ptr.Bool(true),
					},
				},
			},
		}

		getRsp, err := s.opts.Client.GetClientAccounts(ctx, getReq)
		if err != nil {
			return
		}

		if len(getRsp.ClientAccounts) == 0 {
			err = errors.New("storage service did not return a ClientAccount")

			return
		}

		if len(getRsp.ClientAccounts) != 1 {
			err = errors.New("storage service did return more than one ClientAccount")

			return
		}

		clientAccount := getRsp.ClientAccounts[0]

		if clientAccount.Password == nil ||
			clientAccount.Password.Value == nil {
			err = errors.New("clientAccount.Password.Value not set")

			return
		}

		hashed := hashPassword(s.opts.Salt, *req.Input.Password)

		if hashed != *clientAccount.Password.Value {
			err = errors.New("hashed != *clientAccount.Password.Value")

			return
		}

		jwtPayload := JwtPayload{
			ClientAccountId: *clientAccount.Id,
		}

		b, err := json.Marshal(jwtPayload)
		if err != nil {
			return
		}

		object, err := s.signer.Sign(b)
		if err != nil {
			return
		}

		s, err := object.CompactSerialize()
		if err != nil {
			return
		}

		rsp.Output = &sdk.AuthenticateClientAccountOutput{
			Token: &sdk.Token{
				Value: ptr.String(s),
			},
		}

		return
	}()
	if err != nil {
		rsp.Meta = &sdk.ResponseMeta{
			Errors: []sdk.Error{
				{
					Message: &sdk.Text{
						Formatting: &sdk.FormattingKind.Plain,
						Value:      ptr.String(err.Error()),
					},
				},
			},
		}
	}

	return
}
