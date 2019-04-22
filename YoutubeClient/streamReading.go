package youtubeclient

import (
	dgvoice "CookingBoy/DiscordVoice"

	"github.com/bwmarrin/discordgo"
)

// On testing, Now using a preset of radios
func GetWaifuStream(vc *discordgo.VoiceConnection, directURL string, name string) {
	IsPlaying[vc.GuildID] = true
	NowPlayingChan <- name
	dgvoice.PlayAudioFile(vc, directURL, StopPlayerChans[vc.GuildID], nil)
	IsPlaying[vc.GuildID] = false
	NowPlayingChan <- ""
}
