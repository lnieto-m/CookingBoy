package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"

	"github.com/bwmarrin/discordgo"
)

// General help message
func (S *Sakamoto) help(args []string) {

	if len(args) > 0 {
		if args[0] == "soundbox" {
			S.displaySoundBoxHelp()
			return
		} else if args[0] == "radio" {
			S.displayRadioHelp()
			return
		} else if args[0] == "search" {
			S.displaySearchHelp()
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

	radioField := &discordgo.MessageEmbedField{
		Name:  ":headphones: Radio",
		Value: "`s!radio <radio_name>`\nType `s!help radio` to display avalaible radios",
	}

	searchField := &discordgo.MessageEmbedField{
		Name:  ":mag: Image Search",
		Value: "`search`,`search_sort`,`search_range`\nType `s!help search` for more details about image search",
	}

	fields := []*discordgo.MessageEmbedField{
		musicField,
		soundBoxField,
		radioField,
		searchField,
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

// SoundBox help message
func (S *Sakamoto) displaySoundBoxHelp() {

	author := &discordgo.MessageEmbedAuthor{
		Name:    "Avalaible sounds",
		IconURL: "https://pbs.twimg.com/profile_images/627117609444581380/7YG7kxA4_400x400.png",
	}

	description := "`JEAGER`,`JEANNE`,`cklair`,`whee`,`bruh`,`oof`,`marionon`,`thomas`,`sanic`,`running`,`SPITONHIM`,`dewae`,`johncena`,`sensibilite`,`qualifie`,`vega`,`rengar`,`quenouille`,`mince`,`troposphere`,`ratz`,`doremi`,`guile`,`zombie`,`ally`"

	soundBoxHelpMessage := &discordgo.MessageEmbed{
		Title:       "",
		Description: description,
		Author:      author,
		Color:       0x98BDF0,
	}

	S.discordSession.ChannelMessageSendEmbed(S.discordMessageCreate.ChannelID, soundBoxHelpMessage)
}

func (S *Sakamoto) displayRadioHelp() {

	description := ""

	for key, value := range youtubeclient.LoadedRadios {
		description += "`" + key + "`: " + value[2] + "\n"
	}

	author := &discordgo.MessageEmbedAuthor{
		Name:    "Avalaible radios",
		IconURL: "https://c-ash.smule.com/sf/s33/arr/f3/22/1f68ac92-e47a-499f-91fc-19b284ea6be3_256.jpg",
	}

	radioHelpMessage := &discordgo.MessageEmbed{
		Description: description,
		Author:      author,
		Color:       0x98BDF0,
	}

	S.discordSession.ChannelMessageSendEmbed(S.discordMessageCreate.ChannelID, radioHelpMessage)
}

func (S *Sakamoto) displaySearchHelp() {

	author := &discordgo.MessageEmbedAuthor{
		Name:    "Search Engine using Imgur",
		IconURL: "https://i.imgur.com/tF3g1JA.jpg",
	}

	footer := &discordgo.MessageEmbedFooter{
		Text: "Powered by Imgur.(no lol)",
	}

	searchHelpMessage := &discordgo.MessageEmbed{
		Description: "**Current search options:**\nSort: `" + SEARCHSORT + "`\nRange: `" + SEARCHRANGE + "`\nUse `search_sort` and `search_range` to change them.\n\n**Avaliable options:**\nSort:`top`,`rising`,`viral`,`time`\nRange:`day`,`week`,`month`,`year`,`all`",
		Author:      author,
		Footer:      footer,
	}
	S.discordSession.ChannelMessageSendEmbed(S.discordMessageCreate.ChannelID, searchHelpMessage)
}
