package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func (S *Sakamoto) getInfoForEmbed() *discordgo.MessageEmbed {

	content := ""
	titleContent := "No song playing."
	image := &discordgo.MessageEmbedThumbnail{}

	if youtubeclient.NowPlaying.URL != "" {
		titleContent = "[" + youtubeclient.NowPlaying.Title + "](" + youtubeclient.NowPlaying.URL + ")"
		image = &discordgo.MessageEmbedThumbnail{
			URL: youtubeclient.NowPlaying.ThumbnailImageURL,
		}
	}

	for id, videoElem := range youtubeclient.SongsQueues[S.discordMessageCreate.GuildID] {
		content += strconv.Itoa(id+1) + ". [" + videoElem.Title + "](" + videoElem.URL + ")\n"
	}
	if content == "" {
		content = "No song queued."
	}

	field := &discordgo.MessageEmbedField{
		Name:  "Next",
		Value: content,
	}

	// TODO : ADD PAGINATION
	log.Println("CONTENT : ", len(youtubeclient.SongsQueues[S.discordMessageCreate.GuildID]))

	table := []*discordgo.MessageEmbedField{
		field,
	}

	message := &discordgo.MessageEmbed{
		Title:       "Now Playing",
		Description: titleContent,
		Fields:      table,
		Thumbnail:   image,
	}

	return message
}

func (S *Sakamoto) displayQueue(args []string) {
	log.Println("lol")
	S.getVoiceConn()
	message := S.getInfoForEmbed()
	S.discordSession.ChannelMessageSendEmbed(S.discordMessageCreate.ChannelID, message)
}
