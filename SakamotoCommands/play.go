package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
)

func (S *Sakamoto) play(args []string) {
	// log.Println(args[0])
	S.getVoiceConn()
	if len(args) > 0 {
		switch checkLinkValidity(args[0]) {
		case VIDEO:
			go youtubeclient.PlayVideo(args[0], S.voiceConn)
		case PLAYLIST:
			youtubeclient.QueuePlaylist(args[0], S.voiceConn)
		case NONVALIDLINK:
			S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "Please enter a valid Youtube video or playlist link.")
		}
	}
}

func (S *Sakamoto) stop(args []string) {
	S.getVoiceConn()
	youtubeclient.SongsQueues[S.voiceConn.ChannelID] = []string{}
	youtubeclient.StopPlayerChans[S.voiceConn.ChannelID] <- true
}
