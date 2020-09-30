package main

import (
	"design-patterns/singleton/hungry"
	"design-patterns/singleton/lazy"
	"testing"
)

/*
$ go test -bench=. -benchtime=3s
goos: linux
goarch: amd64
pkg: design-patterns/singleton
BenchmarkHungry-4                 	1000000000	         0.274 ns/op
BenchmarkLazyDirect-4             	843014878	         4.32 ns/op
BenchmarkLazyCheckNil-4           	837470934	         4.22 ns/op
BenchmarkHungryParallel-4         	1000000000	         0.402 ns/op
BenchmarkLazyDirectParallel-4     	1000000000	         1.12 ns/op
BenchmarkLazyCheckNilParallel-4   	1000000000	         1.12 ns/op
PASS
ok  	design-patterns/singleton	11.313s
*/

func BenchmarkHungry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		instance := hungry.GetInstance()
		_ = instance
	}
}

func BenchmarkLazyDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		instance := lazy.GetInstanceDirect()
		_ = instance
	}
}

func BenchmarkLazyCheckNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		instance := lazy.GetInstanceCheckNil()
		_ = instance
	}
}

func BenchmarkHungryParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			instance := hungry.GetInstance()
			_ = instance
		}
	})
}

func BenchmarkLazyDirectParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			instance := lazy.GetInstanceDirect()
			_ = instance
		}
	})
}

func BenchmarkLazyCheckNilParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			instance := lazy.GetInstanceCheckNil()
			_ = instance
		}
	})
}
