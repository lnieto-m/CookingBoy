package main

import (
	sakamotocommands "CookingBoy/SakamotoCommands"
	youtubeclient "CookingBoy/YoutubeClient"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func reactionsHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	// Get the message where the reactions were added
	message, err := s.ChannelMessage(m.ChannelID, m.MessageID)
	if err != nil {
		log.Println(err)
		return
	}

	// Only handle other users reactions
	// prevent the auto remove when adding the base reactions to a queue message
	if message.Author.ID != s.State.User.ID || m.UserID == s.State.User.ID {
		return
	}

	// Check if can find queue message in cache
	youtubeclient.ManageQueuePage(s, m)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	skmt := sakamotocommands.Start(s, m)

	// Ignore self message
	if m.Author.ID == s.State.User.ID {
		return
	}

	// 's!' base prefix for commands
	if strings.HasPrefix(m.Content, "s!") {
		skmt.Execute(m.Content[2:])
	}
}

func loadRadio() {
	ready, readyCount, readyTotal := make(chan bool, 1), 0, 0

	// Load json containing radio name and their ID
	// Stored in youtubeclient.LoadedRadios
	data, err := ioutil.ReadFile("radio.json")
	if err != nil {
		log.Println(err.Error())
		return
	}
	radioMap := make(map[string]string)
	err = json.Unmarshal(data, &radioMap)
	if err != nil {
		log.Println(err.Error())
		return
	}
	for key, value := range radioMap {
		go youtubeclient.GetLinksFromScript(key, value, ready)
		readyTotal++
	}

	// Wait for radio initialization to be complete
linksLoop:
	for {
		select {
		case <-ready:
			readyCount++
			if readyCount == readyTotal {
				break linksLoop
			}
		}
	}
	close(ready)
}

func main() {
	discord, err := discordgo.New("Bot NTY2NTUyNDIzOTcyMzM5NzIy.XLGtzA.ktqWKJ6dWudmgiioNT2J_dvpQH8")
	if err != nil {
		log.Println("Error creating Discord session, ", err)
		return
	}

	log.Print("Loading radios...")
	loadRadio()
	log.Println("Done.")

	discord.AddHandler(messageCreate)
	discord.AddHandler(reactionsHandler)

	err = discord.Open()
	if err != nil {
		log.Println("Error opening connection, ", err)
		return
	}

	stopNowPlaying := make(chan bool, 1)

	// Start the routine handling the 'Playing ...' display message
	go sakamotocommands.UpdateGameStatus(discord, stopNowPlaying)

	log.Println("Sakamoto at your service.")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill, syscall.SIGSEGV)
	<-signals
	stopNowPlaying <- true

	discord.Close()
	close(stopNowPlaying)
	log.Println("See you later.")
}
