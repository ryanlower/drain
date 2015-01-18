package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ryanlower/drain/parser"
	"github.com/ryanlower/drain/reporters"
)

// Drain is used to maintain a list of Reporters
// registered to recieve parsed logs for processing
type Drain struct {
	reporters []reporters.Reporter
}

// AddReporter adds a reporter of type t
// to the list of drain reporters
func (d *Drain) AddReporter(t string) error {
	reporter, err := reporters.New(t)
	if err != nil {
		panic(err)
	}

	d.reporters = append(d.reporters, reporter)

	return nil
}

// Handler takes a http.Request,
// checks authorization via HTTP basic auth (if setup, see authenticated)
// parses the request into a parser.ParsedLogLine via parser.Parse
// and sends the ParsedLogLine to registered Reporters
// If all goes well, writes an OK status to the http.ResponseWriter
func (d *Drain) Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if !authenticated(r) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// TODO, don't break if no body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	parsed, err := parser.Parse(body)

	if parsed != nil {
		d.report(parsed)
	}

	w.WriteHeader(http.StatusOK)
}

func (d *Drain) report(hit *parser.ParsedLogLine) {
	for _, reporter := range d.reporters {
		go reporter.Report(hit)
	}
}

// Listens for logs at /drain on env PORT
func main() {
	port := os.Getenv("PORT")

	drain := new(Drain)
	// use Log and Redis reporters by default
	// TODO, allow customization
	drain.AddReporter("log")
	drain.AddReporter("redis")
	drain.AddReporter("librato")

	http.HandleFunc("/drain", drain.Handler)

	log.Printf("Listening on port %v ...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
	}
}

// Helper HTTP basic auth function
// Returns false if AUTH_PASSWORD env is set and provided password doesn't match
// true otherwise
func authenticated(r *http.Request) bool {
	auth := os.Getenv("AUTH_PASSWORD")

	_, password, _ := r.BasicAuth()
	if auth != "" && auth != password {
		// AUTH_PASSWORD is set and provided password doesn't match
		return false
	}
	return true
}
