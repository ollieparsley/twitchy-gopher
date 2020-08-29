package twitch

// GetRoot - the base API request that is used to verify the users details
func (c *Client) GetRoot() (*RootOutput, *ErrorOutput) {
	output := new(RootOutput)
	errorOutput := c.sendAPIRequest("GET", "", nil, output)
	return output, errorOutput
}
