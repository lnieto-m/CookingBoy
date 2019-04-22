package youtubeclient

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

// GetLinksFromScript retrieves the direct URL for a given stream ID using the getDirectStreamLink.sh script
func GetLinksFromScript(name string, ID string, ready chan bool) {
	run := exec.Command("sh", "getDirectStreamLink.sh", ID)
	var stderr bytes.Buffer
	run.Stderr = &stderr
	out, err := run.Output()
	if err != nil {
		log.Println(err, "\n", stderr.String())
		ready <- true
		return
	}
	cleanedURL := strings.Trim(string(out), "\n")

	serv, err := YoutubeStart()
	if err != nil {
		log.Println(err.Error())
		ready <- true
		return
	}
	video, err := GetVideoInfos(serv, "snippet", ID)
	if err != nil {
		log.Println(err.Error())
		ready <- true
		return
	}

	LoadedRadios[name] = []string{cleanedURL, video.Title, "[" + video.Title + "](" + video.URL + ")"}
	ready <- true
}
