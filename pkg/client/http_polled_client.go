package client

import "net/http"

type PooledHTTPClient struct {
	httpClient  HTTPClient
	maxPoolSize int
	pool        chan int
}

func NewPooledHTTPClient(maxPoolSize int) *PooledHTTPClient {
	return &PooledHTTPClient{
		httpClient: &http.Client{
			Timeout: 10, // TODO read from cfg or pass a param
		},
		pool:        make(chan int, maxPoolSize),
		maxPoolSize: maxPoolSize,
	}
}

func (c *PooledHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if c.maxPoolSize > 0 {
		c.pool <- -1
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
