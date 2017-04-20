package memory

import (
	"testing"
	"time"
)

func BenchmarkNew(b *testing.B) {
	println(1111)

	time.Sleep(time.Second * 10)
}