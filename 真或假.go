package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.") // 打印文本
	var command = "walk outside"  // 声明一个变量
	var exit = strings.Contains(command, "outside")  // 声明变量exit里面的strings包里面的函数Contains里面的变量command是否包含outside
	fmt.Println("You leave the cave:", exit)  // 打印exit，如果有，则为true
}
