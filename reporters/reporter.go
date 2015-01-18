package reporters

import (
	"errors"
	"os"

	"github.com/ryanlower/drain/parser"
)

// Reporter interface
// To be a reporter, you must implement Report, taking a parser.ParsedLogLine
type Reporter interface {
	init()
	Report(hit *parser.ParsedLogLine)
}

// New creates a new Reporter of specified type
// and initializes it
func New(t string) (Reporter, error) {
	var reporter Reporter
	switch t {
	case "log":
		reporter = new(Log)
	case "redis":
		reporter = new(Redis)
	case "librato":
		reporter = new(Librato)
	default:
		return nil, errors.New("Unknown reporter type: " + t)
	}

	reporter.init()

	return reporter, nil
}

func envOrDefault(key string, defaultValue string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}
	return defaultValue
}
