package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
	"time"
)

// Check if the given radio exists and plays it
func (S *Sakamoto) playRadio(args []string) {
	if len(args) == 0 {
		S.displayRadioHelp()
		return
	}
	err := S.getVoiceConn()
	if err != nil {
		log.Println(err.Error())
		return
	}
	if value, ok := youtubeclient.LoadedRadios[args[0]]; ok {
		if youtubeclient.IsPlaying[S.discordMessageCreate.GuildID] {
			youtubeclient.StopPlayerChans[S.discordMessageCreate.GuildID] <- true
			time.Sleep(750 * time.Millisecond)
		}
		youtubeclient.PlayStream(S.voiceConn, value[0], value[1])
		return
	}
	S.displayRadioHelp()
}
