package youtubeclient

import (
	"log"

	dgvoice "CookingBoy/DiscordVoice"

	"github.com/bwmarrin/discordgo"
)

// PlayVideo use a given Video object to play its sounds through the given VoiceConnexion
func PlayVideo(video Video, VoiceConn *discordgo.VoiceConnection) {
	inQueue := 0

	defer func() {
		if inQueue == 1 && len(SongsQueues[VoiceConn.GuildID]) > 0 {
			go PlayVideo(SongsQueues[VoiceConn.GuildID][0], VoiceConn)
			if len(SongsQueues[VoiceConn.GuildID][1:]) > 0 {
				SongsQueues[VoiceConn.GuildID] = SongsQueues[VoiceConn.GuildID][1:]
			} else {
				SongsQueues[VoiceConn.GuildID] = []Video{}
			}
		}
	}()

	log.Println(VoiceConn.GuildID)

	url := getCleannedURL(video.URL)
	log.Println(url)

	switch state := IsPlaying[VoiceConn.GuildID]; state {
	case false:
		inQueue = 1
		IsPlaying[VoiceConn.GuildID] = true
		NowPlaying = video
		dgvoice.PlayAudioFile(VoiceConn, url, StopPlayerChans[VoiceConn.GuildID], PauseChan[VoiceConn.GuildID])
		IsPlaying[VoiceConn.GuildID] = false
		NowPlaying = Video{}
	case true:
		SongsQueues[VoiceConn.GuildID] = append(SongsQueues[VoiceConn.GuildID], video)
		log.Printf("Song queued: %v\n", SongsQueues[VoiceConn.GuildID])
	}
}
