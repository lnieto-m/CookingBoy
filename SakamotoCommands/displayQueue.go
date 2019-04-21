package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func (S *Sakamoto) getInfoForEmbed() (*discordgo.MessageEmbed, []string) {

	stop := 0
	content := ""
	titleContent := "No song playing."
	image := &discordgo.MessageEmbedThumbnail{}
	linkList := []string{}

	if youtubeclient.NowPlaying.URL != "" {
		titleContent = "[" + youtubeclient.NowPlaying.Title + "](" + youtubeclient.NowPlaying.URL + ")"
		image = &discordgo.MessageEmbedThumbnail{
			URL: youtubeclient.NowPlaying.ThumbnailImageURL,
		}
	}

	for id, videoElem := range youtubeclient.SongsQueues[S.discordMessageCreate.GuildID] {
		if stop == 0 {
			contentToAdd := strconv.Itoa(id+1) + ". [" + videoElem.Title + "](" + videoElem.URL + ")\n"
			if len(content)+len(contentToAdd) > 1024 {
				stop = 1
			} else {
				content += contentToAdd
			}
		}
		linkList = append(linkList, strconv.Itoa(id+1)+". ["+videoElem.Title+"]("+videoElem.URL+")\n")
	}
	if content == "" {
		content = "No song queued."
	}

	field := &discordgo.MessageEmbedField{
		Name:  "Next",
		Value: content,
	}

	fields := []*discordgo.MessageEmbedField{
		field,
	}

	message := &discordgo.MessageEmbed{
		Title:       "Now Playing",
		Description: titleContent,
		Thumbnail:   image,
		Fields:      fields,
	}

	return message, linkList
}

func (S *Sakamoto) displayQueue(args []string) {
	log.Println("lol")
	S.getVoiceConn()
	message, total := S.getInfoForEmbed()

	pageIndexs := [][2]int{}
	totalMessageLen, currentPageLen, lastPage, lastID := 0, 0, 0, 0
	for id, value := range total {
		totalMessageLen += len(value)
		currentPageLen += len(value)
		if currentPageLen > 1024 {
			pageIndexs = append(pageIndexs, [2]int{lastPage, id})
			lastPage = id
			currentPageLen = len(value)
		}
		lastID = id
	}

	if totalMessageLen > 1024 {
		totalPageCounter := len(pageIndexs)
		if currentPageLen > 0 {
			totalPageCounter++
		}
		message.Footer = &discordgo.MessageEmbedFooter{
			Text: "Page (1/" + strconv.Itoa(totalPageCounter) + ")",
		}
	}

	queue, err := S.discordSession.ChannelMessageSendEmbed(S.discordMessageCreate.ChannelID, message)
	if err != nil {
		log.Println("displayQueue: ", err.Error())
		return
	}

	// log.Println(to)

	if totalMessageLen > 1024 {
		if currentPageLen > 0 {
			pageIndexs = append(pageIndexs, [2]int{lastPage, lastID + 1})
		}
		log.Printf("%v\n", pageIndexs)
		youtubeclient.PushToQueueCache(&youtubeclient.QueueMessage{
			GuildID:     queue.GuildID,
			ChannelID:   queue.ChannelID,
			MessageID:   queue.ID,
			SongList:    total,
			PageRange:   pageIndexs,
			CurrentPage: 0,
		})
		err = S.discordSession.MessageReactionAdd(queue.ChannelID, queue.ID, "⬅")
		if err != nil {
			log.Println("Emoji ", err)
		}
		S.discordSession.MessageReactionAdd(queue.ChannelID, queue.ID, "➡")
	}

}
