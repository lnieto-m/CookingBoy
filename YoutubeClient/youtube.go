package youtubeclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

// youtube.go : Base go code to use the youtube API
// Added some custom methods to retrieve videos and playlist data

const missingClientSecretsMessage = `
Please configure OAuth 2.0
`

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	tokenCacheDir := ".credentials"
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("youtube-go-quickstart.json")), nil
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

// GetVideoInfos use a youtube id to return a Video object
func GetVideoInfos(service *youtube.Service, part string, id string) (Video, error) {
	video := Video{}
	call := service.Videos.List(part)
	call = call.Id(id)
	response, err := call.Do()
	if err != nil {
		return video, err
	}
	for _, item := range response.Items {
		video.ID = id
		video.Title = item.Snippet.Title
		video.URL = "https://www.youtube.com/watch?v=" + id
		thumbnail := ""
		if item.Snippet != nil && item.Snippet.Thumbnails != nil && item.Snippet.Thumbnails.Standard != nil {
			thumbnail = item.Snippet.Thumbnails.Standard.Url
		}
		video.ThumbnailImageURL = thumbnail
		return video, err
	}
	return video, err
}

// GetPlaylistVideos use a playlist id to return a list of Video objects
func GetPlaylistVideos(service *youtube.Service, part string, id string) []Video {
	pageToken := ""
	hasEntered := 0
	idLists := []Video{}
	for {
		call := service.PlaylistItems.List(part)
		call = call.PlaylistId(id)
		call = call.MaxResults(50)
		if hasEntered != 0 {
			call = call.PageToken(pageToken)
		}
		response, err := call.Do()
		if err != nil {
			log.Println(err)
			return idLists
		}
		pageToken = response.NextPageToken
		hasEntered = 1
		for _, items := range response.Items {
			thumbnail := ""
			if items.Snippet != nil && items.Snippet.Thumbnails != nil && items.Snippet.Thumbnails.Standard != nil {
				thumbnail = items.Snippet.Thumbnails.Standard.Url
			}
			video := Video{
				URL:               "https://www.youtube.com/watch?v=" + items.ContentDetails.VideoId,
				ThumbnailImageURL: thumbnail,
				Title:             items.Snippet.Title,
				ID:                items.ContentDetails.VideoId,
			}
			idLists = append(idLists, video)
		}
		if pageToken == "" {
			return idLists
		}
	}
}

// YoutubeStart return a new youtube.service using given credentials
func YoutubeStart() (*youtube.Service, error) {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/youtube-go-quickstart.json
	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)
	service, err := youtube.New(client)

	handleError(err, "Error creating YouTube client")

	return service, err
}
