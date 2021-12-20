package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
)

func pprofChecker() {

	var isCpuPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCpuPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()
	if isCpuPprof {

		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed,err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer func() {
			pprof.StopCPUProfile()
			file.Close()
		}()
	}

	if isMemPprof {

		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed,err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		defer func() {
			file.Close()
		}()
	}

}
