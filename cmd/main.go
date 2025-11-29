package main

import (
	"fmt"

	"github.com/sandronister/log_analyser/config"
	"github.com/sandronister/log_analyser/internal/di"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	readFile, err := di.NewUseCaseReadFile()

	if err != nil {
		panic(err)
	}

	stats, err := readFile.Read(config.FolderPath)

	if err != nil {
		panic(err)
	}

	fmt.Println("================= Resumo do Log ==============================")

	fmt.Printf("Total de linhas processadas: %d\n", stats.TotalLines)
	fmt.Printf("Total de erros encontrados: %d\n", stats.TotalErrors)

	fmt.Println("\nContagem de status HTTP:")
	for status, count := range stats.StatusCounts {
		fmt.Printf("Status %d: %d\n", status, count)
	}

	fmt.Println("\nContagem por IP:")
	for ip, count := range stats.IPCounts {
		fmt.Printf("IP %s: %d\n", ip, count)
	}

	fmt.Println("\nContagem por caminho:")
	for path, count := range stats.PathCounts {
		fmt.Printf("Caminho %s: %d\n", path, count)
	}

	fmt.Println("==============================================================")
}
