// generated by metactl sdk gen
package mql

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type HackernewsServer struct {
	opts HackernewsServerOpts
}

type HackernewsServerOpts struct {
	Service HackernewsService
}

func NewHackernewsServer(opts HackernewsServerOpts) http.Handler {
	return HackernewsServer{opts: opts}
}

func (s HackernewsServer) send(w http.ResponseWriter, rsp interface{}) (err error) {
	w.Header().Set(ContentTypeHeader, ContentTypeJson)
	w.Header().Set(AsgTypeHeader, reflect.TypeOf(rsp).Name())

	err = json.NewEncoder(w).Encode(rsp)
	if err != nil {
		return
	}

	return
}

func (s HackernewsServer) getService() Service {
	getPostFeedsEndpoint := s.opts.Service.GetGetPostFeedsEndpoint()
	getPostsEndpoint := s.opts.Service.GetGetPostsEndpoint()
	getSocialAccountsEndpoint := s.opts.Service.GetGetSocialAccountsEndpoint()

	return Service{
		Name:       String(s.opts.Service.Name()),
		SdkVersion: String(Version),
		Endpoints: &Endpoints{
			LookupService:     &LookupServiceEndpoint{},
			GetPostFeeds:      &getPostFeedsEndpoint,
			GetPosts:          &getPostsEndpoint,
			GetSocialAccounts: &getSocialAccountsEndpoint,
		},
	}
}

func (s HackernewsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get(AsgTypeHeader) {
	case LookupServiceRequestName:
		var req LookupServiceRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return
		}

		svc := s.getService()
		rsp := LookupServiceResponse{
			Output: &LookupServiceOutput{
				Service: &svc,
			},
		}

		err = s.send(w, rsp)
		if err != nil {
			return
		}
	case GetPostFeedsRequestName:
		var req GetPostFeedsRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return
		}

		rsp := s.opts.Service.GetPostFeeds(r.Context(), req)

		err = s.send(w, rsp)
		if err != nil {
			return
		}
	case GetPostsRequestName:
		var req GetPostsRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return
		}

		rsp := s.opts.Service.GetPosts(r.Context(), req)

		err = s.send(w, rsp)
		if err != nil {
			return
		}
	case GetSocialAccountsRequestName:
		var req GetSocialAccountsRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return
		}

		rsp := s.opts.Service.GetSocialAccounts(r.Context(), req)

		err = s.send(w, rsp)
		if err != nil {
			return
		}
	}
}
