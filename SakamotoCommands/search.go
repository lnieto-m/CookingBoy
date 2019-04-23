package sakamotocommands

import (
	imgurgo "CookingBoy/ImgurGo"
	"log"
	"strings"
)

func avalaibleOption(query string, options []string) bool {
	for _, value := range options {
		if value == query {
			return true
		}
	}
	return false
}

func (S *Sakamoto) search(args []string) {
	if len(args) == 0 {
		return
	}
	query := strings.Join(args, " ")
	log.Printf("Sort: %s Window: %s Query: %s\n", SEARCHSORT, SEARCHRANGE, query)
	imgURL := imgurgo.GetImage(SEARCHSORT, SEARCHRANGE, "\""+query+"\"")
	if imgURL == "" {
		return
	}
	S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, imgURL)
}

func (S *Sakamoto) searchSort(args []string) {
	if len(args) != 0 && avalaibleOption(args[0], SearchSortOptions) {
		SEARCHSORT = args[0]
	}
}

func (S *Sakamoto) searchRange(args []string) {
	if len(args) != 0 && avalaibleOption(args[0], SearchRangeOptions) {
		SEARCHRANGE = args[0]
	}
}
