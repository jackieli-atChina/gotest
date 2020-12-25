package main

import (
	"log"
	"os"
	"runtime/pprof"
)

func main(){
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
}
