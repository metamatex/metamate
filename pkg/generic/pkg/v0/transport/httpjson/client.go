package httpjson

import (
	"bytes"
	"encoding/json"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"net/http"
)

type Client struct {
	client *http.Client
	token  string
	f      generic.Factory
	addr   string
}

func NewClient(f generic.Factory, c *http.Client, addr string, token string) (Client) {
	return Client{
		client: c,
		addr:   addr,
		f:      f,
		token:  token,
	}
}

func (c Client) Send(gReq generic.Generic) (gRsp generic.Generic, err error) {
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(gReq.ToStringInterfaceMap())
	if err != nil {
		return
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.addr, b)
	if err != nil {
		return
	}
	httpReq.Header.Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_JSON)
	httpReq.Header.Set(METAMATE_TYPE_HEADER, gReq.Type().Name())
	httpReq.Header.Set(AUTHORIZATION_HEADER, "Bearer "+c.token)

	res, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	m := map[string]interface{}{}
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return
	}

	return c.f.FromStringInterfaceMap(gReq.Type().Edges.Type.Response(), m)
}

func Send(f generic.Factory, client *http.Client, addr string, token string, gReq generic.Generic) (gRsp generic.Generic, err error) {
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(gReq.ToStringInterfaceMap())
	if err != nil {
		return
	}

	httpReq, err := http.NewRequest(http.MethodPost, addr, b)
	if err != nil {
		return
	}
	httpReq.Header.Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_JSON)
	httpReq.Header.Set(METAMATE_TYPE_HEADER, gReq.Type().Name())
	httpReq.Header.Set(AUTHORIZATION_HEADER, "Bearer "+token)

	res, err := client.Do(httpReq)
	if err != nil {
		return
	}

	m := map[string]interface{}{}
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return
	}

	return f.FromStringInterfaceMap(gReq.Type().Edges.Type.Response(), m)
}
