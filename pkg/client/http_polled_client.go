package client

import (
	"net/http"
	"time"
)

type PooledHTTPClient struct {
	httpClient  HTTPClient
	maxPoolSize int
	pool        chan struct{}
}

func NewPooledHTTPClient(maxPoolSize int) *PooledHTTPClient {
	return &PooledHTTPClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second, // TODO read from cfg or pass a param
		},
		pool:        make(chan struct{}, maxPoolSize),
		maxPoolSize: maxPoolSize,
	}
}

func (c *PooledHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if c.maxPoolSize > 0 {
		c.pool <- struct{}{}
		defer func() {
			<-c.pool
		}()
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	return resp, nil
}
