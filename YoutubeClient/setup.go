package youtubeclient

const missingClientSecretsMessage = `
Please configure OAuth 2.0
`

// SongsQueues : Current Song queues sorted by server ID.
var SongsQueues = map[string][]string{}

var StopPlayerChans = map[string]chan bool{}

var IsPlaying = map[string]bool{}
