package main

import (
	"fmt"
)

func main() {
	var year = 2100
	var leap = year%400 == 0 || (year%400 == 0 && year%100 != 0) // 如果被整除则为true，不能整除则为false，所以2100不是闰年，输出为false
	if leap {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	var haveTorch = true //非逻辑运算符，
	var litTorch = false

	if !haveTorch || !litTorch {
		fmt.Println(!haveTorch, !litTorch)
	} // haveTorch为true则为false,litTorch为false则为true
	// 速查 3-4
}
