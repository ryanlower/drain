package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ryanlower/drain/parser"
	"github.com/ryanlower/drain/reporters"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/drain", drainHandler)

	log.Printf("Listening on port %v ...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
	}
}

func drainHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	parsed, err := parser.Parse(body)

	report(parsed)

	w.WriteHeader(http.StatusOK)
}

func report(hit *parser.ParsedLogLine) {
	// use Log and Redis reporters by default
	// TODO, allow customization
	new(reporters.Log).Report(hit)
	new(reporters.Redis).Report(hit)
}
