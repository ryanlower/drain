package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTimestamp(t *testing.T) {
	body := []byte(`<321>1 2014-12-25T19:07:10.076560+00:00 host heroku router - at=info method=GET path="/test/123" host=example.com request_id=12345-abcde fwd="127.0.0.1" dyno=web.11 connect=12ms service=123ms status=200 bytes=1234`)

	parsed, _ := Parse(body)
	assert.Equal(t, "2014-12-25T19:07:10.076560+00:00", parsed.Timestamp)
}

func TestParsePath(t *testing.T) {
	body := []byte(`<321>1 2014-12-25T19:07:10.076560+00:00 host heroku router - at=info method=GET path="/test/123" host=example.com request_id=12345-abcde fwd="127.0.0.1" dyno=web.11 connect=12ms service=123ms status=200 bytes=1234`)

	parsed, _ := Parse(body)
	assert.Equal(t, "/test/123", parsed.Path)
}

func TestParseHost(t *testing.T) {
	body := []byte(`<321>1 2014-12-25T19:07:10.076560+00:00 host heroku router - at=info method=GET path="/test/123" host=example.com request_id=12345-abcde fwd="127.0.0.1" dyno=web.11 connect=12ms service=123ms status=200 bytes=1234`)

	parsed, _ := Parse(body)
	assert.Equal(t, "example.com", parsed.Host)
}

func TestParseStatus(t *testing.T) {
	body := []byte(`<321>1 2014-12-25T19:07:10.076560+00:00 host heroku router - at=info method=GET path="/test/123" host=example.com request_id=12345-abcde fwd="127.0.0.1" dyno=web.11 connect=12ms service=123ms status=200 bytes=1234`)

	parsed, _ := Parse(body)
	assert.Equal(t, "200", parsed.Status)
}

func TestParseError(t *testing.T) {
	body := []byte(`something else`)

	parsed, err := Parse(body)
	assert.Nil(t, parsed)
	assert.Equal(t, "Can't parse: body doesn't match regex", err.Error())
}
