package blazingdocsgo

import (
	"io"
	"net/http"

	"github.com/blazingdocs/blazingdocs-go/config"
)

type Client struct {
	Config     config.Config
	HttpClient http.Client
}

func (client *Client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", client.Config.BaseURL+url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)

}

func (client *Client) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", client.Config.BaseURL+url, body)
	if err != nil {
		return nil, err

	}

	req.Header.Set("Content-Type", contentType)

	return client.Do(req)

}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-API-KEY", client.Config.ApiKey)
	return client.HttpClient.Do(req)
}
