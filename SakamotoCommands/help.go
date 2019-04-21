package sakamotocommands

import (
	"github.com/bwmarrin/discordgo"
)

func (S *Sakamoto) help(args []string) {

	if len(args) > 0 {
		if args[0] == "soundbox" {
			S.displaySoundBoxHelp()
			return
		}
	}

	musicField := &discordgo.MessageEmbedField{
		Name:  ":musical_note: Music",
		Value: "`join`,`play`,`queue`,`stop`,`skip`,`leave`,`pause`",
	}

	soundBoxField := &discordgo.MessageEmbedField{
		Name:  ":sound: SoundBox",
		Value: "`s!sound <sound_name>`\nType `s!help soundbox` to display avalaible sounds",
	}

	fields := []*discordgo.MessageEmbedField{
		musicField,
		soundBoxField,
	}

	author := &discordgo.MessageEmbedAuthor{
		Name:    "Usage",
		IconURL: "https://cdn.discordapp.com/avatars/566552423972339722/2ae2df97a8c1b0fd1d46ee9589fb3a04.png",
	}

	helpMessage := &discordgo.MessageEmbed{
		Title:       "",
		Description: "`s!<base_command> <args>`",
		Fields:      fields,
		Author:      author,
		Color:       0x98BDF0,
	}

	S.discordSession.ChannelMessageSendEmbed(S.discordMessageCreate.ChannelID, helpMessage)
}

func (S *Sakamoto) displaySoundBoxHelp() {

	author := &discordgo.MessageEmbedAuthor{
		Name:    "Avalaible sounds",
		IconURL: "https://pbs.twimg.com/profile_images/627117609444581380/7YG7kxA4_400x400.png",
	}

	description := "`JEAGER`,`JEANNE`,`cklair`,`whee`,`bruh`,`oof`,`marionon`,`thomas`,`sanic`,`running`,`SPITONHIM`,`dewae`,`johncena`,`sensibilite`,`qualifie`,`vega`,`rengar`,`quenouille`,`mince`,`troposphere`,`ratz`,`doremi`,`guile`"

	soundBoxHelpMessage := &discordgo.MessageEmbed{
		Title:       "",
		Description: description,
		Author:      author,
		Color:       0x98BDF0,
	}

	S.discordSession.ChannelMessageSendEmbed(S.discordMessageCreate.ChannelID, soundBoxHelpMessage)
}
