package youtubeclient

// SongsQueues : Current Song queues sorted by server ID.
var SongsQueues = map[string][]Video{}

// VoiceConnexions : "GuildID" : "VoiceChannel"
var VoiceConnexions = map[string]string{}

// StopPlayerChans : Store differnt channels for servers
var StopPlayerChans = map[string]chan bool{}

// IsPlaying : Player state
var IsPlaying = map[string]bool{}

// NowPlaying : link to current song playing or empty if no song
var NowPlaying = Video{}

// QueueMessageCache stores the last queue message for pagination
var QueueMessageCache = []QueueMessage{}

// Video stores data from Youtube API to limit API requests :^)
type Video struct {
	URL               string
	ThumbnailImageURL string
	Title             string
	ID                string
}

type QueueMessage struct {
	GuildID   string
	ChannelID string
	MessageID string
	SongList  []string
}
