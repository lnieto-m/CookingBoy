package main

import (
	sakamotocommands "CookingBoy/SakamotoCommands"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot NTY2NTUyNDIzOTcyMzM5NzIy.XLGtzA.ktqWKJ6dWudmgiioNT2J_dvpQH8")
	if err != nil {
		log.Println("Error creating Discord session, ", err)
		return
	}

	discord.AddHandler(messageCreate)

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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	skmt := sakamotocommands.Start(s, m)

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "s!") {
		skmt.Execute(m.Content[2:])
	}
}
