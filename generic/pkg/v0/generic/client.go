package generic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

func Send(f Factory, client *http.Client, addr string, username string, password string, gReq Generic) (gRsp Generic, err error) {
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(gReq.ToStringInterfaceMap())
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, addr, b)
	if err != nil {
		return
	}
	req.Header.Set(ContentTypeHeader, ContentTypeJson)
	req.Header.Set(AsgTypeHeader, gReq.Type().Name())

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	rsp, err := client.Do(req)
	if err != nil {
		return
	}

	if rsp.StatusCode != 200 {
		var msg string
		if rsp.Status != "" {
			msg = rsp.Status
		} else {
			msg = fmt.Sprintf("status code: %v", rsp.StatusCode)
		}

		err = errors.New(msg)

		return
	}

	m := map[string]interface{}{}
	err = json.NewDecoder(rsp.Body).Decode(&m)
	if err != nil {
		return
	}

	return f.FromStringInterfaceMap(gReq.Type().Edges.Type.Response(), m)
}
