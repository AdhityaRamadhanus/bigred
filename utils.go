package main

import "strconv"

func respSimpleString(payload string) []byte {
	return []byte("+" + payload + "\r\n")
}

func respBulkString(payload []byte) []byte {
	bytesBulk := []byte("$" + strconv.Itoa(len(payload)) + "\r\n")
	bytesBulk = append(bytesBulk, payload...)
	bytesBulk = append(bytesBulk, '\r', '\n')
	return bytesBulk
}

func respInteger(payload int) []byte {
	return []byte(":" + strconv.Itoa(payload) + "\r\n")
}

func respError(message string) []byte {
	return []byte("-ERR " + message + "\r\n")
}
