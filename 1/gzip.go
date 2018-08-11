package main

import (
	"compress/gzip"
	"os"
)

func main() {
	file, err := os.Create("gzip-sample.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "gzip-sample.txt"
	writer.Write([]byte("gzip.Write example\n"))
	writer.Close()
}
