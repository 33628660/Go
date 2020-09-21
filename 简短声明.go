package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 代码4-2 更简洁的倒数程序
	/*var count = 0
	  for count = 10; count > 0; count-- {
	      fmt.Println(count)
	  }
	      fmt.Println(count)*/ // 处于作用域内

	// 代码4-3 在for循环中使用简短声明
	/*for count := 10; count > 0; count-- {
	    fmt.Println(count)  // :=变量将不再处于作用域内

	}*/

	// 代码4-4 在if语句中使用简短声明
	/*rand.Seed(time.Now().UnixNano())
	  if num := rand.Intn(3); num == 0 {   // 随机数0-2,等同于0打印文本
	    fmt.Println("Space Adventures")
	  } else if num == 1 {   // 如果等同于1,打印该文本

	    fmt.Println("SpaceX")
	  } else {    // 否则打印该文本
	    fmt.Println("Virgin Galactic")
	  }*/

	// 代码4-5 在switch中使用简短声明
	rand.Seed(time.Now().UnixNano())
	switch num := rand.Intn(10); num { // 在switch筛选中，伪随机数大于2以上则打印第四行出现的文本,伪随机数
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num)
	}
}
