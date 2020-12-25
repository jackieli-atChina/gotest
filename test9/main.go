package main
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type LogProcess struct {
	//定义了两个管道，一个读取，一个写入
	rc chan string
	wc chan string

	//文件路径
	path string
}

func (l *LogProcess) ReadFromFile() {
	//打开文件返回一个reader
	file, err := os.Open(l.path)
	if err != nil {
		log.Fatal(err)
	}
	//(0,2)是指向文件最后一行
	file.Seek(0, 2)
	//新建一个reader，读取打开的文件
	read := bufio.NewReader(file)
	for {
		//循环读取文件返回一个bute的一个切片参数意思为换行
		bytes, err := read.ReadBytes('\n')
		//如果错误为io.EOF代表这是文件末尾，休眠500毫秒跳过本次读取
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatal(err)
		}
		//把读取的数据转换为string放入读取管道
		l.rc <- string(bytes)
	}

}
func (l *LogProcess) Process() {
	//处理方法从读取管道中拿到数据进行处理，我们进行最简单的处理
	for {
		data := <-l.rc
		//strings.ToUppser方法为把字母全部转换为大写
		res := strings.ToUpper(data)
		//把处理结果放入写入管道
		l.wc <- res
	}

}
func (l *LogProcess) Write() {
	//从写入管道中拿到数据，进行写入，因为现在为调测，我们进行最简单的fmt操作
	for {
		fmt.Print(<-l.wc)
	}

}

func main() {
	//这里我们实现结构体的时候新增了文件路径为当前目录下的log文件
	l := &LogProcess{
		rc:   make(chan string),
		wc:   make(chan string),
		path: "./test8/access.log",
	}
	go l.ReadFromFile()
	go l.Process()
	go l.Write()

	time.Sleep(30 * time.Second)
}