package youtubeclient

import (
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// Create an embed Message to send it using the EditMessage Fuction
// Similar to sakamoto.getInfoForEmbed() method (merge possible ?)
// This one needs a starting point to get the corresponding queue page
func updateQueueVisual(S *discordgo.Session, M *discordgo.MessageReactionAdd, index int, message *QueueMessage) {
	titleContent := "No song playing."
	image := &discordgo.MessageEmbedThumbnail{}
	if index >= 0 && index < len(message.PageRange) {
		content := ""
		for _, value := range message.SongList[message.PageRange[index][0]:message.PageRange[index][1]] {
			content += value
		}

		// Update the page index
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

		footer := &discordgo.MessageEmbedFooter{
			Text: "Page (" + strconv.Itoa(message.CurrentPage+1) + "/" + strconv.Itoa(len(message.PageRange)) + ")",
		}

		editedMessage := &discordgo.MessageEmbed{
			Title:       "Now Playing",
			Description: titleContent,
			Thumbnail:   image,
			Fields:      fields,
			Footer:      footer,
		}

		_, err := S.ChannelMessageEditEmbed(message.ChannelID, message.MessageID, editedMessage)
		if err != nil {
			log.Println("ManageQueuePage.go: ", err)
		}
	}
}

// ManageQueuePage edits an embed message using the QueueCache
// Needed to display a oong song queue
// Arrows emoji let the user navigate through the song queue
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

		// Remove reactions added by users (except the bot itself) to allow multiple uses
		err := S.MessageReactionRemove(M.ChannelID, M.MessageReaction.MessageID, M.Emoji.Name, M.UserID)
		if err != nil {
			log.Println(err)
		}
	}
}
