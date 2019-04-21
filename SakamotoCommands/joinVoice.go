package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
)

func (S *Sakamoto) getVoiceConn() {
	ok := false
	if S.voiceConn, ok = S.discordSession.VoiceConnections[S.discordMessageCreate.GuildID]; ok {
		return
	}
	log.Println("Error encountered joining voice channel, ")
}

func (S *Sakamoto) leaveVoice(args []string) {
	if S.performOriginChannelCheck() == false {
		S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "You must be in my voice channel to use this command.")
		return
	}
	S.stop(args)
	S.getVoiceConn()
	S.voiceConn.Disconnect()
}

func (S *Sakamoto) joinVoice(args []string) {
	channelOrigin, err := S.discordSession.State.Channel(S.discordMessageCreate.ChannelID)
	if err != nil {
		log.Println("Could not find channel. ", err)
		return
	}

	guild, err := S.discordSession.State.Guild(channelOrigin.GuildID)
	if err != nil {
		log.Println("Could not find guild. ", err)
		return
	}

	log.Println("Origin :", channelOrigin.ID, "Guild: ", guild.ID)

	for _, vs := range guild.VoiceStates {
		if vs.UserID == S.discordMessageCreate.Author.ID {
			S.voiceConn, err = S.discordSession.ChannelVoiceJoin(guild.ID, vs.ChannelID, false, false)
			if err != nil {
				log.Println("Error encountered joining voice channel, ", err)
				return
			}

			log.Println(S.discordMessageCreate.GuildID)
			youtubeclient.StopPlayerChans[S.discordMessageCreate.GuildID] = make(chan bool, 1)
			youtubeclient.IsPlaying[S.discordMessageCreate.GuildID] = false
			youtubeclient.SongsQueues[S.discordMessageCreate.GuildID] = []youtubeclient.Video{}
			youtubeclient.VoiceConnexions[S.discordMessageCreate.GuildID] = S.voiceConn.ChannelID
			youtubeclient.PauseChan[S.discordMessageCreate.GuildID] = make(chan bool, 1)
			youtubeclient.PauseStates[S.discordMessageCreate.GuildID] = false
			// S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "At your service.")
		}
	}
}
