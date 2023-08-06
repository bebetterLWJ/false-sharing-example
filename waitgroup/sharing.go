package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type MyAtomic interface {
	IncreaseAll()
	IncreaseA()
	IncreaseB()
}

type NoPad struct {
	a uint64
	b uint64
	c uint64
}

func (n *NoPad) IncreaseAll() {
	atomic.AddUint64(&n.a, 1)
	atomic.AddUint64(&n.b, 1)
	atomic.AddUint64(&n.c, 1)
}

func (n *NoPad) IncreaseA() {
	atomic.AddUint64(&n.a, 1)
}

func (n *NoPad) IncreaseB() {
	atomic.AddUint64(&n.b, 1)
}

// a、_p1 占满一个 cpu 缓存行
type Pad struct {
	a   uint64
	_p1 [7]uint64
	b   uint64
	_p2 [7]uint64
	c   uint64
	_p3 [7]uint64
}

func (p *Pad) IncreaseAll() {
	atomic.AddUint64(&p.a, 1)
	atomic.AddUint64(&p.b, 1)
	atomic.AddUint64(&p.c, 1)
}

func (p *Pad) IncreaseA() {
	atomic.AddUint64(&p.a, 1)
}

func (p *Pad) IncreaseB() {
	atomic.AddUint64(&p.b, 1)
}

func main() {
	pad := &Pad{}
	pad.IncreaseAll()
	fmt.Printf("%d\n", pad.a)
	fmt.Println(unsafe.Sizeof(Pad{}))
}
