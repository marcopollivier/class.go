package main

import (
	"fmt"
	"runtime"
)

func main() {

	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)

	fmt.Println(stats.HeapAlloc / 1024)

}
