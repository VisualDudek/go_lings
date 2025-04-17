package main

import (
	"fmt"
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

func BenchmarkMyFuncCustomLabels(b *testing.B) {
	for _, size := range []int{10, 100, 1_000} {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			for i := range b.N {
				_ = i
				myFunc()
			}
		})
	}
}
