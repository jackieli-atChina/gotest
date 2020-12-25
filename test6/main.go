package main

import (
	"fmt"
	"strings"
	"time"
)

type logProcess struct {
	// 传递数据
	rc chan string
	wc chan string
	// 文件读取路径
	path string
	// db连接
	influxDBDsn string
}
// 读取模块
func (f *logProcess) Read(){
		line :="message"
		f.rc <-line
}
// 解析模块
func (f *logProcess) Process(){
	data := <- f.rc
	f.wc <- strings.ToUpper(data)
}
// 写入模块
func (f *logProcess) Write(){
	fmt.Println(<-f.wc)
}
func main() {
	l := &logProcess{
		rc: make(chan string),
		wc: make(chan string),
		path: "/",
		influxDBDsn: "/",
	}
	// 并发执行
	go l.Read()
	go l.Process()
	go l.Write()

	time.Sleep(1*time.Second)
}