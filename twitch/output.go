package twitch

import (
	"time"
)

//
// Twitch Error
//
type ErrorOutput struct {
	Error   string `json:"error"`
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

//
// Root
//
type RootOutput struct {
	Token RootToken `json:"token"`
}

type RootToken struct {
	Authorization RootTokenAuthorization `json:"authorization"`
	Username      string                 `json:"user_name"`
	Valid         bool                   `json:"valid"`
	ClientID      string                 `json:"client_id"`
}

type RootTokenAuthorization struct {
	Scopes    []string  `json:"scopes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//
// Blocks
//
type BlocksOutput struct {
	Blocks []BlocksBlock
}
type BlocksBlock struct {
	UpdatedAt time.Time       `json:"updated_at"`
	ID        int64           `json:"_id"`
	User      BlocksBlockUser `json:"user"`
}
type BlocksBlockUser struct {
	ID          int64     `json:"_id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Type        string    `json:"type"`
	Bio         string    `json:"bio"`
	Logo        string    `json:"logo"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
type BlocksInput struct {
	UserID int64 // The user ID for user whos block list you want to return
	Limit  int   // Maximum number of objects in array. Default is 25. Maximum is 100.
	Offset int   // Object offset for pagination. Default is 0.
}
