package entity

type LogEntry struct {
	IP         string
	Timestamp  int64
	Method     string
	Path       string
	StatusCode int
	Size       int64
}
