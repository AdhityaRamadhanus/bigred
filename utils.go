package main

import "strconv"

// respSimpleString return resp comply simple string
// example
// input : "OK"
// output: "+OK\r\n"

func respSimpleString(payload string) []byte {
	return []byte("+" + payload + "\r\n")
}

// respBulkString return resp comply bulk string
// example
// input : "Hello"
// output: "$5\r\nHello\r\n"

func respBulkString(payload []byte) []byte {
	bytesBulk := []byte("$" + strconv.Itoa(len(payload)) + "\r\n")
	bytesBulk = append(bytesBulk, payload...)
	bytesBulk = append(bytesBulk, '\r', '\n')
	return bytesBulk
}

// respInteger return resp comply integer
// example
// input : 5
// output: ":5\r\n"

func respInteger(payload int) []byte {
	return []byte(":" + strconv.Itoa(payload) + "\r\n")
}

// respError return resp comply error string
// example
// input : "Unknown Command"
// output: "-ERR Unknown Command\r\n"

func respError(message string) []byte {
	return []byte("-ERR " + message + "\r\n")
}
