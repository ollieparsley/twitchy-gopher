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

	// Specify the API version
	req.Header.Add("Accept", fmt.Sprintf("application/vnd.twitchtv.v%d+json", c.version))

	// Authorization headers
	req.Header.Add("Authorization", "OAuth "+c.oauthConfig.AccessToken)
	req.Header.Add("Client-ID", c.oauthConfig.ClientID)

	return req
}

func (c *Client) performRequest(req *http.Request, output interface{}) *ErrorOutput {
	// Make the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errorToOutput(err)
	}

	// JSON decoding
	code := resp.StatusCode
	if code == 204 {
		return nil
	} else if 200 <= code && code <= 299 {
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

func (c *Client) sendRequest(method string, path string, params map[string]string, output interface{}) *ErrorOutput {
	// Create the request
	req := c.createRequest(method, path, params)

	// Perform the request
	return c.performRequest(req, output)
}

func errorToOutput(err error) *ErrorOutput {
	return &ErrorOutput{
		Message: err.Error(),
		Error:   "Twitchy error",
		Status:  -1,
	}
}

//
// BLOCKS
//

// GetRoot - the base API request that is used to verify the users details
func (c *Client) GetRoot() (*RootOutput, *ErrorOutput) {
	output := new(RootOutput)
	errorOutput := c.sendRequest("GET", "", nil, output)
	return output, errorOutput
}

// ListBlocks - return a list of users from a users' block list
func (c *Client) ListBlocks(input *ListBlocksInput) (*ListBlocksOutput, *ErrorOutput) {
	params := map[string]string{}
	if input.Limit != 0 {
		params["limit"] = strconv.Itoa(input.Limit)
	}
	if input.Offset != 0 {
		params["offset"] = strconv.Itoa(input.Offset)
	}
	output := new(ListBlocksOutput)
	errorOutput := c.sendRequest("GET", fmt.Sprintf("users/%d/blocks", input.UserID), params, output)
	return output, errorOutput
}

// BlockUser - Block a user (target) on behalf of another user
func (c *Client) BlockUser(input *BlockUserInput) (*BlockUserOutput, *ErrorOutput) {
	output := new(BlockUserOutput)
	errorOutput := c.sendRequest("PUT", fmt.Sprintf("users/%d/blocks/%d", input.UserID, input.TargetUserID), nil, output)
	return output, errorOutput
}

// UnblockUser - Unblock a user (target) on behalf of another user
func (c *Client) UnblockUser(input *UnblockUserInput) (*UnblockUserOutput, *ErrorOutput) {
	output := new(UnblockUserOutput)
	errorOutput := c.sendRequest("DELETE", fmt.Sprintf("users/%d/blocks/%d", input.UserID, input.TargetUserID), nil, output)
	return output, errorOutput
}

//
// CHANNEL FEED
//

// ListChannelFeedPosts - List channel feed posts
func (c *Client) ListChannelFeedPosts(input *ListChannelFeedPostsInput) (*ListChannelFeedPostsOutput, *ErrorOutput) {
	output := new(ListChannelFeedPostsOutput)
	errorOutput := c.sendRequest("GET", fmt.Sprintf("feed/%d/posts", input.ChannelID), nil, output)
	return output, errorOutput
}

// CreateChannelFeedPost - create a post for a channel feed
func (c *Client) CreateChannelFeedPost(input *CreateChannelFeedPostInput) (*CreateChannelFeedPostOutput, *ErrorOutput) {
	params := map[string]string{}
	params["content"] = input.Content
	if input.Share == true {
		params["share"] = "true"
	} else {
		params["share"] = "false"
	}
	output := new(CreateChannelFeedPostOutput)
	errorOutput := c.sendRequest("POST", fmt.Sprintf("feed/%d/posts", input.ChannelID), params, output)
	return output, errorOutput
}

// GetChannelFeedPost - Get a single channel feed post
func (c *Client) GetChannelFeedPost(input *GetChannelFeedPostInput) (*GetChannelFeedPostOutput, *ErrorOutput) {
	output := new(GetChannelFeedPostOutput)
	errorOutput := c.sendRequest("GET", fmt.Sprintf("feed/%d/posts/%s", input.ChannelID, input.PostID), nil, output)
	return output, errorOutput
}

// DeleteChannelFeedPost - Delete a single channel feed post
func (c *Client) DeleteChannelFeedPost(input *DeleteChannelFeedPostInput) (*DeleteChannelFeedPostOutput, *ErrorOutput) {
	output := new(DeleteChannelFeedPostOutput)
	errorOutput := c.sendRequest("DELETE", fmt.Sprintf("feed/%d/posts/%s", input.ChannelID, input.PostID), nil, output)
	return output, errorOutput
}

// CreateChannelFeedPostReaction - create a reaction to a post on a channel feed
func (c *Client) CreateChannelFeedPostReaction(input *CreateChannelFeedPostReactionInput) (*CreateChannelFeedPostReactionOutput, *ErrorOutput) {
	params := map[string]string{}
	params["emote_id"] = input.EmoteID
	output := new(CreateChannelFeedPostReactionOutput)
	errorOutput := c.sendRequest("POST", fmt.Sprintf("feed/%d/posts/%s/reactions", input.ChannelID, input.PostID), params, output)
	return output, errorOutput
}

// DeleteChannelFeedPostReaction - Delete a single channel feed post reaction
func (c *Client) DeleteChannelFeedPostReaction(input *DeleteChannelFeedPostReactionInput) (*DeleteChannelFeedPostReactionOutput, *ErrorOutput) {
	params := map[string]string{}
	params["emote_id"] = input.EmoteID
	output := new(DeleteChannelFeedPostReactionOutput)
	errorOutput := c.sendRequest("DELETE", fmt.Sprintf("feed/%d/posts/%s/reactions", input.ChannelID, input.PostID), nil, output)
	return output, errorOutput
}
