package main

import "sync/atomic"

//原子包

var atomicX int64

func AddX() {

	atomic.AddInt64(&atomicX, 1)
}
