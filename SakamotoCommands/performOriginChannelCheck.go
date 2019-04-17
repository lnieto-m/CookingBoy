package sakamotocommands

import (
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
)

func (S *Sakamoto) performOriginChannelCheck() bool {

	// Perform Voice Channel connexion test
	voiceID := ""
	log.Println("performOriginChannelCheck :", youtubeclient.VoiceConnexions[S.discordMessageCreate.GuildID])
	if voiceID, ok := youtubeclient.VoiceConnexions[S.discordMessageCreate.GuildID]; ok == false {
		if voiceID == "" {
		}
		return false
	}
	log.Println(voiceID)

	// Check Guild(Server) existence
	guild, err := S.discordSession.State.Guild(S.discordMessageCreate.GuildID)
	if err != nil {
		log.Println(err)
		return false
	}

	// Check if command's user is in the bot voice channel
	for _, voiceState := range guild.VoiceStates {
		if voiceState.UserID == S.discordMessageCreate.Author.ID && voiceState.ChannelID == youtubeclient.VoiceConnexions[S.discordMessageCreate.GuildID] {
			log.Println("performOriginCheck", voiceState.UserID, S.discordMessageCreate.Author.ID, voiceState.ChannelID)
			return true
		}
	}
	return false
}
