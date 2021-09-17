package sync

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestPool(t *testing.T) {
	v := int32(0)
	a := sync.Pool{New: func() interface{} {
		atomic.AddInt32(&v, 1)
		b := make([]int, 1024)
		return &b
	}}
	a.Put(a.Get())
	get2 := a.Get()
	a.Put(get2)
	fmt.Println(atomic.LoadInt32(&v))
	get3 := a.Get()
	a.Put(get3)
	fmt.Println(atomic.LoadInt32(&v))
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}
