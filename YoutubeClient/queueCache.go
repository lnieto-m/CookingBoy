package youtubeclient

func PushToQueueCache(songList QueueMessage) {
	if len(QueueMessageCache) > 5 {
		QueueMessageCache = QueueMessageCache[1:]
		QueueMessageCache = append(QueueMessageCache, songList)
		return
	}
	QueueMessageCache = append(QueueMessageCache, songList)
}
