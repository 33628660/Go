package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD" // 包内变量

func main() {
	// 代码4-6 变量作用域规则
	/*year := 2018      // 申明变量
	  // 分支判断month月份变量随机数，1-12月
	  rand.Seed(time.Now().UnixNano())
	  switch month := rand.Intn(12) + 1; month {
	  // 输出值是否为2月份，如果是天数+1
	  case 2:
	       day := rand.Intn(28) + 1
	       fmt.Println(era, year, month, day)
	  // 如果是4/6/9/11月份，天数+1
	  case 4, 6, 9, 11:
	       day := rand.Intn(30) + 1
	       fmt.Println(era, year, month, day)
	  // 如果不满足以上条件，则默认天数为30
	  default :
	       day := rand.Intn(31) + 1
	       fmt.Println(era, year, month, day)
	   }*/
	// 代码4-7 重构后的随机日期选取程序
	rand.Seed(time.Now().UnixNano())
	year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31
	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
	// 先定义变量月的随机数，定义变量天的随机数，然后分支循环月的天数，如果随机数是2月，则打印28天+1，反之正常打印其他月份天数+1，然后依次打印出，字符串，年份，月份，天数，相对于4-6，4-7不需要满足条件

}
