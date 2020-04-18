package main

import (
	"fmt"
	"strings"
)

func main() {
	// Go 1.10から追加されたstrings.Builder 書き出し専用のbytes.Buffer
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example1\n"))
	builder.Write([]byte("strings.Builder example2\n"))
	builder.Write([]byte("strings.Builder example3\n"))
	fmt.Println(builder.String())
}
