package reporters

import (
	"encoding/json"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/ryanlower/drain/parser"
)

// Redis is a reporter that publishes hit statues to redis
// The address of the redis server to be published can be
// controlled using the REDIS_ADDRESS env (defaults to "localhost:6379")
// Setting the optional REDIS_PASSWORD env will cause the client to AUTH
// using that password on connection
type Redis struct {
	Reporter
	connection redis.Conn
}

func (r *Redis) init() {
	r.connect()
}

// Report publishes the parsed log lines status code to redis
func (r *Redis) Report(hit *parser.ParsedLogLine) {
	hitJSON, err := json.Marshal(hit)
	if err != nil {
		log.Panic(err)
	} else {
		r.connection.Do("PUBLISH", "drain.hits", hitJSON)
	}
}

func (r *Redis) connect() {
	address := envOrDefault("REDIS_ADDRESS", "localhost:6379")
	conn, err := redis.Dial("tcp", address)

	// TODO, how to handle errors in reporters?
	if err != nil {
		log.Panic(err)
	}

	// AUTH if REDIS_PASSWORD env is set
	if password := os.Getenv("REDIS_PASSWORD"); password != "" {
		conn.Do("AUTH", password)
	}

	r.connection = conn
}
