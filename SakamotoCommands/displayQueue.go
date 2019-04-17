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

	log.Printf("%v\n", len(message.Fields))
	for _, item := range message.Fields {
		log.Printf("n: %v v: %v%v\n", item.Name, item.Value, len(item.Value))
	}

	queue, err := S.discordSession.ChannelMessageSendEmbed(S.discordMessageCreate.ChannelID, message)
	if err != nil {
		log.Println("displayQueue: ", err.Error())
		return
	}
	// log.Println(to)
	if len(total) > 25 {
		youtubeclient.PushToQueueCache(youtubeclient.QueueMessage{
			GuildID:   queue.GuildID,
			ChannelID: queue.ChannelID,
			MessageID: queue.ID,
			SongList:  total,
		})
		err = S.discordSession.MessageReactionAdd(queue.ChannelID, queue.ID, "⬅")
		if err != nil {
			log.Println("Emoji ", err)
		}
		S.discordSession.MessageReactionAdd(queue.ChannelID, queue.ID, "➡")
	}

}
