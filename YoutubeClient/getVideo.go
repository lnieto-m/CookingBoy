package youtubeclient

import (
	"log"

	"github.com/otium/ytdl"
)

// Use the otium/ytdl package to get a direct link to the mp4 file of a youtube video
func getCleannedURL(URL string) string {
	video, err := ytdl.GetVideoInfo(URL)

	if err != nil {
		log.Println(err)
		return ""
	}

	for _, format := range video.Formats {
		if format.AudioEncoding == "opus" || format.AudioEncoding == "aac" || format.AudioEncoding == "vorbis" {
			data, err := video.GetDownloadURL(format)
			if err != nil {
				log.Println(err)
			}
			url := data.String()
			return url
		}
	}
	return ""
}
