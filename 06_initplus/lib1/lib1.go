package lib1

import "fmt"

//当前Lib1提供的API
//函数开头字母大写代表外界可以调用，如果小写开头表示只能当前文件内调用
func Lib1Test() {
	fmt.Println("Lib1Test()...")
}
func init() {
	fmt.Println("lib1.init()...")
}
