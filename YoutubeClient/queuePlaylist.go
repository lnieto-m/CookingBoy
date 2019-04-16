package youtubeclient

import (
	"log"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func QueuePlaylist(URL string, VoiceConn *discordgo.VoiceConnection) {

	service, err := YoutubeStart()
	if err != nil {
		log.Println(err)
		return
	}

	playlistRe := regexp.MustCompile(`(?m)list=([\w|-]*)`)

	for _, match := range playlistRe.FindAllStringSubmatch(URL, -1) {
		log.Println(match[1])
		videoIds := GetPlaylistLinks(service, "contentDetails", match[1])
		for _, item := range videoIds {
			SongsQueues[VoiceConn.ChannelID] = append(SongsQueues[VoiceConn.ChannelID], "https://www.youtube.com/watch?v="+item)
		}

		// TODO : ADD LECTURE

		log.Printf("%v\n", SongsQueues[VoiceConn.ChannelID])
	}
}
