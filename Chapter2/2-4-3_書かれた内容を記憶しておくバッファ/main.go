package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	// Write()メソッドで書き込まれた内容をためておいて
	buffer.Write([]byte("bytes.Buffer example1\n"))
	buffer.Write([]byte("bytes.Buffer example2\n"))
	buffer.Write([]byte("bytes.Buffer example3\n"))
	// 後でまとめて表示できる
	fmt.Println(buffer.String())
}
