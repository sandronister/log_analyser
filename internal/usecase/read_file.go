package usecase

import (
	"fmt"

	"github.com/sandronister/log_analyser/internal/entity"
	"github.com/sandronister/log_analyser/internal/infra/fs"
	"github.com/sandronister/log_analyser/internal/ports"
)

type ReadFileUseCase struct {
	parser ports.Parser
}

func NewReadFileUseCase(parser ports.Parser) *ReadFileUseCase {
	return &ReadFileUseCase{
		parser: parser,
	}
}

func (r *ReadFileUseCase) getLogEntries(files []string) ([]entity.LogEntry, error) {
	var logEntries []entity.LogEntry
	for _, file := range files {
		entries, err := r.parser.ParserFile(file)
		if err != nil {
			fmt.Println("Error parsing file:", err)
			return nil, err
		}
		logEntries = append(logEntries, entries...)
	}

	return logEntries, nil
}

func (r *ReadFileUseCase) compute(entries []entity.LogEntry) (*entity.Stats, error) {
	stats := &entity.Stats{
		TotalLines:   len(entries),
		StatusCounts: make(map[int]int),
		IPCounts:     make(map[string]int),
		PathCounts:   make(map[string]int),
	}

	for _, entry := range entries {
		stats.StatusCounts[entry.StatusCode]++
		stats.IPCounts[entry.IP]++
		stats.PathCounts[entry.Path]++
		if entry.StatusCode >= 400 {
			stats.TotalErrors++
		}
	}

	return stats, nil
}

func (r *ReadFileUseCase) Read(folder string) (*entity.Stats, error) {
	files, err := fs.ReadFolderRecursive(folder)

	if err != nil {
		return nil, err
	}

	entries, err := r.getLogEntries(files)

	if err != nil {
		return nil, err
	}

	return r.compute(entries)
}
