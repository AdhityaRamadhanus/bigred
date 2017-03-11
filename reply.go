package main

import (
	"net"
)

func replyOK(conn net.Conn) error {
	_, err := conn.Write(respSimpleString("OK"))
	return err
}

func replyNil(conn net.Conn) error {
	_, err := conn.Write([]byte("$-1\r\n"))
	return err
}

func replyBulkString(conn net.Conn, message []byte) error {
	lenBytes, messageBytes := respBulkString(message)
	_, err := conn.Write(lenBytes)
	if err != nil {
		return err
	}
	_, err = conn.Write(messageBytes)
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte("\r\n"))
	return err
}

func replyInteger(conn net.Conn, payload int) error {
	_, err := conn.Write(respInteger(payload))
	return err
}

func replySimpleString(conn net.Conn, message string) error {
	_, err := conn.Write(respSimpleString(message))
	return err
}

func replyErrArgLength(conn net.Conn, command string) error {
	_, err := conn.Write(respError("wrong number of arguments for '" + command + "' command"))
	return err
}

func replyErrUnknownCommand(conn net.Conn, command string) error {
	_, err := conn.Write(respError("unknown command '" + command + "'"))
	return err
}
