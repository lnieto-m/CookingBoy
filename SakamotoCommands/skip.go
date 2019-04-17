package sakamotocommands

import youtubeclient "CookingBoy/YoutubeClient"

func (S *Sakamoto) skip(args []string) {
	if S.performOriginChannelCheck() == false {
		S.discordSession.ChannelMessageSend(S.discordMessageCreate.ChannelID, "You must be in my voice channel to use this command.")
		return
	}
	if _, ok := youtubeclient.StopPlayerChans[S.discordMessageCreate.GuildID]; ok {
		youtubeclient.StopPlayerChans[S.discordMessageCreate.GuildID] <- true
	}
}
