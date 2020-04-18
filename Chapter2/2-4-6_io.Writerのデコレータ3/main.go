package main

import (
	"bufio"
	"os"
)

func main() {
	// Flush()メソッドでためておいた結果を吐き出す。
	// Flush()しないとそのまま消滅する。
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer1\n")
	buffer.WriteString("bufio.Writer2\n")
	buffer.Flush()
	buffer.WriteString("bufio.Writer3\n")
	buffer.WriteString("bufio.Writer4\n")
	buffer.WriteString("bufio.Writer5\n")
	buffer.Flush()
	buffer.WriteString("bufio.Writer6\n") // これは表示されずに終わる
}
