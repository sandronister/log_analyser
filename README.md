# Log Analyser

An efficient HTTP log analyzer written in Go that processes log files in batches and generates detailed statistics about web server traffic.

## 📋 Funcionalidades

- **Apache log analysis**: Processes logs in the Apache Common Log Format
- **Comprehensive statistics**: Collects data about HTTP status codes, IPs, endpoints, and errors
- **Configurable processing**: Configure everything via environment variables or a `.env` file
- **Recursive reading**: Automatically processes every file in a folder and its subfolders
- **Clean architecture**: Implements Clean Architecture with dependency injection
- **Optimized performance**: Supports batch processing with configurable workers

## 🚀 Começando

### Prerequisites

- Go 1.24.1 or newer
- Log files in Apache Common Log Format

### Installation

1. Clone the repository:
```bash
git clone https://github.com/sandronister/log_analyser.git
cd log_analyser
```

2. Download the dependencies:
```bash
go mod download
```

### Configuration

The project includes a pre-configured `.env` file with default values. You can modify it as needed or set the variables directly in your system.

#### Available environment variables:

##### Required:
- `FOLDER_PATH`: Path to the folder containing the log files


#### Included `.env` file:

```env
# Server settings
FOLDER_PATH=log_files
```

### How to use:

1. **Basic configuration**: The project already comes with default settings in the `.env` file

2. **Prepare logs**: Place your log files in the `log_files/` directory or update `FOLDER_PATH` in `.env`

3. **Run the analysis**:
```bash
# Direct execution
go run cmd/main.go

# Or build and run
go build -o log-analyser cmd/main.go
./log-analyser
```

4. **Customize settings**: Edit the `.env` file as needed:
```bash
# Example for logs in another directory
FOLDER_PATH=/var/log/apache2
BATCH_SIZE=2000
WORKER_COUNT=8
```

## 📊 Saída

The program outputs a detailed report with the following information:

- **Total lines processed**: Total number of log entries
- **Total errors found**: Count of HTTP status codes >= 400
- **HTTP status count**: Distribution per status code
- **Count per IP**: Request frequency per IP address
- **Count per path**: Distribution of hits by endpoint/path

### Example output:

```
================= Log Summary ================================
Total lines processed: 15420
Total errors found: 234

HTTP status count:
Status 200: 12500
Status 404: 150
Status 500: 84
Status 302: 2686

Count per IP:
IP 192.168.1.1: 450
IP 10.0.0.1: 320
IP 203.0.113.0: 280

Count per path:
Path /: 5600
Path /api/users: 2300
Path /static/style.css: 1800
==============================================================
```

## 🏗️ Arquitetura

The project follows Clean Architecture principles with dependency injection:

```
.env                    # Environment settings

cmd/                    # Application entry point
├── main.go            # Application bootstrap

config/                 # Configuration
├── viper_config.go    # Configuration management with Viper

internal/              
├── di/                # Dependency injection
│   └── NewReadFile.go # Factory for use cases
├── entity/            # Domain entities
│   ├── log_entry.go   # Log entry structure
│   └── stats.go       # Statistics and KV structure
├── infra/             # Infrastructure layer
│   ├── fs/            # File system
│   │   └── file_reader.go  # Recursive directory reader
│   └── parser/        # Log parsers
│       └── apache_common.go # Apache Common format parser
├── ports/             # Interfaces/ports
│   └── parser.go      # Parser interface
└── usecase/           # Use cases/business rules
    └── read_file.go   # Log analysis logic

log_files/             # Directory with log files
├── teste.log          # Sample file
```

## 📝 Formato de Log Suportado

The analyzer supports the Apache Common Log Format:

```
127.0.0.1 - - [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326
```

Components:
- **Client IP**: IP address of the request
- **Timestamp**: Date and time of the request
- **HTTP method**: GET, POST, PUT, etc.
- **Path**: URL/endpoint accessed
- **Status code**: HTTP response (200, 404, 500, etc.)
- **Size**: Bytes transferred

## 🛠️ Tecnologias Utilizadas

- **Go 1.24.1**: Main programming language
- **Viper**: Advanced configuration and environment variable management
- **Clean Architecture**: Layered architectural pattern
- **Apache Common Log Parser**: Specialized parser with optimized regex
- **Dependency Injection**: Pattern for flexibility and testability
- **File System**: Recursive reading and batch processing

## 🤝 Contribuindo

1. Fork this repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 Licença

This project is licensed under the MIT License. See the `LICENSE` file for more details.

## 👨‍💻 Autor

**Sandro Nister**
- GitHub: [@sandronister](https://github.com/sandronister)

---

⭐ If this project was useful to you, consider leaving a star!
