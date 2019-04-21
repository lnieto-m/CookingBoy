package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"

	"github.com/bwmarrin/discordgo"
)

// UpdateGameStatus : Go routine that receive the song name from a channel and use the discord session
// to update the "Playing ... " message
func UpdateGameStatus(S *discordgo.Session, stop chan bool) {
	for {
		select {
		case name := <-youtubeclient.NowPlayingChan:
			S.UpdateStatus(0, name)
		case <-stop:
			return
		}
	}
}
