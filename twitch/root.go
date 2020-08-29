package twitch

import "time"

// RootOutput the output for the Root object
type RootOutput struct {
	Token RootToken `json:"token"`
}

//RootToken the token object in the root response
type RootToken struct {
	Authorization RootTokenAuthorization `json:"authorization"`
	Username      string                 `json:"user_name"`
	Valid         bool                   `json:"valid"`
	ClientID      string                 `json:"client_id"`
}

//RootTokenAuthorization The auth object within the root token section
type RootTokenAuthorization struct {
	Scopes    []string  `json:"scopes"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetRoot - the base API request that is used to verify the users details
func (c *Client) GetRoot() (*RootOutput, *ErrorOutput) {
	output := new(RootOutput)
	errorOutput := c.sendAPIRequest("GET", "", nil, output)
	return output, errorOutput
}
