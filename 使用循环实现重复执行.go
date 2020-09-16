package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*var count = 10  // 申明变量赋值
	for count > 0 { // 循环变量大于0
		fmt.Println(count)      // 打印值
		time.Sleep(time.Second) // 设置时间条件
		count--                 // 满足条件停止
	}
	fmt.Println("Liftoff")*/

	/*var degrees = 0 // 申明变量赋值
	for {           // for循环
		fmt.Println(degrees) // 打印值0
		degrees++            // 后位数+1
		if degrees >= 360 {  // 判断变量大于等于360
			degrees = 0
			if rand.Intn(2) == 0 { // 生成伪随机数，是否相等
				break // 跳出循环，触发停止条件
			}
		}
	}*/
	// 速查 3-6
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {   // 判断伪随机数是否相等伪
			break  // 跳出循环
		}
		count--    // 满足条件停止
	}
	if count == 0 {   // 如果变量值相等于0，执行该结果
		fmt.Println("Liftoff!")
	} else {    // 否则执行该语句
		fmt.Println("Launch failed")
	}

}
