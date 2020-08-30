package twitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
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

// ErrorOutput - Twitch Error
type ErrorOutput struct {
	Error   string `json:"error"`
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

//User the user secton within a block
type User struct {
	ID          int64     `json:"_id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Type        string    `json:"type"`
	Bio         string    `json:"bio"`
	Logo        string    `json:"logo"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

//Subscription the details of a subscription and the related user
type Subscription struct {
	ID          string    `json:"_id"`
	CreatedAt   time.Time `json:"created_at"`
	SubPlan     string    `json:"sub_plan"`
	SubPlanName string    `json:"sub_plan_name"`
	User        User      `json:"user"`
}

//Channel details of a channel
type Channel struct {
	ID                           string    `json:"_id"`
	Name                         string    `json:"name"`
	DisplayName                  string    `json:"display_name"`
	Mature                       bool      `json:"mature"`
	Status                       string    `json:"status"`
	BroadcasterLanguage          string    `json:"broadcaster_language"`
	Game                         string    `json:"game"`
	Language                     string    `json:"language"`
	CreatedAt                    time.Time `json:"created_at,omitempty"`
	UpdatedAt                    time.Time `json:"updated_at,omitempty"`
	Logo                         string    `json:"logo"`
	VideoBanner                  string    `json:"video_banner"`
	ProfileBanner                string    `json:"profile_banner"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color"`
	Partner                      bool      `json:"partner"`
	URL                          string    `json:"url"`
	Views                        int64     `json:"views"`
	Followers                    int64     `json:"followers"`
	BroadcasterTyoe              string    `json:"broadcaster_type"`
	StreamKey                    string    `json:"stream_key"`
	Email                        string    `json:"email"`
}

//Team the user secton within a block
type Team struct {
	ID          int64     `json:"_id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Banner      string    `json:"banner"`
	Background  string    `json:"background"`
	Info        string    `json:"Info"`
	Logo        string    `json:"logo"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

//Preview details of the different preview image sizes
type Preview struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Large    string `json:"large"`
	Template string `json:"template"`
}

//Thumbnails details of the different thumbnail sizes
type Thumbnails struct {
	Small    []interface{} `json:"small"`
	Medium   []interface{} `json:"medium"`
	Large    []interface{} `json:"large"`
	Template []interface{} `json:"template"`
}

//Resolutions details of the different resolutions
type Resolutions struct {
	Chunked string `json:"chunked"`
	High    string `json:"high"`
	Low     string `json:"low"`
	Medium  string `json:"medium"`
	Mobile  string `json:"mobile"`
}

//FPS details of the different frames per second
type FPS struct {
	Chunked float64 `json:"chunked"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Medium  float64 `json:"medium"`
	Mobile  float64 `json:"mobile"`
}

//Video details of a video
type Video struct {
	ID            string      `json:"_id"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	BroadcastID   int         `json:"broadcast_id"`
	BroadcastType string      `json:"broadcast_type"`
	Status        string      `json:"status"`
	TagList       string      `json:"tag_list"`
	Views         int64       `json:"views"`
	URL           string      `json:"url"`
	Language      string      `json:"language"`
	Viewable      string      `json:"viewable"`
	ViewableAt    *time.Time  `json:"viewable_at,omitempty"`
	RecordedAt    *time.Time  `json:"recorded_at,omitempty"`
	CreatedAt     *time.Time  `json:"created_at,omitempty"`
	Game          string      `json:"game"`
	Length        int64       `json:"length"`
	Preview       Preview     `json:"preview"`
	Thumbnails    Thumbnails  `json:"thumbnails"`
	Paywalled     bool        `json:"paywalled"`
	FPS           FPS         `json:"fps"`
	Resolutions   Resolutions `json:"resolutions"`
	Channel       Channel     `json:"channel"`
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
