package nasa

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PawlikMateusz/TWF0ZXVzeiBHb2dvQXBwcyBOQVNB/pkg/client"
)

type RemoteProvider struct {
	httpClient client.HTTPClient
	baseURL    string
	apiKey     string
}

func NewRemoteProvider(baseURL, apiKey string, maxConcurrentRequests int) *RemoteProvider {
	httpClient := client.NewPooledHTTPClient(maxConcurrentRequests)
	return &RemoteProvider{
		baseURL:    "www.google.com",
		httpClient: httpClient,
	}
}

func (rp *RemoteProvider) GetImagesURLs(startDate, endDate time.Time) (urls []string, err error) {
	req, err := http.NewRequest("GET", rp.baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %s", err)
	}
	q := req.URL.Query()
	q.Add("api_key", rp.apiKey)
	q.Add("date", "")
	for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {

		q.Set("date", d.Format("2006-01-02"))
		req.URL.RawQuery = q.Encode()
		fmt.Printf("\nMaking request %s", req.URL.String())
		// resp, err := rp.httpClient.Do(req)
	}

	return []string{
		"www.testurl.pl",
		fmt.Sprint(fmt.Sprint(rp.apiKey)),
		fmt.Sprint(fmt.Sprint(rp.baseURL)),
	}, nil
}
