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

	// PauseChan allow the communication between the user and the ffmpeg process
	// send True to halt the process and false to resume it
	PauseChan = map[string]chan bool{}

	// PauseStates stores the states of pause
	// True if the song player is paused else false
	PauseStates = map[string]bool{}

	// NowPlayingChan receive the song currently playing to display it in the 'Playing ... ' message of a discord user
	// This channel is binded to the sakamotocommands.UpdateGameStatus method and should be used only by this function
	NowPlayingChan = make(chan string, 1)
)

// Video stores data from Youtube API to limit API requests :^)
type Video struct {
	URL               string
	ThumbnailImageURL string
	Title             string
	ID                string
}

// QueueMessage stores data for a queue message to allow futur modifications through arrow emojis reactions
type QueueMessage struct {
	GuildID     string
	ChannelID   string
	MessageID   string
	SongList    []string
	PageRange   [][2]int
	CurrentPage int
}
