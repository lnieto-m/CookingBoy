package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
	"regexp"
)

func (S *Sakamoto) play(args []string) {
	S.getVoiceConn()
	service, err := youtubeclient.YoutubeStart()
	if err != nil {
		log.Println(err)
		return
	}
	if len(args) > 0 {
		switch checkLinkValidity(args[0]) {
		case VIDEO:
			re := regexp.MustCompile(`(?m)v=([\w|-]*)`)
			id := ""
			for _, idMatch := range re.FindAllStringSubmatch(args[0], -1) {
				id = idMatch[1]
			}
			video, err := youtubeclient.GetVideoInfos(service, "snippet", id)
			if err != nil {
				log.Println(err)
				return
			}
			go youtubeclient.PlayVideo(video, S.voiceConn)
		case PLAYLIST:
			youtubeclient.QueuePlaylist(args[0], S.voiceConn)
		case NONVALIDLINK:
			S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "Please enter a valid Youtube video or playlist link.")
		}
	}
}

func (S *Sakamoto) stop(args []string) {
	S.getVoiceConn()
	youtubeclient.SongsQueues[S.discordMessageCreate.GuildID] = []youtubeclient.Video{}
	youtubeclient.StopPlayerChans[S.discordMessageCreate.GuildID] <- true
}
