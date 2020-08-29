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
	ID                           string    `json:"_id"`
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
	BroadcasterTyoe              string    `json:"broadcaster_type"`
	StreamKey                    string    `json:"stream_key"`
	Email                        string    `json:"email"`
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
