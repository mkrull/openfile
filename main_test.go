package main

import (
	"os"
	"testing"
)

func TestGetFile(t *testing.T) {
	os.Remove(name)
	_, err := getFile(name)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
