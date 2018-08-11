package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer foo\n"))
	buffer.Write([]byte("bytes.Buffer hoge\n"))
	fmt.Println(buffer.String())
}
