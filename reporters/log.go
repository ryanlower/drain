package reporters

import (
	"log"

	"github.com/ryanlower/drain/parser"
)

// Log is a reporter that logs via Report
type Log struct{}

// Report logs the parsed log line
func (r *Log) Report(hit *parser.ParsedLogLine) {
	log.Print(hit)
}
