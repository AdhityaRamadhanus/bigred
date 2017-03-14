package main

import (
	"fmt"
	"testing"
)

func TestRespInteger5(t *testing.T) {
	bytesInt := respInteger(666)
	if string(bytesInt) != ":666\r\n" {
		t.Error("respInteger not returning resp comply string")
	}
}

func TestRespIntegerNeg(t *testing.T) {
	bytesInt := respInteger(-1)
	if string(bytesInt) != ":-1\r\n" {
		t.Error("respInteger not returning resp comply string")
	}
}

func TestRespError(t *testing.T) {
	bytesError := respError("Unknown Command")
	if string(bytesError) != "-ERR Unknown Command\r\n" {
		t.Error("respError not returning resp comply string")
	}
}

func TestRespBulkString(t *testing.T) {
	bulkString := "SET mykey Hello"
	bytesBulk := respBulkString([]byte(bulkString))
	if string(bytesBulk) != fmt.Sprintf("$%d\r\n%s\r\n", len(bulkString), bulkString) {
		t.Error("respError not returning resp comply string")
	}
}

func TestRespSimpleString(t *testing.T) {
	bulkString := "OK"
	bytesBulk := respSimpleString("OK")
	if string(bytesBulk) != fmt.Sprintf("+%s\r\n", bulkString) {
		t.Error("respError not returning resp comply string")
	}
}
