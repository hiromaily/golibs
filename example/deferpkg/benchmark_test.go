package deferpkg

import (
	"fmt"
	"testing"
)

func withDefer() {
	fmt.Printf("%s", "-")
	defer func() {
		fmt.Printf("%s", "end")
	}()
}

func withoutDefer() {
	fmt.Printf("%s", "-")
	fmt.Printf("%s", "end")
}

func BenchmarkWithDefer(b *testing.B) {
	//5541 ns/o

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withDefer()
	}
	b.StopTimer()
}

func BenchmarkWithoutDefer(b *testing.B) {
	//5238 ns/op

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withoutDefer()
	}
	b.StopTimer()
}
