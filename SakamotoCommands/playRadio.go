package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
	"time"
)

func (S *Sakamoto) playRadio(args []string) {
	if len(args) == 0 {
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
			time.Sleep(250 * time.Millisecond)
		}
		youtubeclient.GetWaifuStream(S.voiceConn, value[0], value[1])
	}
}
