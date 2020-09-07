package main

import (
	"fmt"
)

func main() {
	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")
	var age = 41                                       // 声明一个age为41的变量
	var minor = age < 18                               // 声明判断18是否小于age等于minor
	fmt.Printf("At age %v, am I minor %v", age, minor) // 打印minor文本数值，判断结果是否为false
}
