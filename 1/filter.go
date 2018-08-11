package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(nil)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}
