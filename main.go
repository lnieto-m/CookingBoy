package main

import (
	sakamotocommands "CookingBoy/SakamotoCommands"
	youtubeclient "CookingBoy/YoutubeClient"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func reactionsHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	message, err := s.ChannelMessage(m.ChannelID, m.MessageID)
	if err != nil {
		log.Println(err)
		return
	}
	if message.Author.ID != s.State.User.ID || m.UserID == s.State.User.ID {
		return
	}
	// Check if can find queue message in cache
	// s.State.User.ID != m.UserID
	youtubeclient.ManageQueuePage(s, m)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	skmt := sakamotocommands.Start(s, m)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "s!") {
		skmt.Execute(m.Content[2:])
	}
}

func main() {
	discord, err := discordgo.New("Bot NTY2NTUyNDIzOTcyMzM5NzIy.XLGtzA.ktqWKJ6dWudmgiioNT2J_dvpQH8")
	if err != nil {
		log.Println("Error creating Discord session, ", err)
		return
	}

	discord.AddHandler(messageCreate)
	discord.AddHandler(reactionsHandler)

	err = discord.Open()
	if err != nil {
		log.Println("Error opening connection, ", err)
		return
	}

	log.Println("Sakamoto at your service.")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill, syscall.SIGSEGV)
	<-signals

	discord.Close()
	log.Println("See you later.")
}
