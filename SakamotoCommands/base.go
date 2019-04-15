package sakamotocommands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Sakamoto : Base struct
type Sakamoto struct {
	commandList          map[string]interface{}
	discordSession       *discordgo.Session
	discordMessageCreate *discordgo.MessageCreate
	voiceConn            *discordgo.VoiceConnection
}

// Start : base setup
func Start(s *discordgo.Session, m *discordgo.MessageCreate) Sakamoto {
	S := Sakamoto{}
	S.discordSession = s
	S.discordMessageCreate = m
	S.commandList = map[string]interface{}{
		"join": func(args []string) { S.joinVoice(args) },
		"play": func(args []string) { S.play(args) },
	}
	return S
}

// Execute : execute a given command and args
func (S *Sakamoto) Execute(commandInput string) {
	commandList := strings.Split(commandInput, " ")
	args := []string{}
	if len(commandList[1:]) > 0 {
		args = commandList[1:]
	}
	if command, ok := S.commandList[commandList[0]]; ok {
		print(len(commandInput[1:]))
		command.(func([]string))(args)
	}
}
