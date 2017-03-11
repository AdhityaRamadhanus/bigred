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

func ParseRequest(conn io.ReadCloser) (*Commands, error) {
	reader := bufio.NewReader(conn)
	requestStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.Wrap(err, "can't parse new line")
	}
	if requestStr[0] == '*' {
		var args []string
		var argsCount int
		_, err := fmt.Sscanf(requestStr, "*%d", &argsCount)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse array length")
		}
		for i := 0; i <= argsCount-1; i++ {
			commandBytes, err := parseBulkString(reader)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse Resp String")
			}
			args = append(args, string(commandBytes))
		}
		return &Commands{
			Name: strings.ToLower(args[0]),
			Args: args[1:],
		}, nil
	}
	return nil, errors.New("Not a bulk strings")
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
