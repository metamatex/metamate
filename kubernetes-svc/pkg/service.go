package pkg

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/gen/v0/sdk"

	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	CacertPath = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	TokenPath  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	Endpoint   = "https://kubernetes.default.svc/api/v1/namespaces/%v/services%v"
)

type Service struct {
	client    *http.Client
	token     string
	namespace string
}

func (Service) Name() (string) {
	return "kubernetes"
}

func NewService() (svc Service, err error) {
	caCert, err := ioutil.ReadFile(CacertPath)
	if err != nil {
		return
	}

	token, err := ioutil.ReadFile(TokenPath)
	if err != nil {
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	svc = Service{client: c, token: string(token), namespace: "default"}

	return
}

func (Service) GetGetServicesEndpoint() (sdk.GetServicesEndpoint) {
	return sdk.GetServicesEndpoint{
		Filter: &sdk.GetServicesRequestFilter{
			Or: []sdk.GetServicesRequestFilter{
				{
					Mode: &sdk.GetModeFilter{
						Kind: &sdk.EnumFilter{
							Is: &sdk.GetModeKind.Collection,
						},
					},
				},
				{
					Mode: &sdk.GetModeFilter{
						Kind: &sdk.EnumFilter{
							Is: sdk.String(sdk.GetModeKind.Id),
						},
						Id: &sdk.IdFilter{
							Kind: &sdk.EnumFilter{
								Is: &sdk.IdKind.ServiceId,
							},
						},
					},
				},
			},
		},
	}
}

func (s Service) GetServices(ctx context.Context, req sdk.GetServicesRequest) (rsp sdk.GetServicesResponse) {
	var svcs []sdk.Service
	var errs []error
	switch *req.Mode.Kind {
	case sdk.GetModeKind.Id:
		switch *req.Mode.Id.Kind {
		case sdk.IdKind.ServiceId:
			svcs, errs = s.GetServicesModeId(*req.Mode.Id.ServiceId)
		}
	case sdk.GetModeKind.Collection:
		svcs, errs = s.GetServicesModeCollection(s.namespace)
	}

	rsp.Services = svcs

	if len(errs) != 0 {
		for _, err := range errs {
			rsp.Errors = append(rsp.Errors, sdk.Error{
				Message: sdk.String(err.Error()),
			})
		}
	}

	return
}

func (s Service) GetServicesModeId(serviceId sdk.ServiceId) (services []sdk.Service, errs []error) {
	namespace, name := resolveIdValue(*serviceId.Value)

	rq, err := http.NewRequest(http.MethodGet, getIdUrl(namespace, name), nil)
	if err != nil {
		return
	}

	rq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	rp, err := s.client.Do(rq)
	if err != nil {
		return
	}

	k8sSvc := K8sService{}
	err = json.NewDecoder(rp.Body).Decode(&k8sSvc)
	if err != nil {
		return
	}

	svc, err := svcFromK8sSvc(k8sSvc)
	if err != nil {
		return
	}

	services = append(services, svc)

	return
}

func (s Service) GetServicesModeCollection(namespace string) (svcs []sdk.Service, errs []error) {
	err := func() (err error) {
		rq, err := http.NewRequest(http.MethodGet, getCollectionUrl(namespace), nil)
		if err != nil {
			return
		}

		rq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

		rp, err := s.client.Do(rq)
		if err != nil {
			return
		}

		r := struct {
			Items []K8sService
		}{}
		err = json.NewDecoder(rp.Body).Decode(&r)
		if err != nil {
			return
		}

		for _, k8sSvc := range r.Items {
			if !containsSvc(k8sSvc) {
				continue
			}

			svc, err := svcFromK8sSvc(k8sSvc)
			if err != nil {
				errs = append(errs, err)
			}

			svcs = append(svcs, svc)
		}

		return
	}()
	if err != nil {
		errs = append(errs, err)
	}

	return
}

func getCollectionUrl(namespace string) (string) {
	return fmt.Sprintf(Endpoint, namespace, "")
}

func getIdUrl(namespace, name string) (string) {
	return fmt.Sprintf(Endpoint, namespace, name)
}

func genIdValue(namespace, name string) (string) {
	return namespace + "/" + name
}

func resolveIdValue(value string) (namespace, name string) {
	spl := strings.Split(value, "/")

	return spl[0], spl[1]
}

func containsSvc(k8sSvc K8sService) (bool) {
	return k8sSvc.Metadata.Annotations.Transport != "" || k8sSvc.Metadata.Annotations.Port != ""
}

func svcFromK8sSvc(k8sSvc K8sService) (svc sdk.Service, err error) {
	svc.Id = &sdk.ServiceId{}
	svc.Id.Value = sdk.String(genIdValue(k8sSvc.Metadata.Namespace, k8sSvc.Metadata.Name))

	svc.Url = &sdk.Url{}
	svc.Url.Value = sdk.String("http://" + k8sSvc.Metadata.Name)

	svc.Transport = sdk.String(k8sSvc.Metadata.Annotations.Transport)

	i, err := strconv.ParseInt(k8sSvc.Metadata.Annotations.Port, 10, 32)
	if err != nil {
		err = errors.New(fmt.Sprintf("error parsing port for service %v: %v", *svc.Id.Value, err.Error()))

		return
	}

	svc.Port = sdk.Int32(int32(i))

	return
}

type K8sService struct {
	Metadata struct {
		Name        string
		Namespace   string
		Annotations struct {
			Transport string `json:"metamate.io/v0.service.transport"`
			Port      string `json:"metamate.io/v0.service.port"`
		}
	}
}
