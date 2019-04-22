package sakamotocommands

import (
	dgvoice "CookingBoy/DiscordVoice"
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
	"time"
)

// Play a sound avalaible in the soundbox
// If the sound does not exist, show the soundbox commands
func (S *Sakamoto) soundBox(args []string) {
	if len(args) <= 0 || S.performOriginChannelCheck() == false {
		return
	}
	if path, ok := Soundbox[args[0]]; ok {
		if youtubeclient.IsPlayingSound[S.discordMessageCreate.GuildID] {
			log.Println("Sound already playing.")
			return
		}
		err := S.getVoiceConn()
		if err != nil {
			log.Println(err.Error())
			return
		}
		S.pause(nil)
		if youtubeclient.IsPlaying[S.discordMessageCreate.GuildID] {
			time.Sleep(750 * time.Millisecond)
		}
		youtubeclient.IsPlayingSound[S.discordMessageCreate.GuildID] = true
		dgvoice.PlayAudioFile(S.voiceConn, path, nil, nil)
		youtubeclient.IsPlayingSound[S.discordMessageCreate.GuildID] = false
		if youtubeclient.IsPlaying[S.discordMessageCreate.GuildID] {
			time.Sleep(250 * time.Millisecond)
		}
		S.pause(nil)
		return
	}
	S.displaySoundBoxHelp()
}
