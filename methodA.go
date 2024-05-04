package main

import (
	"llbcastBench/bclist"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func processorA(validate *atomic.Int64, listen *bclist.BroadcastListener[*int]) {
	for {
		val := listen.Wait()
		if val == nil {
			return
		}
		doWork(*val)
		validate.Add(1)
	}
}

func methodA(numPr, numIt int) bres {
	ret := bres{}
	benchTime := time.Now()

	l := bclist.NewBroadcastList[*int]()

	ret.tInit = time.Since(benchTime)
	benchTime = time.Now()

	vld := &atomic.Int64{}

	wg := &sync.WaitGroup{}
	wg.Add(numPr)
	for i := 0; i < numPr; i++ {
		go func() {
			processorA(vld, l.GetListener())
			wg.Done()
		}()
	}

	ret.tStart = time.Since(benchTime)
	benchTime = time.Now()

	for i := 0; i < numIt; i++ {
		num := i
		l.Broadcast(&num)
	}

	l.Broadcast(nil)

	ret.tBcast = time.Since(benchTime)
	benchTime = time.Now()

	wg.Wait()

	ret.tDone = time.Since(benchTime)

	if int64(numIt)*int64(numPr) != vld.Load() {
		log.Println(int64(numIt)*int64(numPr), "!=", vld.Load())
	}

	return ret
}
