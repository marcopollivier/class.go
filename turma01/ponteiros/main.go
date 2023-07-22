package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	i := rand.Intn(100)

	fmt.Printf("%v at %p\n", i, &i)

	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)

	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MB", bToMb(stats.Alloc))
	fmt.Printf("\tTotalAlloc = %v MB", bToMb(stats.TotalAlloc))
	fmt.Printf("\tSys = %v MB", bToMb(stats.Sys))
	fmt.Printf("\tNumGC = %v\n", stats.NumGC)

}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
