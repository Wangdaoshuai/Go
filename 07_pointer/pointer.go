package main

import "fmt"

func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}
func swap1(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp

}
func main() {
	var a int = 10
	var b int = 20

	//swap
	swap(a, b)
	fmt.Println("a=", a, "b=", b)
	swap1(&a, &b)
	fmt.Println("a=", a, "b=", b)

	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int //二级指针
	pp = &p      //用p的地址赋值于pp
	fmt.Println(&p)
	fmt.Println(pp)
}
