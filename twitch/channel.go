package twitch

import (
	"fmt"
	"time"
)

//GetChannelByIDInput the inputs used with the get channel by id endpoint
type GetChannelByIDInput struct {
	ChannelID int64
}

//UpdateChannelInput the inputs used with the update channel endpoint
type UpdateChannelInput struct {
	ChannelID          int64
	Status             string
	Game               string
	Delay              int64
	ChannelFeedEnabled bool
}

//GetChannelEditorsInput the inputs used with the get channel editors endpoint
type GetChannelEditorsInput struct {
	ChannelID int64
}

//GetChannelEditorsOutput the outputs used with the get channel editors endpoint
type GetChannelEditorsOutput struct {
	Users []User `json:"users"`
}

//GetChannelFollowersInput the inputs used with the get channel by id endpoint
type GetChannelFollowersInput struct {
	ChannelID int64
	Limit     int64
	Offset    int64
	Cursor    string
	Direction string
}

//GetChannelFollowersOutput the outputs used with the get channel editors endpoint
type GetChannelFollowersOutput struct {
	Cursor  string                            `json:"cursor"`
	Total   int64                             `json:"total"`
	Follows []GetChannelFollowersFollowOutput `json:"follows"`
}

//GetChannelFollowersFollowOutput the outputs used with the get channel editors endpoint
type GetChannelFollowersFollowOutput struct {
	CreatedAt     time.Time `json:"created_at"`
	Notifications bool      `json:"notifications"`
	User          User      `json:"user"`
}

// GetChannel - the channel details for the authenticated user
func (c *Client) GetChannel() (*Channel, *ErrorOutput) {
	output := new(Channel)
	errorOutput := c.sendAPIRequest("GET", "channel", nil, output)
	return output, errorOutput
}

// GetChannelByID - Get a single channel feed post
func (c *Client) GetChannelByID(input *GetChannelByIDInput) (*Channel, *ErrorOutput) {
	output := new(Channel)
	errorOutput := c.sendAPIRequest("GET", fmt.Sprintf("channels/%d", input.ChannelID), nil, output)
	return output, errorOutput
}

// UpdateChannel - Updates a channel metadata
func (c *Client) UpdateChannel(input *UpdateChannelInput) (*Channel, *ErrorOutput) {
	params := map[string]string{
		"status":               input.Status,
		"game":                 input.Game,
		"delay":                fmt.Sprintf("%d", input.Delay),
		"channel_feed_enabled": "true",
	}
	if input.ChannelFeedEnabled == false {
		params["channel_feed_enabled"] = "false"
	}
	output := new(Channel)
	errorOutput := c.sendAPIRequest("PUT", fmt.Sprintf("channel/%d", input.ChannelID), params, output)
	return output, errorOutput
}

// GetChannelEditors - Get a the editors for a channel
func (c *Client) GetChannelEditors(input *GetChannelEditorsInput) (*GetChannelEditorsOutput, *ErrorOutput) {
	output := new(GetChannelEditorsOutput)
	errorOutput := c.sendAPIRequest("GET", fmt.Sprintf("channels/%d/editors", input.ChannelID), nil, output)
	return output, errorOutput
}

// GetChannelFollowers - Get a the followers for a channel
func (c *Client) GetChannelFollowers(input *GetChannelFollowersInput) (*GetChannelFollowersOutput, *ErrorOutput) {
	params := map[string]string{
		"limit":     fmt.Sprintf("%d", input.Limit),
		"offset":    fmt.Sprintf("%d", input.Offset),
		"cursor":    input.Cursor,
		"direction": input.Direction,
	}
	output := new(GetChannelFollowersOutput)
	errorOutput := c.sendAPIRequest("GET", fmt.Sprintf("channels/%d/follows", input.ChannelID), params, output)
	return output, errorOutput
}
