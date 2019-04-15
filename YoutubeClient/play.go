package youtubeclient

import (
	"log"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

func PlayVideo(URL string, VoiceConn *discordgo.VoiceConnection) {
	log.Println(VoiceConn.ChannelID)
	url := getCleannedURL(URL)
	switch state := IsPlaying[VoiceConn.ChannelID]; state {
	case false:
		IsPlaying[VoiceConn.ChannelID] = true
		dgvoice.PlayAudioFile(VoiceConn, url, StopPlayerChans[VoiceConn.ChannelID])
		IsPlaying[VoiceConn.ChannelID] = false
	case true:
		SongsQueues[VoiceConn.ChannelID] = append(SongsQueues[VoiceConn.ChannelID], URL)
	}
}
