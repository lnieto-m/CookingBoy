package sakamotocommands

import "github.com/bwmarrin/discordgo"

//
type Sakamoto struct {
	commandList          map[string]interface{}
	discordSession       *discordgo.Session
	discordMessageCreate *discordgo.MessageCreate
}

// Start : base setup
func Start(s *discordgo.Session, m *discordgo.MessageCreate) Sakamoto {
	S := Sakamoto{}
	S.discordSession = s
	S.discordMessageCreate = m
	S.commandList = map[string]interface{}{
		"join": func(args []string) { S.joinVoice(args) },
	}
	return S
}

// Execute : execute a given command and args
func (S *Sakamoto) Execute(commandInput []string) {
	if len(commandInput) > 0 {
		if command, ok := S.commandList[commandInput[0]]; ok {
			print(len(commandInput[1:]))
			command.(func([]string))(commandInput)
		}
	}
}
