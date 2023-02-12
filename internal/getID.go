package getYoutube

import (
	"regexp"
)

func GetID(url string) string {
	var id string

	regex := regexp.MustCompile(`^(?:https?://)?(?:www\.)?(?:youtu\.be/|youtube\.com/watch\?v=)([A-Za-z0-9_-]{11})`)

	match := regex.FindStringSubmatch(url)
	if len(match) > 0 {
		id = match[1]
	}

	return id
}
