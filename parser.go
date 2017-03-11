package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"

	"strings"

	"github.com/pkg/errors"
)

type Commands struct {
	Name string
	Args []string
}

// Taken from https://github.com/docker/go-redis-server/blob/master/parser.go

// ParseRequest Parsing Request from redis client and return Commands which consist of a command name and its arguments
func ParseRequest(conn io.ReadCloser) (*Commands, error) {
	reader := bufio.NewReader(conn)
	requestStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.Wrap(err, "can't parse new line")
	}
	// Request should come as an array of resp bulk string
	if requestStr[0] == '*' {
		var args []string
		var argsCount int
		_, err := fmt.Sscanf(requestStr, "*%d", &argsCount)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse array length")
		}
		// Get the arguments, it's expected to be a bulk string
		for i := 0; i <= argsCount-1; i++ {
			commandBytes, err := parseBulkString(reader)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse Resp String")
			}
			args = append(args, string(commandBytes))
		}
		return &Commands{
			Name: strings.ToLower(args[0]),
			Args: args[1:], // Totally unsafe
		}, nil
	}
	return nil, errors.New("Cannot parse request")
}

func parseBulkString(r *bufio.Reader) ([]byte, error) {

	line, err := r.ReadString('\n')
	if err != nil {
		return nil, errors.New("Malformed resp string")
	}
	var argSize int
	if _, err := fmt.Sscanf(line, "$%d", &argSize); err != nil {
		return nil, errors.New("Failed to get string length")
	}

	data, err := ioutil.ReadAll(io.LimitReader(r, int64(argSize)))
	if err != nil {
		return nil, err
	}

	if len(data) != argSize {
		return nil, errors.New("Lenght of actual data not same as in meta data")
	}

	if b, err := r.ReadByte(); err != nil || b != '\r' {
		return nil, errors.New("Missing CR")
	}
	if b, err := r.ReadByte(); err != nil || b != '\n' {
		return nil, errors.New("Missing LF")
	}

	return data, nil
}
