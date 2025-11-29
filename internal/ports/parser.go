package ports

import (
	"github.com/sandronister/log_analyser/internal/entity"
)

type Parser interface {
	ParserFile(filepath string) ([]entity.LogEntry, error)
}
