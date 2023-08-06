package main

import (
	"runtime"
	"sync"
	"testing"
)

func testAtomicIncrease(myatomic MyAtomic) {
	// set cpu cores to 4，which makes false-string more likely
	runtime.GOMAXPROCS(4)
	paraNum := 10000
	repeatNum := 10000
	var wg sync.WaitGroup
	for i := 0; i < paraNum; i++ {
		// wg.Add should not in go func,if wg.add in goroutine, can cause "sync: WaitGroup misuse: Add called concurrently with Wait"
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < repeatNum; j++ {
				myatomic.IncreaseA()
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < repeatNum; j++ {
				myatomic.IncreaseB()
			}
		}()
	}
	wg.Wait()
}

func BenchmarkNoPad(b *testing.B) {
	myatomic := &NoPad{}
	b.ResetTimer()
	b.StartTimer()
	testAtomicIncrease(myatomic)
}

func BenchmarkPad(b *testing.B) {
	myatomic := &Pad{}
	b.ResetTimer()
	b.StartTimer()
	testAtomicIncrease(myatomic)
}
