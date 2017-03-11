package main

import (
	"log"
	"net"
	"time"

	"strconv"

	"github.com/allegro/bigcache"
)

type Bigred struct {
	Proto string
	Addr  string
	Cache *bigcache.BigCache
}

func NewBigRed(proto, addr string) (*Bigred, error) {
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
	defer conn.Close()
	for {
		commands, err := ParseRequest(conn)
		if err == nil {
			if err := b.executeCommand(commands, conn); err != nil {
				log.Println(err)
			}
		}
	}
}

func (b *Bigred) executeCommand(cmd *Commands, conn net.Conn) error {
	switch cmd.Name {
	case "get":
		if len(cmd.Args) < 1 {
			_, err := conn.Write([]byte("-ERR wrong number of arguments for 'get' command\r\n"))
			return err
		}
		bytes, err := b.Cache.Get(cmd.Args[0])
		if err != nil || len(bytes) == 0 {
			_, err := conn.Write([]byte("$-1\r\n"))
			return err
		}
		_, err = conn.Write([]byte("$" + strconv.Itoa(len(bytes)) + "\r\n"))
		if err != nil {
			return err
		}
		_, err = conn.Write(bytes)
		if err != nil {
			return err
		}
		_, err = conn.Write([]byte("\r\n"))
		return err
	case "set":
		if len(cmd.Args) < 2 {
			_, err := conn.Write([]byte("-ERR wrong number of arguments for 'set' command\r\n"))
			return err
		}
		b.Cache.Set(cmd.Args[0], []byte(cmd.Args[1]))
		_, err := conn.Write([]byte("+OK\r\n"))
		return err
	case "del":
		if len(cmd.Args) == 0 {
			_, err := conn.Write([]byte("-ERR wrong number of arguments for 'del' command\r\n"))
			return err
		}
		for _, key := range cmd.Args {
			b.Cache.Set(key, nil)
		}
		_, err := conn.Write([]byte(":" + strconv.Itoa(len(cmd.Args)) + "\r\n"))
		return err
	case "dbsize":
		dbsize := b.Cache.Len()
		_, err := conn.Write([]byte(":" + strconv.Itoa(dbsize) + "\r\n"))
		return err
	case "flushall":
		b.Cache.Reset()
		_, err := conn.Write([]byte("+OK\r\n"))
		return err
	case "ping":
		conn.Write([]byte("+PONG\r\n"))
		return nil
	default:
		conn.Write([]byte("-ERR unknown command '" + cmd.Name + "'\r\n"))
		return nil
	}
}
