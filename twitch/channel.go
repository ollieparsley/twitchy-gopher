package twitch

import (
	"fmt"
	"strings"
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

//GetChannelFollowersInput the inputs used with the get channel followers endpoint
type GetChannelFollowersInput struct {
	ChannelID int64
	Limit     int64
	Offset    int64
	Cursor    string
	Direction string
}

//GetChannelFollowersOutput the outputs used with the get channel editors endpoint
type GetChannelFollowersOutput struct {
	Cursor  string                            `json:"_cursor"`
	Total   int64                             `json:"_total"`
	Follows []GetChannelFollowersFollowOutput `json:"follows"`
}

//GetChannelFollowersFollowOutput the outputs used with the get channel editors endpoint
type GetChannelFollowersFollowOutput struct {
	CreatedAt     time.Time `json:"created_at"`
	Notifications bool      `json:"notifications"`
	User          User      `json:"user"`
}

//GetChannelTeamsInput the inputs used with the get channel editors endpoint
type GetChannelTeamsInput struct {
	ChannelID int64
}

//GetChannelTeamsOutput the outputs used with the get channel teams endpoint
type GetChannelTeamsOutput struct {
	Teams []Team `json:"teams"`
}

//GetChannelSubscribersInput the inputs used with the get channel by subscribers endpoint
type GetChannelSubscribersInput struct {
	ChannelID int64
	Limit     int64
	Offset    int64
	Cursor    string
	Direction string
}

//GetChannelSubscribersOutput the outputs used with the get channel subscribers endpoint
type GetChannelSubscribersOutput struct {
	Total         int64          `json:"_total"`
	Subscriptions []Subscription `json:"subscriptions"`
}

//CheckChannelSubscriptionByUserInput the inputs used with the check channel subscription endpoint
type CheckChannelSubscriptionByUserInput struct {
	ChannelID int64
	UserID    int64
}

//GetChannelVideosInput the inputs used with the get channel videos endpoint
type GetChannelVideosInput struct {
	ChannelID      int64
	Limit          int64
	Offset         int64
	BroadcastTypes []string
	Language       string
	Sort           string
}

//GetChannelVideosOutput the outputs used with the get channel videos endpoint
type GetChannelVideosOutput struct {
	Total  int64   `json:"_total"`
	Videos []Video `json:"videos"`
}

//StartChannelCommercialInput the outputs used with the get channel videos endpoint
type StartChannelCommercialInput struct {
	ChannelID int64
	Length    int64
}

//StartChannelCommercialOutput the outputs used with the start channel commercial endpoint
type StartChannelCommercialOutput struct {
	Length     int64  `json:"Length"`
	Message    string `json:"Message"`
	RetryAfter int64  `json:"RetryAfter"`
}

//ResetStreamKeyInput the inputs used with the reset stream key endpoint
type ResetStreamKeyInput struct {
	ChannelID int64
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

// GetChannelTeams - Get a the editors for a channel
func (c *Client) GetChannelTeams(input *GetChannelTeamsInput) (*GetChannelTeamsOutput, *ErrorOutput) {
	output := new(GetChannelTeamsOutput)
	errorOutput := c.sendAPIRequest("GET", fmt.Sprintf("channels/%d/teams", input.ChannelID), nil, output)
	return output, errorOutput
}

// GetChannelSubscribers - Get a the subscribers for a channel
func (c *Client) GetChannelSubscribers(input *GetChannelSubscribersInput) (*GetChannelSubscribersOutput, *ErrorOutput) {
	params := map[string]string{
		"limit":     fmt.Sprintf("%d", input.Limit),
		"offset":    fmt.Sprintf("%d", input.Offset),
		"direction": input.Direction,
	}
	output := new(GetChannelSubscribersOutput)
	errorOutput := c.sendAPIRequest("GET", fmt.Sprintf("channels/%d/subscriptions", input.ChannelID), params, output)
	return output, errorOutput
}

// CheckChannelSubscriptionByUser - Get a single subscription by user ID
func (c *Client) CheckChannelSubscriptionByUser(input *CheckChannelSubscriptionByUserInput) (*Subscription, *ErrorOutput) {
	output := new(Subscription)
	errorOutput := c.sendAPIRequest("GET", fmt.Sprintf("channels/%d/subscriptions/%d", input.ChannelID, input.UserID), nil, output)
	return output, errorOutput
}

// GetChannelVideos - Get a the videos for a channel
func (c *Client) GetChannelVideos(input *GetChannelVideosInput) (*GetChannelVideosOutput, *ErrorOutput) {
	params := map[string]string{
		"limit":          fmt.Sprintf("%d", input.Limit),
		"offset":         fmt.Sprintf("%d", input.Offset),
		"broadcast_type": strings.Join(input.BroadcastTypes, ","),
		"langage":        input.Language,
		"sort":           input.Sort,
	}
	output := new(GetChannelVideosOutput)
	errorOutput := c.sendAPIRequest("GET", fmt.Sprintf("channels/%d/videos", input.ChannelID), params, output)
	return output, errorOutput
}

// StartChannelCommercial - Start a commercial for a channel
func (c *Client) StartChannelCommercial(input *StartChannelCommercialInput) (*StartChannelCommercialOutput, *ErrorOutput) {
	params := map[string]string{
		"length": fmt.Sprintf("%d", input.Length),
	}
	output := new(StartChannelCommercialOutput)
	errorOutput := c.sendAPIRequest("POST", fmt.Sprintf("channel/%d/commercial", input.ChannelID), params, output)
	return output, errorOutput
}

// ResetStreamKey - Reset the stream key for a channel
func (c *Client) ResetStreamKey(input *ResetStreamKeyInput) (*Channel, *ErrorOutput) {
	output := new(Channel)
	errorOutput := c.sendAPIRequest("DELETE", fmt.Sprintf("channels/%d/stream_key", input.ChannelID), nil, output)
	return output, errorOutput
}
