package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
)

// Check the VoiceConnections map to get the voice connection corresponding to the current server
func (S *Sakamoto) getVoiceConn() {
	ok := false
	if S.voiceConn, ok = S.discordSession.VoiceConnections[S.discordMessageCreate.GuildID]; ok {
		return
	}
	log.Println("Error encountered joining voice channel, ")
}

// Leave voice, stopping all the music processes
func (S *Sakamoto) leaveVoice(args []string) {
	if S.performOriginChannelCheck() == false {
		S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "You must be in my voice channel to use this command.")
		return
	}
	S.stop(args)
	S.getVoiceConn()
	S.voiceConn.Disconnect()
}

// The bot join the User's voice channel
// Does nothing if the user is not in a voice channel
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

	for _, vs := range guild.VoiceStates {
		if vs.UserID == S.discordMessageCreate.Author.ID {
			S.voiceConn, err = S.discordSession.ChannelVoiceJoin(guild.ID, vs.ChannelID, false, false)
			if err != nil {
				log.Println("Error encountered joining voice channel, ", err)
				return
			}

			// Setup all the globals map entry corresponding to this server(Guild)
			youtubeclient.StopPlayerChans[S.discordMessageCreate.GuildID] = make(chan bool, 1)
			youtubeclient.IsPlaying[S.discordMessageCreate.GuildID] = false
			youtubeclient.SongsQueues[S.discordMessageCreate.GuildID] = []youtubeclient.Video{}
			youtubeclient.VoiceConnexions[S.discordMessageCreate.GuildID] = S.voiceConn.ChannelID
			youtubeclient.PauseChan[S.discordMessageCreate.GuildID] = make(chan bool, 1)
			youtubeclient.PauseStates[S.discordMessageCreate.GuildID] = false
			S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "At your service.")

			log.Println("Joined channel " + vs.ChannelID + " in server ID: " + S.discordMessageCreate.GuildID)
		}
	}
}
