package sakamotocommands

import youtubeclient "CookingBoy/YoutubeClient"

func (S *Sakamoto) volume(args []string) {
	// One day i'll find Da Wae
	// S.discordSession.VoiceConnections[]
}

func (S *Sakamoto) pause(args []string) {
	if youtubeclient.IsPlaying[S.discordMessageCreate.GuildID] {
		if youtubeclient.PauseStates[S.discordMessageCreate.GuildID] {
			youtubeclient.PauseChan[S.discordMessageCreate.GuildID] <- false
			youtubeclient.PauseStates[S.discordMessageCreate.GuildID] = false
		} else {
			youtubeclient.PauseChan[S.discordMessageCreate.GuildID] <- true
			youtubeclient.PauseStates[S.discordMessageCreate.GuildID] = true
		}
	}

}
