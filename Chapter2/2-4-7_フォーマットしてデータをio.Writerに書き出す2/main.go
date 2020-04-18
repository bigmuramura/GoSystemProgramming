package main

import (
	"encoding/json"
	"os"
)

func main() {
	// JSON整形
	// 2-4-5_インターネットアクセスの送信2 の要領でブラウザにJSONを返すこともできる
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}
