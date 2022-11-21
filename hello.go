// //kome
// /*sasa*/
// //注意事項
// //変数がmsgなど、1文字目が小文字で始まっていた場合はprivateなので他のパッケージからアクセスできない。
// //一文字目がMsgのように大文字だった場合、他のパッケージからアクセス可能

// //package main

// import "fmt"

// func main() {
// 	var msg string
// 	msg = "hello world"
// 	//var msg="hello"
// 	//↑これでもOK

// 	//msg:="hello world"

// 	//var a,b int
// 	//a,b:=10,15
// 	var (
// 		c int
// 		d string
// 	)
// 	c = 20
// 	d = "shimayu"
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(msg)

// }
