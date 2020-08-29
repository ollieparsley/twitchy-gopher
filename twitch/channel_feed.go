package twitch

import (
	"fmt"
)

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
