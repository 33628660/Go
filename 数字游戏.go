package main

import (
	"fmt"
	//"math/rand"
)

func main() {
	//var num = rand.Intn(10) + 1   // 随机数字
	//fmt.Println(num)
	//num = rand.Intn(10) + 1
	//fmt.Println(num)	
     
        const ufo = 56000000   // 在距离56...km的情况下，宇宙飞船需要多少千米每小时，才能到达Malacandra
        var d = 24 * 28
        fmt.Print(ufo/d," km")
       
}
