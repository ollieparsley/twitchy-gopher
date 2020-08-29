package twitch

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestCreateVideo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.twitch.tv/kraken/videos",
		httpmock.NewStringResponder(200, `{"upload":{"url":"https://uploads.twitch.tv/upload/123456","token":"this-is-a-token"},"video":{"title":"Test TitleD","description":null,"broadcast_id":1,"broadcast_type":"upload","status":"created","tag_list":"","views":0,"url":"https://www.twitch.tv/ollieparsleydev/v/103787281","language":"en","viewable":"public","viewable_at":null,"_id":"v123456","recorded_at":"2016-11-27T20:12:25Z","game":null,"length":0,"preview":{"small":"https://vod-secure.twitch.tv/_404/404_processing_80x45.png","medium":"https://vod-secure.twitch.tv/_404/404_processing_320x180.png","large":"https://vod-secure.twitch.tv/_404/404_processing_640x360.png","template":"https://vod-secure.twitch.tv/_404/404_processing_{width}x{height}.png"},"thumbnails":{"small":[],"medium":[],"large":[],"template":[]},"paywalled":false,"fps":{},"resolutions":{},"created_at":"2016-11-27T20:12:25Z","published_at":null,"_links":{"self":"https://api.twitch.tv/kraken/videos/v103787281","channel":"https://api.twitch.tv/kraken/channels/ollieparsleydev"},"channel":{"name":"ollieparsleydev","display_name":"ollieparsleydev"}}}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.CreateVideo(&CreateVideoInput{
		ChannelName: "ollieparsleydev",
		//V5 ChannelID: 139985889,
		Title: "Test upload",
	})

	if errorOutput != nil {
		t.Errorf("CreateVideo errorOutput should have been nil: %+v", errorOutput)
	}

	if output == nil {
		t.Errorf("CreateVideo output shouldn't have been nil")
	}

	if output.Upload.URL != "https://uploads.twitch.tv/upload/123456" {
		t.Errorf("CreateVideo output upload URL was not \"https://uploads.twitch.tv/upload/123456\": %s", output.Upload.URL)
	}

	if output.Upload.Token != "this-is-a-token" {
		t.Errorf("CreateVideo output upload token was not \"this-is-a-token\": %s", output.Upload.Token)
	}

	if output.Video.ID != "v123456" {
		t.Errorf("CreateVideo output video ID was not \"v123456\": %s", output.Video.ID)
	}

}

/*func TestUploadVideoPart(t *testing.T) {
	client := NewClient(&OAuthConfig{
		ClientID:    "pzqv1a6n4r1l7wzto3mor00bzkpmw8c",
		AccessToken: "xte5p3cozk1tbv2gbmnalobl9z77vu",
	}, &http.Client{})

	bytes.NewBufferString("foobar")

	uploadVideoPartOutput, uploadVideoPartError := client.UploadVideoPart(&UploadVideoPartInput{
		VideoID: "",
		Part:    1,
		Token:   createOutput.Upload.Token,
		Body:    buf,
	})
	if uploadPartErrorOutput != nil {
		t.Errorf("TestUploadVideo uploadPartErrorOutput: %+v", uploadPartErrorOutput)
	}

	_, completeErrorOutput := client.CompleteVideo(&CompleteVideoInput{
		VideoID: createOutput.Video.ID,
		Token:   createOutput.Upload.Token,
	})
	if completeErrorOutput != nil {
		t.Errorf("TestUploadVideo completeErrorOutput: %+v", createErrorOutput)
	}

}*/

/*func TestUploadVideo(t *testing.T) {
	client := NewClient(&OAuthConfig{
		ClientID:    "pzqv1a6n4r1l7wzto3mor00bzkpmw8c",
		AccessToken: "xte5p3cozk1tbv2gbmnalobl9z77vu",
	}, &http.Client{})
	client.apiVersion = 4

	createOutput, createErrorOutput := client.CreateVideo(&CreateVideoInput{
		ChannelName: "ollieparsleydev",
		Title:       "Test upload",
	})
	if createErrorOutput != nil {
		t.Errorf("TestUploadVideo createErrorOutput: %+v", createErrorOutput)
	}

	fmt.Printf("\n\nCREATE OUTPUT: %+v\n\n", createOutput)

	buf := bytes.NewBuffer(nil)
	f, _ := os.Open("/home/ollie/streaming.mp4") // Error handling elided for brevity.
	io.Copy(buf, f)                              // Error handling elided for brevity.
	f.Close()

	fmt.Printf("CREATE OUTPUT: %+v", createOutput)

	_, uploadPartErrorOutput := client.UploadVideoPart(&UploadVideoPartInput{
		VideoID: createOutput.Video.ID,
		Part:    1,
		Token:   createOutput.Upload.Token,
		Body:    buf,
	})
	if uploadPartErrorOutput != nil {
		t.Errorf("TestUploadVideo uploadPartErrorOutput: %+v", uploadPartErrorOutput)
	}

	_, completeErrorOutput := client.CompleteVideo(&CompleteVideoInput{
		VideoID: createOutput.Video.ID,
		Token:   createOutput.Upload.Token,
	})
	if completeErrorOutput != nil {
		t.Errorf("TestUploadVideo completeErrorOutput: %+v", createErrorOutput)
	}

}*/
