package youtubeclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func getLinksFromScript(name string, ID string, ready chan bool) {
	run := exec.Command("sh", "getDirectStreamLink.sh", ID)
	var stderr bytes.Buffer
	run.Stderr = &stderr
	out, err := run.Output()
	if err != nil {
		log.Println(err, "\n", stderr.String())
		return
	}
	cleanedURL := strings.Trim(string(out), "\n")

	serv, err := YoutubeStart()
	if err != nil {
		log.Println(err.Error())
		return
	}
	video, err := GetVideoInfos(serv, "snippet", ID)
	if err != nil {
		log.Println(err.Error())
		return
	}

	LoadedRadios[name] = []string{cleanedURL, video.Title, "[" + video.Title + "](" + video.URL + ")"}
	ready <- true
}

func GetRadioLinks(pathToJSON string, ready chan bool) {
	data, err := ioutil.ReadFile(pathToJSON)
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
		go getLinksFromScript(key, value, ready)
	}
}
