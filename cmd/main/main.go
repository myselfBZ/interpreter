package main

import (
	"log"
	"os"

	"github.com/myselfBZ/interpreter/internal/lexer"
)

func open(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("couldn't find the file")
	}
	buff := make([]byte, 1024)
	size, err := file.Read(buff)
	if err != nil {
		log.Fatal("error reading from a file, ", err)
	}
	return string(buff[:size])
}

func main() {
	src := open("test.monkey")
	l := lexer.New(src)
	l.Tokenize()
}
