package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
)

func (S *Sakamoto) play(args []string) {
	// log.Println(args[0])
	S.getVoiceConn()
	go youtubeclient.PlayVideo(args[0], S.voiceConn)
}
