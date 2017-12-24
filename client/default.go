package client

import (
	"errors"
	"net/http"
)

var (
	ClientError = errors.New("Http request is failed!")
)

type HTTPClient interface {
	Send()
}

type Client struct {
	*http.Client
}

func NewHttpClient() *Client {
	return &Client{
		&http.Client{},
	}
}

func (c *Client) Send(req *http.Request) (*http.Response, error) {
	res, err := c.Do(req)
	if err != err {
		panic(err)
	}

	return res, err
}
