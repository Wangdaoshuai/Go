package main

import "fmt"

/*
	知识点一：defer的执行顺序

func func1() {
	fmt.Println("A")
}
func func2() {
	fmt.Println("B")
}
func func3() {
	fmt.Println("C")
}
func main(){
	defer func1()
	defer func2()
	defer func3()
}
*/

//知识点二：defer和return谁先谁后：defer在后面调用
func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}
func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}

/*
func main() {
	//写入defer关键字   压栈的形式执行顺序
	defer fmt.Println("main end1") //函数结束之前执行
	defer fmt.Println("main end2") //函数结束之前执行
	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")


}*/
