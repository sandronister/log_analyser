# Log Analyser

Um analisador de logs HTTP eficiente escrito em Go que processa arquivos de log em lote e gera estatísticas detalhadas sobre o tráfego do servidor web.

## 📋 Funcionalidades

- **Análise de logs Apache**: Processa logs no formato Apache Common Log Format
- **Estatísticas abrangentes**: Coleta dados sobre status HTTP, IPs, endpoints e erros
- **Processamento configurável**: Configuração via variáveis de ambiente ou arquivo `.env`
- **Leitura recursiva**: Processa todos os arquivos em uma pasta e suas subpastas automaticamente
- **Arquitetura limpa**: Implementa Clean Architecture com injeção de dependência
- **Performance otimizada**: Suporte a processamento em lote com workers configuráveis

## 🚀 Começando

### Pré-requisitos

- Go 1.24.1 ou superior
- Arquivos de log no formato Apache Common Log Format

### Instalação

1. Clone o repositório:
```bash
git clone https://github.com/sandronister/log_analyser.git
cd log_analyser
```

2. Baixe as dependências:
```bash
go mod download
```

### Configuração

O projeto inclui um arquivo `.env` pré-configurado com valores padrão. Você pode modificá-lo conforme suas necessidades ou definir as variáveis diretamente no sistema.

#### Variáveis de ambiente disponíveis:

##### Obrigatórias:
- `FOLDER_PATH`: Caminho para a pasta contendo os arquivos de log


#### Arquivo `.env` incluído:

```env
# Configurações do servidor
FOLDER_PATH=log_files
```

### Como usar:

1. **Configuração básica**: O projeto já vem com configurações padrão no arquivo `.env`

2. **Preparar logs**: Coloque seus arquivos de log no diretório `log_files/` ou modifique o `FOLDER_PATH` no `.env`

3. **Executar análise**:
```bash
# Execução direta
go run cmd/main.go

# Ou compilar e executar
go build -o log-analyser cmd/main.go
./log-analyser
```

4. **Personalizar configurações**: Edite o arquivo `.env` conforme necessário:
```bash
# Exemplo para logs em outro diretório
FOLDER_PATH=/var/log/apache2
BATCH_SIZE=2000
WORKER_COUNT=8
```

## 📊 Saída

O programa gera um relatório detalhado com as seguintes informações:

- **Total de linhas processadas**: Número total de entradas de log
- **Total de erros encontrados**: Contagem de códigos de status HTTP >= 400
- **Contagem de status HTTP**: Distribuição por código de status
- **Contagem por IP**: Frequência de requisições por endereço IP
- **Contagem por caminho**: Distribuição de acessos por endpoint/caminho

### Exemplo de saída:

```
================= Resumo do Log ==============================
Total de linhas processadas: 15420
Total de erros encontrados: 234

Contagem de status HTTP:
Status 200: 12500
Status 404: 150
Status 500: 84
Status 302: 2686

Contagem por IP:
IP 192.168.1.1: 450
IP 10.0.0.1: 320
IP 203.0.113.0: 280

Contagem por caminho:
Caminho /: 5600
Caminho /api/users: 2300
Caminho /static/style.css: 1800
==============================================================
```

## 🏗️ Arquitetura

O projeto segue os princípios da Clean Architecture com injeção de dependência:

```
.env                    # Configurações de ambiente

cmd/                    # Ponto de entrada da aplicação
├── main.go            # Bootstrap da aplicação

config/                 # Configurações
├── viper_config.go    # Gerenciamento de configuração com Viper

internal/              
├── di/                # Injeção de dependência
│   └── NewReadFile.go # Factory para casos de uso
├── entity/            # Entidades de domínio
│   ├── log_entry.go   # Estrutura de entrada de log
│   └── stats.go       # Estrutura de estatísticas e KV
├── infra/             # Camada de infraestrutura
│   ├── fs/            # Sistema de arquivos
│   │   └── file_reader.go  # Leitura recursiva de diretórios
│   └── parser/        # Parsers de log
│       └── apache_common.go # Parser para formato Apache Common
├── ports/             # Interfaces/Portas
│   └── parser.go      # Interface para parsers
└── usecase/           # Casos de uso/Regras de negócio
    └── read_file.go   # Lógica de análise de logs

log_files/             # Diretório com arquivos de log
├── teste.log          # Arquivo de exemplo
```

## 📝 Formato de Log Suportado

O analisador suporta o formato Apache Common Log Format:

```
127.0.0.1 - - [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326
```

Componentes:
- **IP do cliente**: Endereço IP da requisição
- **Timestamp**: Data e hora da requisição
- **Método HTTP**: GET, POST, PUT, etc.
- **Caminho**: URL/endpoint acessado
- **Código de status**: Resposta HTTP (200, 404, 500, etc.)
- **Tamanho**: Bytes transferidos

## 🛠️ Tecnologias Utilizadas

- **Go 1.24.1**: Linguagem de programação principal
- **Viper**: Gerenciamento avançado de configuração e variáveis de ambiente
- **Clean Architecture**: Padrão arquitetural com separação de camadas
- **Apache Common Log Parser**: Parser especializado com regex otimizada
- **Injeção de Dependência**: Padrão para flexibilidade e testabilidade
- **Sistema de Arquivos**: Leitura recursiva e processamento em lote

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 👨‍💻 Autor

**Sandro Nister**
- GitHub: [@sandronister](https://github.com/sandronister)

---

⭐ Se este projeto foi útil para você, considere dar uma estrela!