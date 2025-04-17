package main

import (
	"testing"
	"time"
)

func myFunc() {
	time.Sleep(time.Millisecond * 100)
}

func BenchmarkMyFunc(b *testing.B) {
	for i := range b.N {
		_ = i
		myFunc()
	}

}
