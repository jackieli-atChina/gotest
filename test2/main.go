package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	//创建一个文件，
	file, err := os.Create("./test2/log")
	if err != nil {
		log.Fatal(err)
	}
	for{
		//在文件写入数据
		file.Write([]byte("hello world\n"))
		fmt.Print("ok")
		time.Sleep(500 * time.Millisecond)
	}
}