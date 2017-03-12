package main

import (
	"log"
	"net"
	"time"

	"github.com/allegro/bigcache"
)

// Bigred is the main service struct consist of cache engine bigcache
type Bigred struct {
	Proto string
	Addr  string
	Cache *bigcache.BigCache
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
		go b.handleClient(conn)
	}
}

func (b *Bigred) handleClient(conn net.Conn) (err error) {
	defer func() {
		if err != nil {
			log.Println(err)
		}
		conn.Close()
	}()
	for {
		commands, err := ParseRequest(conn)
		if err != nil {
			return err
		}
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
	case "ping":
		return replySimpleString(conn, "PONG")
	default:
		return replyErrUnknownCommand(conn, cmd.Name)
	}
}
