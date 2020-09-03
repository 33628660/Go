package main

import (
	"fmt"
)

func main() {
	var distance = 56000000	//单行声明变量
        var speed = 100800
        var {
             distance = 56000000  // 一次声明一组变量
             speed = 100800
            }
        var distance, speed = 56000000, 100800  // 在同一行内声明多个代码清单
        
        var weight = 149.0  //赋值操作符
        weight = weight * 0.3783
        weight *= 0.3783


        var age = 41    // 增量操作符
        age = age + 1 
        age += 1 
        age++
       
}
