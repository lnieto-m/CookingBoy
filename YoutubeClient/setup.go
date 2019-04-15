package youtubeclient

// SongsQueues : Current Song queues sorted by server ID.
var SongsQueues = map[string][]string{}

// StopPlayerChans : Store differnt channels for servers
var StopPlayerChans = map[string]chan bool{}

// IsPlaying : Player state
var IsPlaying = map[string]bool{}

// NowPlaying : link to current song playing or empty if no song
var NowPlaying = ""
