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
	if len(args) != 0 && avalaibleOption(args[0], SearchSortOptions) && args[0] != SEARCHSORT {
		oldOption := SEARCHSORT
		SEARCHSORT = args[0]
		S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "Search sort option changed from `"+oldOption+"` to `"+SEARCHSORT+"`.")
		return
	}
	S.displaySearchHelp()
}

func (S *Sakamoto) searchRange(args []string) {
	if len(args) != 0 && avalaibleOption(args[0], SearchRangeOptions) && args[0] != SEARCHRANGE {
		oldOption := SEARCHRANGE
		SEARCHRANGE = args[0]
		S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "Search range option changed from `"+oldOption+"` to `"+SEARCHRANGE+"`.")
		return
	}
	S.displaySearchHelp()
}
