package main

import (
	"net/http"
	"os"
)

func main() {
	// 2-4-5_インターネットアクセスの送信2 のRequest構造体を使った版
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できる")
	request.Write(os.Stdout)
}
