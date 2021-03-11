package nasa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/PawlikMateusz/TWF0ZXVzeiBHb2dvQXBwcyBOQVNB/pkg/client"
)

const (
	nasaAPIBaseURL = "https://api.nasa.gov/planetary/apod"
)

type RemoteProvider struct {
	httpClient client.HTTPClient
	baseURL    string
	apiKey     string
}

type Response struct {
	URL string `json:"url"`
}

func NewRemoteProvider(apiKey string, maxConcurrentRequests int) *RemoteProvider {
	httpClient := client.NewPooledHTTPClient(maxConcurrentRequests)
	return &RemoteProvider{
		baseURL:    nasaAPIBaseURL,
		httpClient: httpClient,
		apiKey:     apiKey,
	}
}

func (rp *RemoteProvider) GetImagesURLs(startDate, endDate time.Time) (urls []string, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	errs := make(chan error, 1)

	// iterate thought all days
	for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
		wg.Add(1)
		d := d
		go func(date time.Time) {
			defer wg.Done()

			// create request
			req, err := rp.createRequest(ctx, rp.apiKey, d)
			if err != nil {
				sendAsyncError(errs, fmt.Errorf("Failed to create request: %s", err))
				cancel()
				return
			}

			// send request
			resp, err := rp.httpClient.Do(req)
			if err != nil {
				sendAsyncError(errs, fmt.Errorf("error when making request to external API: %s", err))
				cancel()
				return
			}
			if resp.StatusCode != http.StatusOK {
				sendAsyncError(errs, fmt.Errorf("failed to get response from external API"))
				cancel()
				return
			}
			defer resp.Body.Close()

			// read url form response
			decoder := json.NewDecoder(resp.Body)
			var r Response
			err = decoder.Decode(&r)
			if err != nil {
				sendAsyncError(errs, fmt.Errorf("failed to decode external API response"))
				cancel()
				return
			}
			urls = append(urls, r.URL)
		}(d)
	}
	wg.Wait()
	if ctx.Err() != nil {
		return nil, <-errs
	}
	return urls, nil
}

func (rp *RemoteProvider) createRequest(ctx context.Context, api_key string, date time.Time) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", rp.baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %s", err)
	}
	q := req.URL.Query()
	q.Add("api_key", rp.apiKey)
	q.Add("date", date.Format("2006-01-02"))
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func sendAsyncError(ch chan error, err error) {
	select {
	case ch <- err:
	default:
	}
}
