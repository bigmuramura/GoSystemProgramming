package main

import (
	"io"
	"os"
)

func main() {
	// ファイルと標準出力共に出力
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}
