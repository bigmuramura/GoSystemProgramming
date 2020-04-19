package main

import (
	"fmt"
	"os"
)

func main() {
	num := 1
	flo := 2.0
	str := "3"

	fmt.Fprintf(os.Stdout, "数値 %d, 小数 %f, 文字 %s", num, flo, str)
}
