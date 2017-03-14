package main

import (
	"net"
)

// Helper to write simple string OK to client
func replyOK(conn net.Conn) error {
	_, err := conn.Write(respSimpleString("OK"))
	return err
}

// Helper to write nil object to client, used when get return nil
func replyNil(conn net.Conn) error {
	_, err := conn.Write([]byte("$-1\r\n"))
	return err
}

// Helper to write Bulk string to client

func replyBulkString(conn net.Conn, message []byte) error {
	messageBytes := respBulkString(message)
	_, err := conn.Write(messageBytes)
	return err
}

// Helper to write integer to client

func replyInteger(conn net.Conn, payload int) error {
	_, err := conn.Write(respInteger(payload))
	return err
}

// Helper to write simple string to client

func replySimpleString(conn net.Conn, message string) error {
	_, err := conn.Write(respSimpleString(message))
	return err
}

// Helper to write error that associated with arguments length to the command

func replyErrArgLength(conn net.Conn, command string) error {
	_, err := conn.Write(respError("wrong number of arguments for '" + command + "' command"))
	return err
}

// Helper to write error that associated with unknown command

func replyErrUnknownCommand(conn net.Conn, command string) error {
	_, err := conn.Write(respError("unknown command '" + command + "'"))
	return err
}
