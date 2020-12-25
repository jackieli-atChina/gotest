package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func Counter(wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Millisecond * 200)
		counter++
	}
	wg.Done()
}

func main() {
	cpuf, err := os.Create("cpu_profile")
	if err != nil {
		log.Fatal(err)
	}
	//cpu profile
	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

	memf, err := os.Create("mem_profile")
	if err != nil {
		log.Fatal(err)
	}
	//heap profile
	if err := pprof.WriteHeapProfile(memf); err != nil {
		log.Fatal(err)
	}
	defer memf.Close()


	flag.Parse()
	//远程获取pprof数据
	go func() {
		//log.Println(http.ListenAndServe("localhost:8001", nil))
		log.Println(http.ListenAndServe("0.0.0.0:8001", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Counter(&wg)
	}
	wg.Wait()
	// sleep 10mins, 在程序退出之前可以查看性能参数.
	time.Sleep(60 * time.Second)
}