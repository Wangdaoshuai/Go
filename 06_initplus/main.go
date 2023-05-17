package main

import (
	"fmt"
	_ "gotest/05_init/lib1"      //匿名导包  不使用其方法也不报错
	mylib2 "gotest/05_init/lib2" //起别名
)

func main() {
	fmt.Println("hello")
	//lib2.Lib2Test()
	//Lib1Test()
	mylib2.Lib2Test()
}
