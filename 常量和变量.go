// 到达火星需要多长时间
package main

import (
	"fmt"
)

func main() {
	//const lightSpeed = 299792// km/s	
        //var distance = 56000000 //km/s
        //fmt.Println(distance/lightSpeed, "seconds")
        //distance = 401000000
        //fmt.Println(distance/lightSpeed, "seconds")
        const hoursPerDay = 24   // 申明常量
        var speed = 100800   // 申明变量  km/h
        var distance = 96300000  // 声明变量 km
        fmt.Println(distance/speed/hoursPerDay, "Day")  // 计算飞船从地球到火星的天数
}
