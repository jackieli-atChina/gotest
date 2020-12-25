package main

import (
	"fmt"
	"strings"
	"time"
)
// 定义接口
type Reader interface {
	Read(rc chan string)
}
type Writer interface {
	Write(wc chan string)
}
// 定义实现
type logProcess struct {
	// 传递数据
	rc chan string
	wc chan string
	read Reader
	write Writer
}
type ReadFromFile struct {
	// 文件读取路径
	path string
}
type WriteToDb struct {
	// db连接
	influxDBDsn string
}
// 实现
func (r *ReadFromFile) Read(rc chan string){
	line :="message"
	rc <-line
}
func (w *WriteToDb) Write(wc chan string){
	fmt.Println(<-wc)
}
// 解析模块
func (f *logProcess) Process(){
	data := <- f.rc
	f.wc <- strings.ToUpper(data)
}
func main() {
	r:= &ReadFromFile{
		path:"./access.log",
	}
	w:= &WriteToDb{
		influxDBDsn: "/",
	}
	l := &logProcess{
		rc: make(chan string),
		wc: make(chan string),
		read: r,
		write: w,
	}
	// 并发执行
	go l.read.Read(l.rc)
	go l.Process()
	go l.write.Write(l.wc)

	time.Sleep(1*time.Second)
}