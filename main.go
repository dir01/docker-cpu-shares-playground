package main

import "fmt"
import "sync"
import "os"
import "strconv"


func main() {
	workersCount, _ := strconv.Atoi(os.Getenv("WORKERS_COUNT"))
	var wg sync.WaitGroup
	for i:=1; i<=workersCount; i++ {
		wg.Add(1)	
		go burnCPU(i, &wg)
	}
	wg.Wait()
}

func burnCPU(id int, wg *sync.WaitGroup) {
    defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	for {
		fib(100000)
		// fmt.Printf("Worker %d got new batch of work\n", id)
	}
    fmt.Printf("Worker %d done\n", id)
}

func fib(n int) int {
	if n == 1 {
		return 1
	}
	return fib(n-1) + n
}