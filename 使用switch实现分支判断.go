package main

import (
	"fmt"
)

func main() {
	/*fmt.Println("找出含有go inside的关键字")
	  var command = "go inside"

	  switch command {
	  case "go east":
	        fmt.Println("A是这个嘛？")
	  case "enter cave", "go inside":
	        fmt.Println("go inside,对！是这个")
	  case "read sign":
	        fmt.Println("B是这个嘛？")
	  }*/

	var room = "lake"
	switch {
	case room == "cave":
		fmt.Println("A是这个嘛？")
	case room == "lake":
		fmt.Println("lake,对！是这个")
		fallthrough // 下降另一分支
	case room == "underwater":
		fmt.Println("我也不知道显示什么诶！")
	}	
        // 速查 3-5
}
