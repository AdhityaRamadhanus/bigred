package main

import "strconv"

func respSimpleString(payload string) []byte {
	return []byte("+" + payload + "\r\n")
}

// A bit weird here
func respBulkString(payload []byte) ([]byte, []byte) {
	lenBytes := []byte("$" + strconv.Itoa(len(payload)) + "\r\n")
	return lenBytes, payload
}

func respInteger(payload int) []byte {
	return []byte(":" + strconv.Itoa(payload) + "\r\n")
}

func respError(message string) []byte {
	return []byte("-ERR" + message + "\r\n")
}
