package main

import (
	"sync"
	"time"
)

func processorB(ch <-chan int) {
	for v := range ch {
		doWork(v)
	}
}

func methodB(numPr, numIt int) bres {
	ret := bres{}
	benchTime := time.Now()

	chs := make([]chan int, numPr)
	for c := range chs {
		chs[c] = make(chan int, 512)
	}

	ret.tInit = time.Since(benchTime)
	benchTime = time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(numPr)
	for c := range chs {
		ch := chs[c]
		go func() {
			processorB(ch)
			wg.Done()
		}()
	}

	ret.tStart = time.Since(benchTime)
	benchTime = time.Now()

	for i := 0; i < numIt; i++ {
		for _, c := range chs {
			c <- i
		}
	}
	for _, c := range chs {
		close(c)
	}

	ret.tBcast = time.Since(benchTime)
	benchTime = time.Now()

	wg.Wait()

	ret.tDone = time.Since(benchTime)

	return ret
}
