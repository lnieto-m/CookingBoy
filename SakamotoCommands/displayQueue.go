package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
	"regexp"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func (S *Sakamoto) getInfoForEmbed() *discordgo.MessageEmbed {
	service, err := youtubeclient.YoutubeStart()
	if err != nil {
		log.Println(err)
		return nil
	}

	content := ""
	re := regexp.MustCompile(`(?m)v=(\w*)`)

	imgURL := ""
	titleContent := "No song playing."

	if youtubeclient.NowPlaying != "" {
		for _, titleMatch := range re.FindAllStringSubmatch(youtubeclient.NowPlaying, -1) {
			title, err := youtubeclient.GetLinkTitle(service, "snippet", titleMatch[1])
			if err != nil {
				log.Println(err)
				return nil
			}
			imgURL, err = youtubeclient.GetImageLink(service, "snippet", titleMatch[1])
			if err != nil {
				log.Println(err)
			}
			titleContent = "[" + title + "](" + youtubeclient.NowPlaying + ")"
		}
	}

	for id, rawLink := range youtubeclient.SongsQueues[S.voiceConn.ChannelID] {
		for _, match := range re.FindAllStringSubmatch(rawLink, -1) {
			link, err := youtubeclient.GetLinkTitle(service, "snippet", match[1])
			if err != nil {
				log.Println(err)
				return nil
			}
			content += strconv.Itoa(id+1) + ". [" + link + "](" + rawLink + ")\n"
		}
	}
	if content == "" {
		content = "No song queued."
	}

	field := &discordgo.MessageEmbedField{
		Name:  "Next",
		Value: content,
	}

	table := []*discordgo.MessageEmbedField{
		field,
	}

	image := &discordgo.MessageEmbedThumbnail{
		URL: imgURL,
		// Width:  int(imgWidth),
		// Height: int(imgHeight),
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
