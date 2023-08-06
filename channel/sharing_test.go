package main

import (
	"fmt"
	"runtime"
	"testing"
)

func testAtomicIncrease(myatomic MyAtomic) {
	runtime.GOMAXPROCS(4)
	paraNum := 1000
	repeatNum := 1000
	channelA := make(chan int, 1)
	channelB := make(chan int, 1)
	channelA <- 1
	for i := 0; i < paraNum; i++ {
		go func() {
			for j := 0; j < repeatNum; j++ {
				fmt.Println(i, j)
				<-channelB
				myatomic.IncreaseA()
				channelA <- 0
			}
		}()
		go func() {
			for j := 0; j < repeatNum; j++ {
				<-channelA
				myatomic.IncreaseB()
				channelB <- 0
			}
		}()
	}
	<-channelA

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
