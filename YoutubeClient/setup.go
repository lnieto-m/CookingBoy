package youtubeclient

var (
	// SongsQueues : Current Song queues sorted by server ID.
	SongsQueues = map[string][]Video{}

	// VoiceConnexions : "GuildID" : "VoiceChannel"
	VoiceConnexions = map[string]string{}

	// StopPlayerChans : Store differnt channels for servers
	StopPlayerChans = map[string]chan bool{}

	// IsPlaying : Player state
	IsPlaying = map[string]bool{}

	// NowPlaying : link to current song playing or empty if no song
	NowPlaying = Video{}

	// QueueMessageCache stores the last queue message for pagination
	QueueMessageCache = []*QueueMessage{}

	PauseChan = map[string]chan bool{}

	PauseStates = map[string]bool{}

	NowPlayingChan = make(chan string, 1)
)

// Video stores data from Youtube API to limit API requests :^)
type Video struct {
	URL               string
	ThumbnailImageURL string
	Title             string
	ID                string
}

type QueueMessage struct {
	GuildID     string
	ChannelID   string
	MessageID   string
	SongList    []string
	PageRange   [][2]int
	CurrentPage int
}
