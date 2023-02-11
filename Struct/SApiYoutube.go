package getYoutube

import (
	"fmt"

	"github.com/senseyeio/duration"
)

type SApiYoutube struct {
	Items []struct {
		Snippet struct {
			Title        string `json:"title"`
			ChannelTitle string `json:"channelTitle"`
		} `json:"snippet"`
		ContentDetails struct {
			Duration string `json:"duration"`
		} `json:"contentDetails"`
		Statistics struct {
			ViewCount     string `json:"viewCount"`
			LikeCount     string `json:"likeCount"`
			FavoriteCount string `json:"favoriteCount"`
			CommentCount  string `json:"commentCount"`
		} `json:"statistics"`
	} `json:"items"`
}

func (a *SApiYoutube) GetTitle() string {
	return fmt.Sprintf("\nTitle: %v", a.Items[0].Snippet.Title)
}

func (a *SApiYoutube) GetChannel() string {
	return fmt.Sprintf("\nChannel: %v", a.Items[0].Snippet.ChannelTitle)
}

func (a *SApiYoutube) GetDuration() string {
	duration, _ := duration.ParseISO8601(a.Items[0].ContentDetails.Duration)
	return fmt.Sprintf("\nDuration: %v hours, %v minutes and %v seconds \n", duration.TH, duration.TM, duration.TS)
}

func (a *SApiYoutube) GetViews() string {
	return fmt.Sprintf("\nViews: %v", a.Items[0].Statistics.ViewCount)
}

func (a *SApiYoutube) GetLikes() string {
	return fmt.Sprintf("\nLikes: %v", a.Items[0].Statistics.LikeCount)
}

func (a *SApiYoutube) GetFavorites() string {
	return fmt.Sprintf("\nFavorites: %v", a.Items[0].Statistics.FavoriteCount)
}

func (a *SApiYoutube) GetComments() string {
	return fmt.Sprintf("\nComments: %v", a.Items[0].Statistics.CommentCount)
}
