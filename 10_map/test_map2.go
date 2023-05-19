package main

import (
	"fmt"
)

func printMap(cityMap map[string]string) {
	//cityMap是一个引用传递
	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["england"] = "london"
}
func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["china"] = "beijing"
	cityMap["japan"] = "tokyo"
	cityMap["usa"] = "newyork"

	//遍历
	printMap(cityMap)

	//删除
	delete(cityMap, "china")

	//修改
	cityMap["usa"] = "dc"
	ChangeValue(cityMap)

	fmt.Println("---------")

	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}
}
