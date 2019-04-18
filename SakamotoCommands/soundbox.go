package sakamotocommands

import (
	"github.com/bwmarrin/dgvoice"
)

func (S *Sakamoto) displaySoundBoxHelp() {

}

func (S *Sakamoto) soundBox(args []string) {
	if len(args) <= 0 {
		return
	}
	if path, ok := Soundbox[args[0]]; ok {
		S.getVoiceConn()
		// youtubeclient.PauseChan <- true
		dgvoice.PlayAudioFile(S.voiceConn, path, nil)
		return
	}
	S.displaySoundBoxHelp()
}
