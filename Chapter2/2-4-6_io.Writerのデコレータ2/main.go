package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	// 書き込まれたデータをgzip圧縮して、あらかじめ渡されていたos.Fileに中継
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}
