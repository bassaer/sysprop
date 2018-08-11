package main

import (
	"bufio"
	"os"
)

func main() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("first\n")
	buffer.Flush()
	buffer.WriteString("second\n")
	buffer.Flush()
}
