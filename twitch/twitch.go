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
	apiURL      string
	version     int
	httpClient  *http.Client
	oauthConfig *OAuthConfig
}

//NewClient a nice way of creating a new Client
func NewClient(oauthConfig *OAuthConfig, httpClient *http.Client) *Client {
	apiURL := "https://api.twitch.tv/kraken/"

	return &Client{
		apiURL:      apiURL,
		version:     5,
		httpClient:  httpClient,
		oauthConfig: oauthConfig,
	}
}

func (c *Client) createRequest(method string, path string, params map[string]string) *http.Request {

	// Create new http request and set it all up!
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

	// Set the user-agent
	req.Header.Add("User-Agent", "Twitchy Gopher (https://github.com/ollieparsley/twitchy-gopher")

	// Specify the API version
	req.Header.Add("Accept", fmt.Sprintf("application/vnd.twitchtv.v%d+json", c.version))

	// Authorization headers
	req.Header.Add("Authorization", "OAuth "+c.oauthConfig.AccessToken)
	req.Header.Add("Client-ID", c.oauthConfig.ClientID)

	// Add GET params
	if method == "GET" {
		for key, val := range params {
			req.URL.Query().Add(key, val)
		}
	}

	return req
}

func (c *Client) sendRequest(method string, path string, params map[string]string, output interface{}) *ErrorOutput {

	// Create the request
	req := c.createRequest(method, path, params)

	// Make the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errorToOutput(err)
	}

	// JSON decoding
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		decodeErr := json.NewDecoder(resp.Body).Decode(output)
		if decodeErr != nil {
			return errorToOutput(decodeErr)
		}
		return nil
	}

	errorOutput := &ErrorOutput{}
	json.NewDecoder(resp.Body).Decode(errorOutput)
	return errorOutput
}

func errorToOutput(err error) *ErrorOutput {
	return &ErrorOutput{
		Message: err.Error(),
		Error:   "Twitchy error",
		Status:  -1,
	}
}

// GetRoot - the base API request that is used to verify the users details
func (c *Client) GetRoot() (*RootOutput, *ErrorOutput) {
	output := new(RootOutput)
	errorOutput := c.sendRequest("GET", "", nil, output)
	return output, errorOutput
}

// GetBlocks - return a list of users from a users' block list
func (c *Client) GetBlocks(input *BlocksInput) (*BlocksOutput, *ErrorOutput) {
	params := map[string]string{}
	if input.Limit != 0 {
		params["limit"] = strconv.Itoa(input.Limit)
	}
	if input.Offset != 0 {
		params["offset"] = strconv.Itoa(input.Offset)
	}
	output := new(BlocksOutput)
	errorOutput := c.sendRequest("GET", fmt.Sprintf("users/%d/blocks", input.UserID), params, output)
	return output, errorOutput
}

// BlockUser - Block a user (target) on behalf of another user
func (c *Client) BlockUser(input *BlockUserInput) (*BlockOutput, *ErrorOutput) {
	output := new(BlockOutput)
	errorOutput := c.sendRequest("PUT", fmt.Sprintf("users/%d/blocks/%d", input.UserID, input.TargetUserID), nil, output)
	return output, errorOutput
}
