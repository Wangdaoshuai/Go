package main

import (
	"fmt"
	. "gotest/05_init/lib1"
	. "gotest/05_init/lib2" //加入.的话表示当前lib2的包所有函数都包含进了main包里,可以直接使用方法名 （不建议使用这种方式，容易产生冲突）
)

func main() {
	fmt.Println("hello")
	Lib2Test()
	Lib1Test()
}
