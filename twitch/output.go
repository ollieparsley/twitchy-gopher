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

//
// ROOT
//

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

//User the user secton within a block
type User struct {
	ID          int64     `json:"_id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Type        string    `json:"type"`
	Bio         string    `json:"bio"`
	Logo        string    `json:"logo"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

//
// BLOCKS
//

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

//
// CHANNEL FEED POSTS
//

/*{
  "_total": 8,
  "_cursor": "1454101643075611000",
  "posts": [{
    "id": "20",
    "created_at": "2016-01-29T21:07:23.075611Z",
    "deleted": false,
    "emotes": [ ],
    "reactions": {
      "endorse": {
        "count": 2,
        "user_ids": [ ]
      }
    },
    "body": "Kappa post",
    "user": {
      "display_name": "bangbangalang",
      "_id": 104447238,
      "name": "bangbangalang",
      "type": "user",
      "bio": "i like turtles and cats",
      "created_at": "2015-10-15T19:52:17Z",
      "updated_at": "2016-01-29T21:06:42Z",
      "logo": null
    }
  }]
}*/

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
