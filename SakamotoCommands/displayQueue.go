package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// Prepare a *discordgo.MessageEmbed using the stored song queue in SongsQueue[GuildID]
// Returns a MessageEmbed object, and a []string used to determine if the message will need futur modifications
func (S *Sakamoto) getInfoForEmbed() (*discordgo.MessageEmbed, []string) {

	stop := 0
	content := ""
	titleContent := "No song playing."
	image := &discordgo.MessageEmbedThumbnail{}
	linkList := []string{}

	// Get the song currently playing
	if youtubeclient.NowPlaying.URL != "" {
		titleContent = "[" + youtubeclient.NowPlaying.Title + "](" + youtubeclient.NowPlaying.URL + ")"
		image = &discordgo.MessageEmbedThumbnail{
			URL: youtubeclient.NowPlaying.ThumbnailImageURL,
		}
	}

	// Get elements in song queue
	// Only stores up to 1024 characters due to Discord API limit
	// The whole queue is stored in a []string to perform futur modifications
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

// Send an Embed message containing current song queue infos retrieved by S.getInfoForEmbed()
// if the queue is too long to fit in one EmbedMessage arrows reactions will be added to allow a complete view of the song queue
func (S *Sakamoto) displayQueue(args []string) {

	S.getVoiceConn()
	message, total := S.getInfoForEmbed()

	// Setup Pagination for a long song queue
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

	// Send the created EmbedMessage
	queue, err := S.discordSession.ChannelMessageSendEmbed(S.discordMessageCreate.ChannelID, message)
	if err != nil {
		log.Println("displayQueue: ", err.Error())
		return
	}

	// Stores queue message in a cache if the message is longer than 1024 characters
	// Then add arrow reactions
	if totalMessageLen > 1024 {
		if currentPageLen > 0 {
			pageIndexs = append(pageIndexs, [2]int{lastPage, lastID + 1})
		}
		youtubeclient.PushToQueueCache(&youtubeclient.QueueMessage{
			GuildID:     queue.GuildID,
			ChannelID:   queue.ChannelID,
			MessageID:   queue.ID,
			SongList:    total,
			PageRange:   pageIndexs,
			CurrentPage: 0,
		})

		// Adding arrow reactions
		S.discordSession.MessageReactionAdd(queue.ChannelID, queue.ID, "⬅")
		S.discordSession.MessageReactionAdd(queue.ChannelID, queue.ID, "➡")
	}

}
