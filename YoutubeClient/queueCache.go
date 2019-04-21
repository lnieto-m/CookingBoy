package youtubeclient

// PushToQueueCache : used for queue message pagination
// Stores Song queue messages too long to be displayed in one embed
// Only 5 Messages can be stored
func PushToQueueCache(songList *QueueMessage) {
	if len(QueueMessageCache) > 5 {
		QueueMessageCache = QueueMessageCache[1:]
		QueueMessageCache = append(QueueMessageCache, songList)
		return
	}
	QueueMessageCache = append(QueueMessageCache, songList)
}
