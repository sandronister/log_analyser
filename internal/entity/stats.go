package entity

type Stats struct {
	TotalLines   int
	TotalErrors  int
	StatusCounts map[int]int
	IPCounts     map[string]int
	PathCounts   map[string]int
}

type KV struct {
	Key   string
	Value int
}
