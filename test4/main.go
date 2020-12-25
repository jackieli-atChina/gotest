package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	//远程获取pprof数据
	go func() {
		log.Println(http.ListenAndServe("localhost:8001", nil))
	}()
	/*cpuf, err := os.Create("cpu_profile")
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
	defer memf.Close()*/

	time.Sleep(6000 * time.Second)
}