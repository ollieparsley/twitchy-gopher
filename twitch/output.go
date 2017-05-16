package twitch

import (
	"bytes"
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
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
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

//Channel details of a channel
type Channel struct {
	ID                           int64     `json:"_id"`
	Name                         string    `json:"name"`
	DisplayName                  string    `json:"display_name"`
	Mature                       bool      `json:"mature"`
	Status                       string    `json:"status"`
	BroadcasterLanguage          string    `json:"broadcaster_language"`
	Game                         string    `json:"game"`
	Language                     string    `json:"language"`
	CreatedAt                    time.Time `json:"created_at,omitempty"`
	UpdatedAt                    time.Time `json:"updated_at,omitempty"`
	Logo                         string    `json:"logo"`
	VideoBanner                  string    `json:"video_banner"`
	ProfileBanner                string    `json:"profile_banner"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color"`
	Partner                      bool      `json:"partner"`
	URL                          string    `json:"url"`
	Views                        int64     `json:"views"`
	Followers                    int64     `json:"followers"`
}

//Upload details of an upload
type Upload struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

//Preview details of the different preview image sizes
type Preview struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Large    string `json:"large"`
	Template string `json:"template"`
}

//Thumbnails details of the different thumbnail sizes
type Thumbnails struct {
	Small    []interface{} `json:"small"`
	Medium   []interface{} `json:"medium"`
	Large    []interface{} `json:"large"`
	Template []interface{} `json:"template"`
}

//Video details of a video
type Video struct {
	ID            string `json:"_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	BroadcastID   int    `json:"broadcast_id"`
	BroadcastType string `json:"broadcast_type"`
	Status        string `json:"status"`
	TagList       string `json:"tag_list"`
	Views         int64  `json:"views"`
	URL           string `json:"url"`
	Language      string `json:"language"`
	Viewable      string `json:"viewable"`
	//ViewableAt    time.Time   `json:"viewable_at,omitempty"`
	//RecordedAt    time.Time   `json:"recorded_at,omitempty"`
	//CreatedAt     time.Time   `json:"created_at,omitempty"`
	Game        interface{} `json:"game"`
	Length      int64       `json:"length"`
	Preview     Preview     `json:"preview"`
	Thumbnails  Thumbnails  `json:"thumbnails"`
	Paywalled   bool        `json:"paywalled"`
	FPS         interface{} `json:"fps"`
	Resolutions interface{} `json:"resolutions"`
	Channel     Channel     `json:"channel"`
}

//
// VIDEO
//

//CreateVideoInput the parameters used to create a video - V4
type CreateVideoInput struct {
	ChannelName string
	Title       string
}

//CreateVideoOutput create the skeleton for a video upload
type CreateVideoOutput struct {
	Upload Upload `json:"upload"`
	Video  Video  `json:"video"`
}

//CompleteVideoInput the parameters used to complete a video upload
type CompleteVideoInput struct {
	VideoID string
	Token   string
}

//CompleteVideoOutput output from a completed video
type CompleteVideoOutput struct{}

//UploadVideoPartInput the parameters used to upload a video part
type UploadVideoPartInput struct {
	VideoID string
	Token   string
	Part    int
	Body    *bytes.Buffer
}

//UploadVideoPartOutput output from upoading a video part
type UploadVideoPartOutput struct {
	Upload Upload `json:"upload"`
	Video  Video  `json:"video"`
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
