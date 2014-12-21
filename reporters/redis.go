package reporters

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/ryanlower/drain/parser"
)

// Redis is a reporter that publishes hit statues to redis
// The address of the redis server to be published can be
// controlled using the REDIS_ADDRESS env (defaults to "localhost:6379")
type Redis struct {
	Reporter
}

// Report publishes the parsed log lines status code to redis
func (r *Redis) Report(hit *parser.ParsedLogLine) {
	conn := connect()

	conn.Do("PUBLISH", "drain.statuses", hit.Status)
}

func connect() redis.Conn {
	address := envOrDefault("REDIS_ADDRESS", "localhost:6379")
	// TODO, auth handling
	conn, err := redis.Dial("tcp", address)
	// TODO, how to handle errors in reporters?
	if err != nil {
		log.Panic(err)
	}

	return conn
}
