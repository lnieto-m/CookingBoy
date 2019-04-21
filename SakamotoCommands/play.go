package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
	"regexp"
)

func (S *Sakamoto) play(args []string) {
	if S.performOriginChannelCheck() == false {
		S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "You must be in my voice channel to use this command.")
		return
	}
	S.getVoiceConn()
	service, err := youtubeclient.YoutubeStart()
	if err != nil {
		log.Println(err)
		return
	}
	if len(args) > 0 {

		// Check if either a video, playlist or something else
		switch checkLinkValidity(args[0]) {
		case VIDEO:

			// If a video, get the video id using a regex then calls the youtube API to get other infos needed
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

			// if a playlist, calls the youtube API to get all the playlist song infos then stores it
			youtubeclient.QueuePlaylist(args[0], S.voiceConn)
		case NONVALIDLINK:

			// Send an error message to the server
			S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "Please enter a valid Youtube video or playlist link.")
		}
	}
}

// Stop all the music functions, clearing song queues and stopping song playing
func (S *Sakamoto) stop(args []string) {
	S.getVoiceConn()
	youtubeclient.SongsQueues[S.discordMessageCreate.GuildID] = []youtubeclient.Video{}
	youtubeclient.StopPlayerChans[S.discordMessageCreate.GuildID] <- true
}
