package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/sandronister/log_analyser/internal/entity"
	"github.com/sandronister/log_analyser/internal/ports"
)

var regex = `^(\S+) \S+ \S+ \[([^\]]+)\] "(\S+) (\S+) \S+" (\d{3}) (\d+|-)`

type ApacheCommonParser struct {
	logRegexp *regexp.Regexp
}

func (p *ApacheCommonParser) ParserFile(filepath string) ([]entity.LogEntry, error) {
	f, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	return p.scanFile(scanner)
}

func NewApachecommonParser() ports.Parser {
	return &ApacheCommonParser{
		logRegexp: regexp.MustCompile(regex),
	}
}

func convertTimestamp(dateStr string) (int64, error) {
	layout := "02/Jan/2006:15:04:05 -0700"

	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0, err
	}

	return t.Unix(), nil
}

func (p *ApacheCommonParser) parseLine(line string) (*entity.LogEntry, error) {
	matches := p.logRegexp.FindStringSubmatch(line)

	if len(matches) < 7 {
		return nil, nil
	}

	status, err := strconv.Atoi(matches[5])
	if err != nil {
		return nil, err
	}

	size := 0
	if matches[6] != "-" {
		size, err = strconv.Atoi(matches[6])
		if err != nil {
			return nil, err
		}
	}

	timestamp, err := convertTimestamp(matches[2])
	if err != nil {
		return nil, err
	}

	return &entity.LogEntry{
		IP:         matches[1],
		Timestamp:  timestamp,
		Method:     matches[3],
		Path:       matches[4],
		StatusCode: status,
		Size:       int64(size),
	}, nil

}

func (p *ApacheCommonParser) scanFile(scanner *bufio.Scanner) ([]entity.LogEntry, error) {

	var entries []entity.LogEntry
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		entry, err := p.parseLine(line)
		if err != nil {
			return nil, err
		}

		entries = append(entries, *entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil

}
