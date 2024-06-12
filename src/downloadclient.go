package src

import (
	"net/http"
)

type FfDownloadClient struct {
	Client *http.Client
	Url    string
}

func New(client *http.Client, url string) *FfDownloadClient {
	return &FfDownloadClient{
		Client: client,
		Url:    url,
	}
}

func (ffc *FfDownloadClient) Do(req *http.Request) (*http.Response, error) {
	return ffc.Client.Do(req)
}

func (ffc *FfDownloadClient) ChunkGet(url string, headervalue string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("RANGE", headervalue)
	return ffc.Do(req)
}
