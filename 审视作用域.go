package main

import (
     "fmt"
     "math/rand"
)

func main () {  
      var count = 0   // 声明变量
      for count < 10 {   // 循环10大于0
          var num = rand.Intn(10) +1  //声明num中的伪随机数+1
          fmt.Println(num)   // 打印num
          count++      // 后位数+1  
      
      } // 作用域2
      var count = num  // 循环结束，再次访问变量
      fmt.Println(num)
}  // 作用域2
