package main //程序的包名

import (
	"fmt"
	"time"
)

// main函数
func main() { //函数的{一定和函数名在同一行，否则编译错误
	//golang中的表达式，加"；"与不加都可以，建议不加
	fmt.Println("Hello")

	time.Sleep(1 * time.Second)
}
