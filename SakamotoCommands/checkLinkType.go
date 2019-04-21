package sakamotocommands

import (
	"regexp"
)

func checkLinkValidity(URL string) int {

	// regex missing a check for country extension
	// Currently accept any .extension
	videoRe := regexp.MustCompile(`(?m)(https|http)://(w{3}\.)?youtube\.\w+\/watch\?v=([\w|-]*)`)
	playlistRe := regexp.MustCompile(`(?m)list=([\w|-]*)`)

	// Try matches for a youtube video or playlist link, return NONVALIDLINK if no matches
	for range videoRe.FindAllStringSubmatch(URL, -1) {
		for range playlistRe.FindAllStringSubmatch(URL, -1) {
			return PLAYLIST
		}
		return VIDEO
	}
	return NONVALIDLINK
}
