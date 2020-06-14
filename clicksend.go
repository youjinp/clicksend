package clicksend

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	clicksendURL = `https://rest.clicksend.com/v3`
)

type HttpClientAPI interface {
	Do(req *http.Request) (*http.Response, error)
}

type ClientAPI interface {
	doRequest(opts parameters, dst interface{}) error
	SendSMS(s *SMS) (SMSResponse, error)
}

// Client provides a connection to the Clicksend API
type Client struct {
	// HTTPClient
	HTTPClient HttpClientAPI
	// Username as described
	Username string
	// APIKey as described
	APIKey string
	// BaseURL is the root API endpoint
	BaseURL string
}

// an object to hold variable parameters to perform request.
type parameters struct {
	// Method is HTTP method type.
	Method string
	// Path is postfix for URI.
	Path string
	// Payload for the request.
	Payload interface{}
}

// NewClient builds a new Client pointer using the provided key and a default API base URL
// Accepts `httpClient`, `username`, `apiKey` as arguments
func NewClient(httpClient HttpClientAPI, username string, apiKey string) *Client {
	return &Client{
		HTTPClient: httpClient,
		Username:   username,
		APIKey:     apiKey,
		BaseURL:    clicksendURL,
	}
}

func (client *Client) doRequest(opts parameters, dst interface{}) error {
	url := fmt.Sprintf("%s/%s", client.BaseURL, opts.Path)

	req, err := http.NewRequest(opts.Method, url, nil)
	if err != nil {
		return err
	}

	if opts.Payload != nil {
		payloadData, err := json.Marshal(opts.Payload)
		if err != nil {
			return err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(payloadData))
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(client.Username+":"+client.APIKey)))

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, dst)
	return err
}
