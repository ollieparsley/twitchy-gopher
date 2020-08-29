package twitch

import (
	"bytes"
	"fmt"
	"strings"
)

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
