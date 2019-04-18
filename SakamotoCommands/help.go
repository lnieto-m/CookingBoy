package sakamotocommands

func (S *Sakamoto) help(args []string) {

	help := `
	`

	S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "```"+help+"```")
}
