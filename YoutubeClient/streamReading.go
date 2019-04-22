package youtubeclient

import (
	dgvoice "CookingBoy/DiscordVoice"

	"github.com/bwmarrin/discordgo"
)

// PlayStream takes a Stream direct URL and play it with ffmpeg
func PlayStream(vc *discordgo.VoiceConnection, directURL string, name string) {
	IsPlaying[vc.GuildID] = true
	NowPlayingChan <- name
	dgvoice.PlayAudioFile(vc, directURL, StopPlayerChans[vc.GuildID], PauseChan[vc.GuildID])
	IsPlaying[vc.GuildID] = false
	NowPlayingChan <- ""
}
