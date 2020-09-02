// 我的减重程序
package main

import (
	"fmt"
)
// main 是所有程序的起始函数
func main() {
        fmt.Print("My weight on the surface of Mars is ") 
        fmt.Print(102.0 * 0.3783) // 打印出地球乘以火星的体重
        fmt.Print(" lbs, and I would be ")
        fmt.Print(24 * 365 / 687) // 打印出地球年龄天数在火星的年龄
        fmt.Print(" years old.")
        fmt.Println("My weight on the surface of Mars is", 
        102.0 * 0.3783, "lbs, and I would be ", 24 * 365 / 687, 
        " years old.")     // 在一行内打印出文本,数值,数学表达式
}
