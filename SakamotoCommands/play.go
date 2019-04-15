package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
)

func (S *Sakamoto) play(args []string) {
	// log.Println(args[0])
	S.getVoiceConn()
	if len(args) > 0 {
		go youtubeclient.PlayVideo(args[0], S.voiceConn)
	}
}

func (S *Sakamoto) stop(args []string) {
	S.getVoiceConn()
	youtubeclient.SongsQueues[S.voiceConn.ChannelID] = []string{}
	youtubeclient.StopPlayerChans[S.voiceConn.ChannelID] <- true
}
