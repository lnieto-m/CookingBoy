package sakamotocommands

import (
	dgvoice "CookingBoy/DiscordVoice"
	"time"
)

func (S *Sakamoto) soundBox(args []string) {
	if len(args) <= 0 {
		return
	}
	if path, ok := Soundbox[args[0]]; ok {
		S.getVoiceConn()
		S.pause(nil)
		time.Sleep(750 * time.Millisecond)
		dgvoice.PlayAudioFile(S.voiceConn, path, nil, nil)
		time.Sleep(250 * time.Millisecond)
		S.pause(nil)
		return
	}
	S.displaySoundBoxHelp()
}
