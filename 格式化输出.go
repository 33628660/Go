package main

import (
	"fmt"
)
func main() {
	fmt.Printf("My weight on the surface of Mars is  %v lbs,", 102.0*0.3783)  // 格式化变量
	fmt.Printf("and I would be %v years old.\n", 24*365/687)
	fmt.Printf("My weight on the surface of %v is  %v lbs.\n", "Earth", 149.0) // 指定多个格式化变量
	fmt.Printf("%-15v $%10v\n", "spaceX", 94)  // 填充整数和负数
        fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
