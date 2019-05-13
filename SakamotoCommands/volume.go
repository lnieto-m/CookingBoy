package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
)

// One day
func (S *Sakamoto) volume(args []string) {
	// One day i'll find Da Wae
	// S.discordSession.VoiceConnections[]
}

// Pause the song playing by halting the ffmpeg process
// Also used to pause the song playing when using the soundbox
func (S *Sakamoto) pause(args []string) {
	if youtubeclient.IsPlaying[S.discordMessageCreate.GuildID] {
		if youtubeclient.PauseStates[S.discordMessageCreate.GuildID] {
			youtubeclient.PauseChan[S.discordMessageCreate.GuildID] <- false
			youtubeclient.PauseStates[S.discordMessageCreate.GuildID] = false
		} else {
			youtubeclient.PauseChan[S.discordMessageCreate.GuildID] <- true
			youtubeclient.PauseStates[S.discordMessageCreate.GuildID] = true
		}
	} else {
		log.Println("Couldnot find any song playing. Guild ID :", S.discordMessageCreate.GuildID)
	}
}
