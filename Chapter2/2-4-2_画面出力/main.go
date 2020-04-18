package main

import "os"

func main() {
	// fmt.Printlnと等価
	os.Stdout.Write([]byte("os.Stdout example\n"))
}
