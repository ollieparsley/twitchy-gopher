package twitch

import (
	"fmt"
	"strconv"
)

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
