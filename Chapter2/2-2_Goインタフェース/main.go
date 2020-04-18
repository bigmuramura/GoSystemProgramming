package main

import "fmt"

// Talker ... インタフェースを定義
type Talker interface{
	Talk()
}

// Greeter ... 構造体を宣言
type Greeter struct{
	name string
}

// Talk ... 構造体はTalkerインタフェースで定義されているメソッドを持っている
func (g Greeter) Talk(){
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main(){
	var talker Talker
	talker = &Greeter{"wozozo"}
	talker.Talk()
}