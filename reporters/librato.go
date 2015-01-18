package reporters

import (
	"os"
	"time"

	"github.com/rcrowley/go-librato"
	"github.com/ryanlower/drain/parser"
)

// Librato is a reporter ...
type Librato struct {
	Reporter
	queue []*parser.ParsedLogLine
}

func (l *Librato) init() {
	go l.ticker()
}

// Report ...
func (l *Librato) Report(hit *parser.ParsedLogLine) {
	l.queue = append(l.queue, hit)
}

func (l *Librato) connect(source string) librato.Metrics {
	email := os.Getenv("LIBRATO_EMAIL")
	token := os.Getenv("LIBRATO_TOKEN")

	return librato.NewMetrics(email, token, source)
}

func (l *Librato) ticker() {
	for now := range time.Tick(time.Minute) {
		l.tick(now)
	}
}

func (l *Librato) tick(now time.Time) {
	pending := make([]*parser.ParsedLogLine, len(l.queue))
	if len(l.queue) != 0 {
		// Copy queue into pending, and empty queue
		copy(pending, l.queue)
		l.queue = make([]*parser.ParsedLogLine, 0)
	}

	l.pushStats(pending)
}

func (l *Librato) pushStats(hits []*parser.ParsedLogLine) {
	// TODO, push other stats (connect time, service time etc)
	count := map[string]int{}
	for _, hit := range hits {
		count[hit.Status]++
	}
	for status, n := range count {
		l.pushGauge("drain.statuses", status, n)
	}
}

func (l *Librato) pushGauge(name, source string, count int) {
	metrics := l.connect(source)
	defer metrics.Wait()
	defer metrics.Close()

	gauge := metrics.GetGauge(name)
	gauge <- int64(count)
}
