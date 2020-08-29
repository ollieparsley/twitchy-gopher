package twitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

//OAuthConfig contains all the oauth 2 config
type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
}

//Client the object everything is based on
type Client struct {
	apiURL        string
	apiVersion    int
	uploadURL     string
	uploadVersion int
	httpClient    *http.Client
	oauthConfig   *OAuthConfig
}

//NewClient a nice way of creating a new Client
func NewClient(oauthConfig *OAuthConfig, httpClient *http.Client) *Client {
	apiURL := "https://api.twitch.tv/kraken/"
	uploadURL := "https://uploads.twitch.tv/"

	return &Client{
		apiURL:        apiURL,
		apiVersion:    5,
		uploadURL:     uploadURL,
		uploadVersion: 4,
		httpClient:    httpClient,
		oauthConfig:   oauthConfig,
	}
}

//createBaseRequest Create new http request and set it all up
func (c *Client) createBaseRequest(method string, path string, params map[string]string) *http.Request {
	fullURL := fmt.Sprintf("%s%s", c.apiURL, path)

	// buffer
	var req *http.Request
	if method != "GET" && params != nil {
		data := url.Values{}
		for key, val := range params {
			data.Add(key, val)
		}
		buffer := bytes.NewBufferString(data.Encode())
		req, _ = http.NewRequest(method, fullURL, buffer)
	} else {
		req, _ = http.NewRequest(method, fullURL, nil)
	}

	// Add GET params
	if method == "GET" && params != nil {
		q := req.URL.Query()
		for key, val := range params {
			q.Set(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	// Set the user-agent
	req.Header.Add("User-Agent", "Twitchy Gopher (https://github.com/ollieparsley/twitchy-gopher")

	return c.authorizeRequest(req)
}

//authorizeRequest add the auth headers for the request
func (c *Client) authorizeRequest(req *http.Request) *http.Request {

	// Authorization headers
	req.Header.Add("Authorization", "OAuth "+c.oauthConfig.AccessToken)
	req.Header.Add("Client-ID", c.oauthConfig.ClientID)

	return req
}

func (c *Client) performRequest(req *http.Request, output interface{}) *ErrorOutput {
	// Make the request
	//fmt.Printf("\nURL: %+v\n", req.URL)
	//fmt.Printf("\nHEADERS: %+v\n", req.Header)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return c.errorToOutput(err)
	}

	//buf := new(bytes.Buffer)
	//buf.ReadFrom(resp.Body)
	//bodyString := buf.String()
	//dump, err := httputil.DumpResponse(resp, true)
	//fmt.Printf("BODY: %+v", string(dump))

	// JSON decoding
	code := resp.StatusCode
	if code == 204 || resp.Header.Get("Content-Length") == "0" {
		return nil
	} else if 200 <= code && code <= 299 {
		decodeErr := json.NewDecoder(resp.Body).Decode(output)
		if decodeErr != nil {
			return c.errorToOutput(decodeErr)
		}
		return nil
	}

	errorOutput := &ErrorOutput{}
	json.NewDecoder(resp.Body).Decode(errorOutput)
	return errorOutput
}

func (c *Client) createAPIRequest(method string, path string, params map[string]string) *http.Request {
	req := c.createBaseRequest(method, path, params)
	req.Header.Set("Accept", fmt.Sprintf("application/vnd.twitchtv.v%d+json", c.apiVersion))
	return req
}

func (c *Client) sendAPIRequest(method string, path string, params map[string]string, output interface{}) *ErrorOutput {
	// Create API request
	req := c.createAPIRequest(method, path, params)

	// Specify the API version
	req.Header.Set("Accept", fmt.Sprintf("application/vnd.twitchtv.v%d+json", c.apiVersion))

	// Perform the request
	return c.performRequest(req, output)
}

func (c *Client) createUploadRequest(method string, path string, contentType string, body *bytes.Buffer) *http.Request {
	// Create the request
	req, _ := http.NewRequest(method, c.uploadURL+path, body)

	// Add the content type for the request
	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}
	req.Header.Add("Content-Length", strconv.Itoa(body.Len()))

	// Specify the API version
	req.Header.Set("Accept", fmt.Sprintf("application/vnd.twitchtv.v%d+json", c.uploadVersion))

	// Authorize the request
	req = c.authorizeRequest(req)
	return req
}

func (c *Client) sendUploadRequest(method string, path string, contentType string, body *bytes.Buffer, output interface{}) *ErrorOutput {
	// Create upload request
	req := c.createUploadRequest(method, path, contentType, body)

	// Perform the request
	return c.performRequest(req, output)
}

func (c *Client) errorToOutput(err error) *ErrorOutput {
	return &ErrorOutput{
		Message: err.Error(),
		Error:   "Twitchy error",
		Status:  -1,
	}
}
