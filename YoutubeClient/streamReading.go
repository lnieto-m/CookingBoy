package youtubeclient

import (
	dgvoice "CookingBoy/DiscordVoice"

	"github.com/bwmarrin/discordgo"
)

// On testing ,weird behavior now using a pre defined url
func GetWaifuStream(vc *discordgo.VoiceConnection) {
	dgvoice.PlayAudioFile(vc, "https://manifest.googlevideo.com/api/manifest/hls_playlist/id/92BsujVPeA0.0/itag/95/source/yt_live_broadcast/requiressl/yes/ratebypass/yes/live/1/cmbypass/yes/goi/160/sgoap/gir%3Dyes%3Bitag%3D140/sgovp/gir%3Dyes%3Bitag%3D136/hls_chunk_host/r5---sn-25glen7y.googlevideo.com/playlist_type/DVR/gcr/fr/ei/36i8XNmvI9Hi1gaQ5aK4Bg/initcwndbps/7510/mm/32/mn/sn-25glen7y/ms/lv/mv/m/pl/16/dover/11/keepalive/yes/mt/1555867828/disable_polymer/true/ip/62.210.33.226/ipbits/0/expire/1555889471/sparams/ip,ipbits,expire,id,itag,source,requiressl,ratebypass,live,cmbypass,goi,sgoap,sgovp,hls_chunk_host,playlist_type,gcr,ei,initcwndbps,mm,mn,ms,mv,pl/signature/56C4C6D82C33D92C17A6453D22AD5D44417461C7.5EA137FF15914D9139477D962599DE9B76F18219/key/dg_yt0/playlist/index.m3u8", StopPlayerChans[vc.GuildID], nil)

	// run := exec.Command("sh", "waifuRadio.sh")
	// ffmpegout, err := run.StdoutPipe()
	// if err != nil {
	// 	// OnError("StdoutPipe Error", err)
	// 	return
	// }

	// ffmpegbuf := bufio.NewReaderSize(ffmpegout, 16384)

	// // Starts the ffmpeg command
	// err = run.Start()
	// if err != nil {
	// 	// OnError("RunStart Error", err)
	// 	return
	// }

	// lol := make(chan []int16, 2)
	// // close := make(chan bool)

	// go func() {
	// 	dgvoice.SendPCM(vc, lol)
	// }()

	// for {
	// 	log.Println("reading...")
	// 	// read data from ffmpeg stdout
	// 	audiobuf := make([]int16, 960*2)
	// 	err = binary.Read(ffmpegbuf, binary.LittleEndian, &audiobuf)
	// 	if err == io.EOF || err == io.ErrUnexpectedEOF {
	// 		log.Println(err)
	// 		return
	// 	}
	// 	if err != nil {
	// 		// OnError("error reading from ffmpeg stdout", err)
	// 		return
	// 	}

	// 	// Send received PCM to the sendPCM channel
	// 	select {
	// 	case lol <- audiobuf:

	// 		// case <-close:
	// 		// return
	// 	}
	// }

}
