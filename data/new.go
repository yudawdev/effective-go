package data

import (
	"bytes"
	"fmt"
	"sync"
)

// SyncedBuffer 零值可用的设计模式
type SyncedBuffer struct {
	lock   sync.Mutex   // 零值是未锁定的互斥锁
	buffer bytes.Buffer // 零值是空的可用缓冲区
}

// # [play]
// ./prog.go:14:17: call of fmt.Println copies lock value: play.SyncedBuffer contains sync.Mutex
//
// Go vet failed.
//
// &{{0 0} {[] 0 0}} {{0 0} {[] 0 0}}
func newType() {
	p := new(SyncedBuffer)
	var v SyncedBuffer
	fmt.Println(p, v)
}
