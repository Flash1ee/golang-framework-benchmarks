package main

import (
	"sync"

	"main/benchmarks"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	benchmarks.StartHttp()
	benchmarks.StartEcho()
	benchmarks.StartFiber()
	benchmarks.StartGin()
	wg.Wait()
}
