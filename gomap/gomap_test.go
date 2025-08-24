package gomap

import (
	"math/rand"
	"testing"
)

func TestNew(t *testing.T) {
	m := New[int, any]()
	if m.Size() != 0 {
		t.Error()
	}
	if !m.Empty() {
		t.Error()
	}
}

func TestPut(t *testing.T) {
	m := New[int, any]()
	nums := rand.Perm(10000)
	for _, num := range nums {
		m.Put(num, struct{}{})
	}
}

func TestRemove(t *testing.T) {
	m := New[int, any]()
	nums := rand.Perm(10000)
	for _, num := range nums {
		m.Put(num, struct{}{})
	}

	nums = rand.Perm(10000)
	for _, num := range nums {
		m.Remove(num)
	}
}

func TestGet(t *testing.T) {
	m := New[int, any]()
	nums := rand.Perm(10000)
	for _, num := range nums {
		m.Put(num, struct{}{})
		if _, ok := m.Get(num); !ok {
			t.Error()
		}
	}

	nums = rand.Perm(10000)
	for _, num := range nums {
		m.Remove(num)
		if _, ok := m.Get(num); ok {
			t.Error()
		}
	}
}

func TestClear(t *testing.T) {
	m := New[int, any]()
	nums := rand.Perm(10000)
	for _, num := range nums {
		m.Put(num, struct{}{})
	}

	m.Clear()
	if m.Size() != 0 {
		t.Error()
	}
	if !m.Empty() {
		t.Error()
	}
}

func TestForEach(t *testing.T) {
	m := New[int, int]()
	nums := rand.Perm(10000)
	for _, num := range nums {
		m.Put(num, num)
	}

	m.ForEach(func(key int, value int) bool {
		value *= 2
		return true
	})

	m.ForEach(func(key int, value int) bool {
		if key != value {
			t.Error()
		}
		return true
	})
}

func TestForEachMutable(t *testing.T) {
	m := New[int, int]()
	nums := rand.Perm(10000)
	for _, num := range nums {
		m.Put(num, num)
	}

	m.ForEachMutable(func(key int, value *int) bool {
		*value = *value * 2
		return true
	})

	m.ForEach(func(key int, value int) bool {
		if key*2 != value {
			t.Error()
		}
		return true
	})
}

func benchmarkPut(b *testing.B, m *GoMap[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, m *GoMap[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Remove(n)
		}
	}
}

func benchmarkGet(b *testing.B, m *GoMap[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func BenchmarkPut10(b *testing.B) {
	size := 10
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}

	b.ResetTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkRemove10(b *testing.B) {
	size := 10
	m := New[int, struct{}]()
	for i := 0; i < size; i++ {
		m.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkGet10(b *testing.B) {
	size := 10
	m := New[int, struct{}]()
	for i := 0; i < size; i++ {
		m.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkPut100(b *testing.B) {
	size := 100
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}

	b.ResetTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkRemove100(b *testing.B) {
	size := 100
	m := New[int, struct{}]()
	for i := 0; i < size; i++ {
		m.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkGet100(b *testing.B) {
	size := 100
	m := New[int, struct{}]()
	for i := 0; i < size; i++ {
		m.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkPut1000(b *testing.B) {
	size := 1000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}

	b.ResetTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkRemove1000(b *testing.B) {
	size := 1000
	m := New[int, struct{}]()
	for i := 0; i < size; i++ {
		m.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkGet1000(b *testing.B) {
	size := 1000
	m := New[int, struct{}]()
	for i := 0; i < size; i++ {
		m.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkPut10000(b *testing.B) {
	size := 10000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}

	b.ResetTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkRemove10000(b *testing.B) {
	size := 10000
	m := New[int, struct{}]()
	for i := 0; i < size; i++ {
		m.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkGet10000(b *testing.B) {
	size := 10000
	m := New[int, struct{}]()
	for i := 0; i < size; i++ {
		m.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkPut100000(b *testing.B) {
	size := 100000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}

	b.ResetTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkRemove100000(b *testing.B) {
	size := 100000
	m := New[int, struct{}]()
	for i := 0; i < size; i++ {
		m.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkGet100000(b *testing.B) {
	size := 100000
	m := New[int, struct{}]()
	for i := 0; i < size; i++ {
		m.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkGet(b, m, size)
}
