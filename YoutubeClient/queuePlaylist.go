package youtubeclient

import (
	"log"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

// QueuePlaylist use a URL string to queue playlist element from that url and play it if no elements are queued
func QueuePlaylist(URL string, VoiceConn *discordgo.VoiceConnection) {

	service, err := YoutubeStart()
	if err != nil {
		log.Println(err)
		return
	}

	playlistRe := regexp.MustCompile(`(?m)list=([\w|-]*)`)

	for _, match := range playlistRe.FindAllStringSubmatch(URL, -1) {
		log.Println(match[1])
		videos := GetPlaylistVideos(service, "snippet,contentDetails", match[1])
		SongsQueues[VoiceConn.GuildID] = append(SongsQueues[VoiceConn.GuildID], videos...)

		if IsPlaying[VoiceConn.GuildID] == false && len(SongsQueues[VoiceConn.GuildID]) > 0 {
			go PlayVideo(SongsQueues[VoiceConn.GuildID][0], VoiceConn)
			if len(SongsQueues[VoiceConn.GuildID][1:]) > 0 {
				SongsQueues[VoiceConn.GuildID] = SongsQueues[VoiceConn.GuildID][1:]
			} else {
				SongsQueues[VoiceConn.GuildID] = []Video{}
			}
		}
	}
}
