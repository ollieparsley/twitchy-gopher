package twitch

import (
	"fmt"
	"strconv"
	"time"
)

//ListBlocksInput the input used with the Blocks endpoint
type ListBlocksInput struct {
	UserID int64 // The user ID for user whos block list you want to return
	Limit  int   // Maximum number of objects in array. Default is 25. Maximum is 100.
	Offset int   // Object offset for pagination. Default is 0.
}

//ListBlocksOutput the array of blocks
type ListBlocksOutput struct {
	Blocks []BlockUserOutput
}

//BlockUserInput the inputs used with the block user endpoint
type BlockUserInput struct {
	UserID       int64
	TargetUserID int64
}

//BlockUserOutput a single block
type BlockUserOutput struct {
	UpdatedAt time.Time `json:"updated_at"`
	ID        int64     `json:"_id"`
	User      User      `json:"user"`
}

//UnblockUserInput the inputs used when deleting a block
type UnblockUserInput struct {
	UserID       int64
	TargetUserID int64
}

//UnblockUserOutput currently the output is empty
type UnblockUserOutput struct{}

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
