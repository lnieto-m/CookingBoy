package imgurgo

import (
	"bytes"
	"log"
	"os/exec"
)

// GetImage calls the Imgur app using a python script and returns a link to that image
func GetImage(sort string, window string, query string) string {

	cmd := exec.Command("python3", "ImgurGo/imgur.py", sort, window, query)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		log.Println(err.Error(), stderr.String())
		return ""
	}
	return string(out)
}
