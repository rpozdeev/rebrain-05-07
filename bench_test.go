package main


import (
	"sync"
	"testing"
)

type Counter struct {
	A int
	B int
}

var pool = sync.Pool{
	New: func() interface{} { return new(Counter) },
}

func Increment(c *Counter) {
	c.A++
	c.B++
}

func BenchmarkWithoutPool(b *testing.B) {
	//var c *Counter
	c := &Counter{1, 1}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			b.StopTimer()
			Increment(c)
			b.StartTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	//var c *Counter
	c := &Counter{1, 1}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			c = pool.Get().(*Counter)
			//c.A = 1
			//c.B = 1
			b.StopTimer()
			Increment(c)
			b.StartTimer()
			pool.Put(c)
		}
	}
}
//type Small struct {
//	a int
//}
//
//var pool = sync.Pool{
//	New: func() interface{} { return new(Small) },
//}
//
////go:noinline
//func inc(s *Small) { s.a++ }
//
//func BenchmarkWithoutPool(b *testing.B) {
//	var s *Small
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < 10000; j++ {
//			s = &Small{ a: 1, }
//			b.StopTimer(); inc(s); b.StartTimer()
//		}
//	}
//}
//
//func BenchmarkWithPool(b *testing.B) {
//	var s *Small
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < 10000; j++ {
//			s = pool.Get().(*Small)
//			s.a = 1
//			b.StopTimer(); inc(s); b.StartTimer()
//			pool.Put(s)
//		}
//	}
//}