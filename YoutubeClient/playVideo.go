package youtubeclient

import (
	"log"

	dgvoice "CookingBoy/DiscordVoice"

	"github.com/bwmarrin/discordgo"
)

// PlayVideo use a given Video object to play its sounds through the given VoiceConnexion
func PlayVideo(video Video, VoiceConn *discordgo.VoiceConnection) {
	inQueue := 0

	// When the current song ends, check the next one in queue and plays it if it exists
	// This allow an automatic looping through song queue
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

	// Get the direct mp4 url for a youtube link
	url := getCleannedURL(video.URL)

	log.Println(url)

	// This switch check if a song is currently playing in this server
	// If not, plays the given song link
	// Else queue it
	switch state := IsPlaying[VoiceConn.GuildID]; state {
	case false:
		inQueue = 1
		IsPlaying[VoiceConn.GuildID] = true
		NowPlaying = video
		NowPlayingChan <- video.Title
		dgvoice.PlayAudioFile(VoiceConn, url, StopPlayerChans[VoiceConn.GuildID], PauseChan[VoiceConn.GuildID])
		IsPlaying[VoiceConn.GuildID] = false
		NowPlaying = Video{}
		NowPlayingChan <- ""
	case true:
		SongsQueues[VoiceConn.GuildID] = append(SongsQueues[VoiceConn.GuildID], video)
		log.Printf("Song queued: %v\n", SongsQueues[VoiceConn.GuildID])
	}
}
