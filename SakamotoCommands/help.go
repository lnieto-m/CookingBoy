package sakamotocommands

func (S *Sakamoto) help(args []string) {

	if len(args) > 0 {
		if args[0] == "soundbox" {
			S.displaySoundBoxHelp()
			return
		}
	}
	help := `
	`

	S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "```"+help+"```")
}

func (S *Sakamoto) displaySoundBoxHelp() {
	help := `SoundBox commands:
JEAGER
JEANNE
cklair
whee
bruh
oof
marionon
thomas
sanic
running
SPITONHIM
dewae
johncena
sensibilite
qualifie
vega
rengar
quenouille
mince
troposphere`
	S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "```"+help+"```")
}
