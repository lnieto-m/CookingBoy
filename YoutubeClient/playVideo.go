package youtubeclient

import (
	"log"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

func PlayVideo(URL string, VoiceConn *discordgo.VoiceConnection) {
	inQueue := 0

	defer func() {
		if inQueue == 1 && len(SongsQueues[VoiceConn.ChannelID]) > 0 {
			go PlayVideo(SongsQueues[VoiceConn.ChannelID][0], VoiceConn)
			if len(SongsQueues[VoiceConn.ChannelID][1:]) > 0 {
				SongsQueues[VoiceConn.ChannelID] = SongsQueues[VoiceConn.ChannelID][1:]
			} else {
				SongsQueues[VoiceConn.ChannelID] = []string{}
			}
		}
	}()

	log.Println(VoiceConn.ChannelID)

	url := getCleannedURL(URL)
	log.Println(url)
	switch state := IsPlaying[VoiceConn.ChannelID]; state {
	case false:
		inQueue = 1
		IsPlaying[VoiceConn.ChannelID] = true
		NowPlaying = URL
		dgvoice.PlayAudioFile(VoiceConn, url, StopPlayerChans[VoiceConn.ChannelID])
		IsPlaying[VoiceConn.ChannelID] = false
		NowPlaying = ""
	case true:
		SongsQueues[VoiceConn.ChannelID] = append(SongsQueues[VoiceConn.ChannelID], URL)
		log.Printf("Song queued: %v\n", SongsQueues[VoiceConn.ChannelID])
	}
}
