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

	// IsPlayingSound : Sounds Player state
	IsPlayingSound = map[string]bool{}

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

	// LoadedRadios contain direct links readable by ffmpeg for streams
	// map[radio_name]direct_URL
	// Arrays should have 3 entries: [0] -> CleannedURL, [1] -> Youtube name of the Stream, [2] -> Embed URL+NAME
	LoadedRadios = map[string][]string{}
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
