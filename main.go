package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
)

const name = "/tmp/my-nice-pipe"

func getFile(path string) (*os.File, error) {
	err := syscall.Mkfifo(path, 0666)
	if err != nil {
		return nil, err
	}

	return os.OpenFile(path, os.O_RDONLY, os.ModeNamedPipe)
}

func main() {
	os.Remove(name)
	file, _ := getFile(name)
	reader := bufio.NewReader(file)

	line, _ := reader.ReadBytes('\n')
	fmt.Print(string(line))
}
