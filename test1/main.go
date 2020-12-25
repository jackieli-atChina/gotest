package main

import (
	"fmt"
	"strings"
	"time"
)

type LogProcess struct {
	//定义了两个管道，一个读取，一个写入
	rc chan string
	wc chan string
}

func (l *LogProcess) ReadFromFile() {
	//测试数据
	data := "hello world"
	//模拟数据进入读取管道
	l.rc <- data
}
func (l *LogProcess) Process() {
	//处理方法从读取管道中拿到数据进行处理，我们进行最简单的处理
	data := <-l.rc
	//strings.ToUppser方法为把字母全部转换为大写
	res := strings.ToUpper(data)
	//把处理结果放入写入管道
	l.wc <- res

}
func (l *LogProcess) Write() {
	//从写入管道中拿到数据，进行写入，因为现在为调测，我们进行最简单的fmt操作
	fmt.Println(<-l.wc)
}

func main() {
	l := &LogProcess{
		rc: make(chan string),
		wc: make(chan string),
	}

	go l.ReadFromFile()
	go l.Process()
	go l.Write()

	time.Sleep(1 * time.Second)
}