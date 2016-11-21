package twitch

import (
	"time"
)

// ErrorOutput - Twitch Error
type ErrorOutput struct {
	Error   string `json:"error"`
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//BlockUser the user secton within a block
type BlockUser struct {
	ID          int64     `json:"_id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Type        string    `json:"type"`
	Bio         string    `json:"bio"`
	Logo        string    `json:"logo"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

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
	User      BlockUser `json:"user"`
}

//UnblockUserInput the inputs used when deleting a block
type UnblockUserInput struct {
	UserID       int64
	TargetUserID int64
}

//UnblockUserOutput currently the output is empty
type UnblockUserOutput struct{}
