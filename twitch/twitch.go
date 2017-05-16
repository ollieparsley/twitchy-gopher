package twitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
		return errorToOutput(err)
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
			return errorToOutput(decodeErr)
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

func errorToOutput(err error) *ErrorOutput {
	return &ErrorOutput{
		Message: err.Error(),
		Error:   "Twitchy error",
		Status:  -1,
	}
}

//
// UPLOAD
//

// CreateVideo - Create the skeleton for a video to upload the content to
func (c *Client) CreateVideo(input *CreateVideoInput) (*CreateVideoOutput, *ErrorOutput) {
	output := new(CreateVideoOutput)
	params := map[string]string{}
	params["channel_name"] = input.ChannelName
	//V5 params["channel_id"] = strconv.FormatInt(input.ChannelID, 10)
	params["title"] = input.Title
	errorOutput := c.sendAPIRequest("POST", "videos", params, output)
	return output, errorOutput
}

// UploadVideoPart - Upload a video part
func (c *Client) UploadVideoPart(input *UploadVideoPartInput) (*UploadVideoPartOutput, *ErrorOutput) {
	output := new(UploadVideoPartOutput)
	errorOutput := c.sendUploadRequest("PUT", fmt.Sprintf("upload/%s?upload_token=%s&part=%d", strings.Replace(input.VideoID, "v", "", 1), input.Token, input.Part), "", input.Body, output)
	return output, errorOutput
}

// CompleteVideo - Complete a video upload
func (c *Client) CompleteVideo(input *CompleteVideoInput) (*CompleteVideoOutput, *ErrorOutput) {
	output := new(CompleteVideoOutput)
	errorOutput := c.sendUploadRequest("POST", fmt.Sprintf("upload/%s/complete?upload_token=%s", strings.Replace(input.VideoID, "v", "", 1), input.Token), "application/x-www-form-urlencoded", bytes.NewBuffer([]byte{}), output)
	return output, errorOutput
}

//
// ROOT
//

// GetRoot - the base API request that is used to verify the users details
func (c *Client) GetRoot() (*RootOutput, *ErrorOutput) {
	output := new(RootOutput)
	errorOutput := c.sendAPIRequest("GET", "", nil, output)
	return output, errorOutput
}

//
// BLOCKS
//

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
	errorOutput := c.sendAPIRequest("GET", fmt.Sprintf("users/%d/blocks", input.UserID), params, output)
	return output, errorOutput
}

// BlockUser - Block a user (target) on behalf of another user
func (c *Client) BlockUser(input *BlockUserInput) (*BlockUserOutput, *ErrorOutput) {
	output := new(BlockUserOutput)
	errorOutput := c.sendAPIRequest("PUT", fmt.Sprintf("users/%d/blocks/%d", input.UserID, input.TargetUserID), nil, output)
	return output, errorOutput
}

// UnblockUser - Unblock a user (target) on behalf of another user
func (c *Client) UnblockUser(input *UnblockUserInput) (*UnblockUserOutput, *ErrorOutput) {
	output := new(UnblockUserOutput)
	errorOutput := c.sendAPIRequest("DELETE", fmt.Sprintf("users/%d/blocks/%d", input.UserID, input.TargetUserID), nil, output)
	return output, errorOutput
}

//
// CHANNEL FEED
//

// ListChannelFeedPosts - List channel feed posts
func (c *Client) ListChannelFeedPosts(input *ListChannelFeedPostsInput) (*ListChannelFeedPostsOutput, *ErrorOutput) {
	output := new(ListChannelFeedPostsOutput)
	errorOutput := c.sendAPIRequest("GET", fmt.Sprintf("feed/%d/posts", input.ChannelID), nil, output)
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
	errorOutput := c.sendAPIRequest("POST", fmt.Sprintf("feed/%d/posts", input.ChannelID), params, output)
	return output, errorOutput
}

// GetChannelFeedPost - Get a single channel feed post
func (c *Client) GetChannelFeedPost(input *GetChannelFeedPostInput) (*GetChannelFeedPostOutput, *ErrorOutput) {
	output := new(GetChannelFeedPostOutput)
	errorOutput := c.sendAPIRequest("GET", fmt.Sprintf("feed/%d/posts/%s", input.ChannelID, input.PostID), nil, output)
	return output, errorOutput
}

// DeleteChannelFeedPost - Delete a single channel feed post
func (c *Client) DeleteChannelFeedPost(input *DeleteChannelFeedPostInput) (*DeleteChannelFeedPostOutput, *ErrorOutput) {
	output := new(DeleteChannelFeedPostOutput)
	errorOutput := c.sendAPIRequest("DELETE", fmt.Sprintf("feed/%d/posts/%s", input.ChannelID, input.PostID), nil, output)
	return output, errorOutput
}

// CreateChannelFeedPostReaction - create a reaction to a post on a channel feed
func (c *Client) CreateChannelFeedPostReaction(input *CreateChannelFeedPostReactionInput) (*CreateChannelFeedPostReactionOutput, *ErrorOutput) {
	params := map[string]string{}
	params["emote_id"] = input.EmoteID
	output := new(CreateChannelFeedPostReactionOutput)
	errorOutput := c.sendAPIRequest("POST", fmt.Sprintf("feed/%d/posts/%s/reactions", input.ChannelID, input.PostID), params, output)
	return output, errorOutput
}

// DeleteChannelFeedPostReaction - Delete a single channel feed post reaction
func (c *Client) DeleteChannelFeedPostReaction(input *DeleteChannelFeedPostReactionInput) (*DeleteChannelFeedPostReactionOutput, *ErrorOutput) {
	params := map[string]string{}
	params["emote_id"] = input.EmoteID
	output := new(DeleteChannelFeedPostReactionOutput)
	errorOutput := c.sendAPIRequest("DELETE", fmt.Sprintf("feed/%d/posts/%s/reactions", input.ChannelID, input.PostID), nil, output)
	return output, errorOutput
}
