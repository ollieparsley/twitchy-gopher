package twitch

import (
	"bytes"
	"fmt"
	"strings"
)

//Upload details of an upload
type Upload struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

//CreateVideoInput the parameters used to create a video - V4
type CreateVideoInput struct {
	ChannelName string
	Title       string
}

//CreateVideoOutput create the skeleton for a video upload
type CreateVideoOutput struct {
	Upload Upload `json:"upload"`
	Video  Video  `json:"video"`
}

//CompleteVideoInput the parameters used to complete a video upload
type CompleteVideoInput struct {
	VideoID string
	Token   string
}

//CompleteVideoOutput output from a completed video
type CompleteVideoOutput struct{}

//UploadVideoPartInput the parameters used to upload a video part
type UploadVideoPartInput struct {
	VideoID string
	Token   string
	Part    int
	Body    *bytes.Buffer
}

//UploadVideoPartOutput output from upoading a video part
type UploadVideoPartOutput struct {
	Upload Upload `json:"upload"`
	Video  Video  `json:"video"`
}

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
