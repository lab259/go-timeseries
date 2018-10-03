package timesrs

type StorageEntry struct {
	Event         Event
	Collection    string
	Granularities []GranularityComposite
	Result        *PipelineResult
	Keys          map[string]interface{}
}

// Storage will persists the `Result` obtained by a bucket.
type Storage interface {
	Store(entry *StorageEntry) (error)
}
