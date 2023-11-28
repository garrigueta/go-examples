package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// we'll generate 1,000,000,000 values for our summing operations
	vals := gen(1e9)
	var sum int64
	t := time.Now()
	sqtime := time.Since(t)
	cntime := time.Since(t)

	for i := 1; i < 6; i++ {
		fmt.Println("\nrun number", (i))
		t = time.Now()
		sum = sqadd(vals)
		sqtime = time.Since(t)
		fmt.Println("Sequential operation time: ", sqtime, "total sum is: ", sum)

		t = time.Now()
		sum = cnadd(vals)
		cntime = time.Since(t)
		fmt.Println("Concurrent operation time: ", cntime, "total sum is: ", sum)
		fmt.Println("Concurrent operation took", sqtime-cntime, "less time to compute sum!")
		fmt.Printf("Concurrent operation took %.2f %% as much time to compute sum!\n", 100*float64(cntime)/float64(sqtime))
	}
}

// gen is for generating random values to work with
func gen(max int) []int {
	rand.Seed(time.Now().UnixNano())
	vals := make([]int, max)
	for i := 0; i < max; i++ {
		vals[i] = rand.Intn(10)
	}
	return vals
}

// sqadd function adds numbers in a sequential pattern
func sqadd(vals []int) int64 {
	var sum int64
	for _, n := range vals {
		sum += int64(n)
	}
	return sum
}

// cnadd function adds values using a concurrent pattern
// to take advantage of multiple cpus/cores
func cnadd(vals []int) int64 {

	// determine # of cores on this device
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	var sum int64
	max := len(vals)

	splits := max / cores

	var wg sync.WaitGroup

	for i := 0; i < cores; i++ {

		// here we'll split the input up for concurrent operation
		start := i * splits
		end := start + splits
		split := vals[start:end]

		// now we'll use separate goroutines to compute the sum
		wg.Add(1)
		go func(vals []int) {
			defer wg.Done()

			var splitSum int64

			// compute the sum for each split part of the input
			for _, n := range vals {
				splitSum += int64(n)
			}

			// compute the final total using all splits to obtain the total
			atomic.AddInt64(&sum, splitSum)
		}(split)
	}

	wg.Wait()
	return sum
}
