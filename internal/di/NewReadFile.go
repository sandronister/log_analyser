package di

import (
	"github.com/sandronister/log_analyser/internal/infra/parser"
	"github.com/sandronister/log_analyser/internal/usecase"
)

func NewUseCaseReadFile() (*usecase.ReadFileUseCase, error) {
	parser := parser.NewApachecommonParser()
	readFileUC := usecase.NewReadFileUseCase(parser)
	return readFileUC, nil
}
