package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// %v はなんでも表示できる
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
