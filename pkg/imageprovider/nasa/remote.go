package nasa

import (
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
	var wg sync.WaitGroup
	for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
		wg.Add(1)
		d := d
		go func(date time.Time) {
			defer wg.Done()

			// create request
			req, err := rp.createRequest(rp.apiKey, d)
			if err != nil {
				return
			}

			// send request
			resp, err := rp.httpClient.Do(req)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			// read url form response
			decoder := json.NewDecoder(resp.Body)
			var r Response
			err = decoder.Decode(&r)
			if err != nil {
				return
			}
			fmt.Printf("\nRESP CODE %s", resp.Status)
			urls = append(urls, r.URL)
		}(d)
	}
	wg.Wait()
	return urls, nil
}

func (rp *RemoteProvider) createRequest(api_key string, date time.Time) (*http.Request, error) {
	req, err := http.NewRequest("GET", rp.baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %s", err)
	}
	q := req.URL.Query()
	q.Add("api_key", rp.apiKey)
	q.Add("date", date.Format("2006-01-02"))
	req.URL.RawQuery = q.Encode()
	fmt.Printf("\nURL %s", req.URL.String())
	return req, nil
}
