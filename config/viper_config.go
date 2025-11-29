package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config representa a estrutura de configuração da aplicação
type Config struct {
	FolderPath string `mapstructure:"FOLDER_PATH"`
}

// LoadConfig carrega as configurações das variáveis de ambiente
func LoadConfig() (*Config, error) {
	config := &Config{}

	// Configura o Viper para ler variáveis de ambiente
	viper.AutomaticEnv()

	// Tenta carregar de um arquivo .env se existir
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/log-analyser/")

	// Lê o arquivo de configuração (opcional)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Aviso: Arquivo de configuração não encontrado: %v", err)
	}

	// Faz o binding da configuração para a struct
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("erro ao fazer unmarshal da configuração: %w", err)
	}

	// Valida a configuração
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("configuração inválida: %w", err)
	}

	return config, nil
}

func validateConfig(config *Config) error {

	if config.FolderPath == "" {
		return fmt.Errorf("FOLDER_PATH não pode ser vazio")
	}

	return nil
}

var appConfig *Config

func Init() error {
	var err error
	appConfig, err = LoadConfig()
	if err != nil {
		return fmt.Errorf("falha ao carregar configuração: %w", err)
	}
	return nil
}

func Get() *Config {
	if appConfig == nil {
		log.Fatal("Configuração não foi inicializada. Chame config.Init() primeiro.")
	}
	return appConfig
}
