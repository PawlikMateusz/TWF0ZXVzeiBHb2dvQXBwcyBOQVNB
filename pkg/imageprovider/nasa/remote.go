package nasa

import (
	"fmt"
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
	return []string{
		"www.testurl.pl",
		fmt.Sprint(fmt.Sprint(rp.apiKey)),
		fmt.Sprint(fmt.Sprint(rp.baseURL)),
	}, nil
}
