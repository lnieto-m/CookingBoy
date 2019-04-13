package sakamotocommands

import "log"

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
			S.discordSession.ChannelVoiceJoin(guild.ID, vs.ChannelID, false, false)
			S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "At your service.")
		}
	}
}
