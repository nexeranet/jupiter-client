package client

import "net/http"

type transport struct {
	underlyingTransport http.RoundTripper
	key                 string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("x-api-key", t.key)
	return t.underlyingTransport.RoundTrip(req)
}

func NewTokenClient(server, key string) (*ClientWithResponses, error) {
	clientWithHeader := http.Client{Transport: &transport{underlyingTransport: http.DefaultTransport, key: key}}
	keyHeader := func(c *Client) error {
		c.Client = &clientWithHeader
		return nil
	}
	return NewClientWithResponses(server, keyHeader)
}

