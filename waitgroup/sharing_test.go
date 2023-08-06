package main

import (
	"runtime"
	"sync"
	"testing"
)

func testAtomicIncrease(myatomic MyAtomic) {
	runtime.GOMAXPROCS(4)
	paraNum := 1000
	repeatNum := 1000
	var wg sync.WaitGroup
	for i := 0; i < paraNum; i++ {
		go func() {
			wg.Add(1)
			for j := 0; j < repeatNum; j++ {
				myatomic.IncreaseA()
			}
			wg.Done()
		}()

		go func() {
			wg.Add(1)
			for j := 0; j < repeatNum; j++ {
				myatomic.IncreaseB()
			}
			wg.Done()
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
