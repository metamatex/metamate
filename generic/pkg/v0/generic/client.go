package generic

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	client *http.Client
	f      Factory
	addr   string
}

func NewClient(f Factory, c *http.Client, addr string) Client {
	return Client{
		client: c,
		addr:   addr,
		f:      f,
	}
}

func (c Client) Send(gReq Generic) (gRsp Generic, err error) {
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(gReq.ToStringInterfaceMap())
	if err != nil {
		return
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.addr, b)
	if err != nil {
		return
	}
	httpReq.Header.Set(ContentTypeHeader, ContentTypeJson)
	httpReq.Header.Set(AsgTypeHeader, gReq.Type().Name())

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

func Send(f Factory, client *http.Client, addr string, gReq Generic) (gRsp Generic, err error) {
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(gReq.ToStringInterfaceMap())
	if err != nil {
		return
	}

	httpReq, err := http.NewRequest(http.MethodPost, addr, b)
	if err != nil {
		return
	}
	httpReq.Header.Set(ContentTypeHeader, ContentTypeJson)
	httpReq.Header.Set(AsgTypeHeader, gReq.Type().Name())

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
