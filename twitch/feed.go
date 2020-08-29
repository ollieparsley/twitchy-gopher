package twitch

import (
	"fmt"
	"time"
)

//ChannelFeedPost the input used with the List channel feed posts endpoint
type ChannelFeedPost struct {
	ID        string                 `json:"id"`
	CreatedAt time.Time              `json:"created_at"`
	Deleted   bool                   `json:"deleted"`
	Emotes    []interface{}          `json:"emotes"`
	Body      string                 `json:"body"`
	Reactions map[string]interface{} `json:"reactions"`
	User      User                   `json:"user"`
}

//ListChannelFeedPostsInput the input used with the List channel feed posts endpoint
type ListChannelFeedPostsInput struct {
	ChannelID int64
	Limit     int
	Cursor    string
}

//ListChannelFeedPostsOutput the array of blocks
type ListChannelFeedPostsOutput struct {
	Total  int64  `json:"_total"`
	Cursor string `json:"_cursor"`
	Posts  []ChannelFeedPost
}

//CreateChannelFeedPostInput the details of the channel feed post to create
type CreateChannelFeedPostInput struct {
	ChannelID int64
	Content   string
	Share     bool
}

//CreateChannelFeedPostOutput the output of a channel feed post creation
type CreateChannelFeedPostOutput struct {
	Post  ChannelFeedPost `json:"post"`
	Tweet string          `json:"tweet"`
}

//GetChannelFeedPostInput the single channel feed post input
type GetChannelFeedPostInput struct {
	ChannelID int64
	PostID    string
}

//GetChannelFeedPostOutput a single channel feed post
type GetChannelFeedPostOutput struct {
	ChannelFeedPost
}

//DeleteChannelFeedPostInput the single channel feed post delete input
type DeleteChannelFeedPostInput struct {
	ChannelID int64
	PostID    string
}

//DeleteChannelFeedPostOutput a single channel feed post deletion object
type DeleteChannelFeedPostOutput struct{}

//CreateChannelFeedPostReactionInput creating a reaction with an emote to a channel post
type CreateChannelFeedPostReactionInput struct {
	ChannelID int64
	PostID    string
	EmoteID   string
}

//CreateChannelFeedPostReactionOutput a reaction to a channel feed post
type CreateChannelFeedPostReactionOutput struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	EmoteID   string    `json:"emote_id"`
	User      User      `json:"user"`
}

//DeleteChannelFeedPostReactionInput the single channel feed post reaction delete input
type DeleteChannelFeedPostReactionInput struct {
	ChannelID int64
	PostID    string
	EmoteID   string
}

//DeleteChannelFeedPostReactionOutput a single channel feed post reaction deletion object
type DeleteChannelFeedPostReactionOutput struct{}

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
