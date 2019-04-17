package youtubeclient

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ManageQueuePage(S *discordgo.Session, M *discordgo.MessageReactionAdd) {
	if M.Emoji.Name == "⬅" || M.Emoji.Name == "➡" {

		err := S.MessageReactionRemove(M.ChannelID, M.MessageReaction.MessageID, M.Emoji.Name, M.UserID)
		if err != nil {
			log.Println(err)
		}
	}
}
