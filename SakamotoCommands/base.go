package sakamotocommands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// base.go : Base setup to use the Bot commands

// Sakamoto : Base struct
type Sakamoto struct {
	commandList          map[string]interface{}     // command list [commandName] - > corresponding function
	discordSession       *discordgo.Session         // Current discord session
	discordMessageCreate *discordgo.MessageCreate   // Stored message create event
	voiceConn            *discordgo.VoiceConnection // Stored current voice connexion
}

// Start : base setup
// Stores current session, message event and avalaible commands then return a Sakamoto object
func Start(s *discordgo.Session, m *discordgo.MessageCreate) Sakamoto {
	S := Sakamoto{}
	S.discordSession = s
	S.discordMessageCreate = m
	S.commandList = map[string]interface{}{
		"join":  func(args []string) { S.joinVoice(args) },
		"play":  func(args []string) { S.play(args) },
		"queue": func(args []string) { S.displayQueue(args) },
		"stop":  func(args []string) { S.stop(args) },
		"skip":  func(args []string) { S.skip(args) },
		"leave": func(args []string) { S.leaveVoice(args) },
		"help":  func(args []string) { S.help(args) },
		"sound": func(args []string) { S.soundBox(args) },
		"pause": func(args []string) { S.pause(args) },
	}
	return S
}

// Execute : execute a given command and args
// Inspect the command map to get the corresponding function else does nothing
func (S *Sakamoto) Execute(commandInput string) {
	commandList := strings.Split(commandInput, " ")
	args := []string{}
	if len(commandList[1:]) > 0 {
		args = commandList[1:]
	}
	if command, ok := S.commandList[commandList[0]]; ok {
		command.(func([]string))(args)
	}
}
