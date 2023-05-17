package main

import (
	"fmt"
)

// 切片截取
func main() {
	s := []int{1, 2, 3}

	//[0,2)
	s1 := s[0:2]
	fmt.Println(s1)

	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)
	//s1和s底层指向同一块地址

	//copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3)
	//将s中的值 拷贝到s2
	copy(s2, s)
	fmt.Println(s2)
}
