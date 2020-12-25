package  main

import (
	"fmt"
	"time"
)

func main(){
	// 创建一个channel
	c := make(chan string)
	go func() {
		time.Sleep(1*time.Second)
		// 发送数据到channel中
		c <- "test"
	}()
	// 阻塞接收到数据
	msg := <-c
	fmt.Println(msg)
}

/*func test(c1,c2){
	select {
	case v: = <-c1:
		fmt.Println(v)
	case v: = <-c2:
		fmt.Println(v)
	default:
		fmt.Println("none")
	}
}*/