package main

import (
	"fmt"
)

func main() {
	var command = "go east"   // 声明一个变量赋值
	if command == "go east" { // 第一句if判断为true
		fmt.Println("You head further up the mountain.")
	} else if command == "go east" { // 第二句判断为true
		fmt.Println("You enter the cave where you live out rest of your life.")
	} else { //如果1-2为false，则打印第三句
		fmt.Println("Didn't quite get that.")
	}
	// 速查3-3
}
