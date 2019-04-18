package youtubeclient

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func updateQueueVisual(S *discordgo.Session, M *discordgo.MessageReactionAdd, index int, message *QueueMessage) {
	titleContent := "No song playing."
	image := &discordgo.MessageEmbedThumbnail{}
	if index >= 0 && index < len(message.PageRange) {
		content := ""
		for _, value := range message.SongList[message.PageRange[index][0]:message.PageRange[index][1]] {
			content += value
		}
		message.CurrentPage = index

		if NowPlaying.URL != "" {
			titleContent = "[" + NowPlaying.Title + "](" + NowPlaying.URL + ")"
			image = &discordgo.MessageEmbedThumbnail{
				URL: NowPlaying.ThumbnailImageURL,
			}
		}

		field := &discordgo.MessageEmbedField{
			Name:  "Next",
			Value: content,
		}

		fields := []*discordgo.MessageEmbedField{
			field,
		}

		editedMessage := &discordgo.MessageEmbed{
			Title:       "Now Playing",
			Description: titleContent,
			Thumbnail:   image,
			Fields:      fields,
		}

		_, err := S.ChannelMessageEditEmbed(message.ChannelID, message.MessageID, editedMessage)
		if err != nil {
			log.Println("ManageQueuePage.go: ", err)
		}
	}
}

func ManageQueuePage(S *discordgo.Session, M *discordgo.MessageReactionAdd) {
	if M.Emoji.Name == "⬅" || M.Emoji.Name == "➡" {
		for _, cacheMessage := range QueueMessageCache {
			if cacheMessage.MessageID == M.MessageID {
				switch M.Emoji.Name {
				case "⬅":
					updateQueueVisual(S, M, cacheMessage.CurrentPage-1, cacheMessage)
				case "➡":
					updateQueueVisual(S, M, cacheMessage.CurrentPage+1, cacheMessage)
				}
			}
		}
		err := S.MessageReactionRemove(M.ChannelID, M.MessageReaction.MessageID, M.Emoji.Name, M.UserID)
		if err != nil {
			log.Println(err)
		}
	}
}
