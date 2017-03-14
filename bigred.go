package main

import (
	"log"
	"net"
	"time"

	"fmt"
	"os"
	"runtime"

	"path/filepath"

	"github.com/allegro/bigcache"
)

// Bigred is the main service struct consist of cache engine bigcache
type Bigred struct {
	Proto   string
	Addr    string
	Cache   *bigcache.BigCache
	clients int
}

// NewBigRed is Bigred constructor accepting proto and addr
// Ex: NewBigRed("tcp", "localhost:6399")
func NewBigRed(proto, addr string) (*Bigred, error) {
	// TODO : parameterized the default config
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(1 * time.Minute))
	if err != nil {
		return nil, err
	}
	return &Bigred{
		Proto: proto,
		Addr:  addr,
		Cache: cache,
	}, nil
}

// Run is the main loop of Bigred service
func (b *Bigred) Run() error {
	l, err := net.Listen(b.Proto, b.Addr)
	defer l.Close()
	if err != nil {
		return err
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		b.clients++
		log.Println("Accepting Connection", conn.RemoteAddr().String())
		go b.handleClient(conn)
	}
}

func (b *Bigred) handleClient(conn net.Conn) (err error) {
	defer func() {
		if err != nil && err.Error() != "EOF" {
			log.Println(err)
		} else {
			log.Println("Closing Connection")
		}
		conn.Close()
		b.clients--
	}()
	for {
		commands, err := ParseRequest(conn)
		if err != nil {
			return err
		}
		// log.Println(commands.Args, commands.Name)
		if err := b.executeCommand(commands, conn); err != nil {
			return err
		}
	}
	// return nil
}

// Since only a few method will be supported, i figure why not just using switch case
func (b *Bigred) executeCommand(cmd *Commands, conn net.Conn) error {
	switch cmd.Name {
	case "get":
		if len(cmd.Args) < 1 {
			return replyErrArgLength(conn, cmd.Name)
		}
		bytes, err := b.Cache.Get(cmd.Args[0])
		if err != nil || len(bytes) == 0 {
			return replyNil(conn)
		}
		return replyBulkString(conn, bytes)
	case "set":
		if len(cmd.Args) < 2 {
			return replyErrArgLength(conn, cmd.Name)
		}
		b.Cache.Set(cmd.Args[0], []byte(cmd.Args[1]))
		return replyOK(conn)
	case "del":
		if len(cmd.Args) == 0 {
			return replyErrArgLength(conn, cmd.Name)
		}
		for _, key := range cmd.Args {
			b.Cache.Set(key, nil)
		}
		return replyInteger(conn, len(cmd.Args))
	case "dbsize":
		dbsize := b.Cache.Len()
		return replyInteger(conn, dbsize)
	case "flushall":
		b.Cache.Reset()
		return replyOK(conn)
	case "info":
		// Give Info about server
		return replyBulkString(conn, []byte(b.InfoServer()+b.InfoClients()))
	case "ping":
		return replySimpleString(conn, "PONG")
	default:
		return replyErrUnknownCommand(conn, cmd.Name)
	}
}

func (b *Bigred) InfoServer() string {
	executableDir, err := filepath.Abs(os.Args[0])
	if err != nil {
		executableDir = "unknown"
	}
	// Key order not preserved
	mapInfo := map[string]interface{}{
		"server": map[string]interface{}{
			"bigred_version": "0.0.1",
			"os":             runtime.GOOS,
			"arch_bits":      runtime.GOARCH,
			"process_id":     os.Getpid(),
			"port":           b.Addr,
			"executable":     executableDir,
		},
	}
	var info string
	for opt, mapVal := range mapInfo {
		info += "#" + opt + "\n"
		for key, val := range mapVal.(map[string]interface{}) {
			info += fmt.Sprintf("%s: %v\n", key, val)
		}
	}
	return info
}

func (b *Bigred) InfoClients() string {
	info := "#Clients\n"
	info += fmt.Sprintf("connected_clients: %v\n", b.clients)
	return info
}
