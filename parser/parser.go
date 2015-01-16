package parser

import (
	"errors"
	"regexp"
)

// ParsedLogLine represents a log line parsed by Parse
type ParsedLogLine struct {
	Timestamp   string
	Path        string
	Host        string
	IP          string
	ConnectTime string // TODO, make ConnectTime & ServiceTime floats
	ServiceTime string
	Status      string
}

// TODO, match generic key=value data rather than hardcoding specifics
var regex = regexp.MustCompile(`<\d+>\d\s(.+?)\s.+path="(\S+)".+host=(\S+).+fwd="(\S+)".+connect=(\d+)ms.+service=(\d+)ms.+status=(\d+)`)

// Parse parses the provided body, returning a ParsedLogLine
func Parse(body []byte) (*ParsedLogLine, error) {
	if match := regex.FindSubmatch(body); match != nil {
		parsed := &ParsedLogLine{
			Timestamp:   string(match[1]),
			Path:        string(match[2]),
			Host:        string(match[3]),
			IP:          string(match[4]),
			ConnectTime: string(match[5]),
			ServiceTime: string(match[6]),
			Status:      string(match[7]),
		}
		return parsed, nil
	}

	return nil, errors.New("Can't parse: body doesn't match regex")
}
