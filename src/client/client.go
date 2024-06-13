package client

import (
	"net/http"
)

type Client struct {
	Client *http.Client
	Url    string
}

func New(client *http.Client, url string) *Client {
	return &Client{
		Client: client,
		Url:    url,
	}
}

func (ffc *Client) Do(req *http.Request) (*http.Response, error) {
	return ffc.Client.Do(req)
}

func (ffc *Client) ChunkGet(url string, headervalue string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("RANGE", headervalue)
	return ffc.Do(req)
}
