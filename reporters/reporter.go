package reporters

import (
	"os"

	"github.com/ryanlower/drain/parser"
)

// Reporter interface
// To be a reporter, you must implement Report, taking a parser.ParsedLogLine
type Reporter interface {
	Report(hit *parser.ParsedLogLine)
}

func envOrDefault(key string, defaultValue string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}
	return defaultValue
}
