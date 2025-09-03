package gomap

type GoMap[K comparable, V any] struct {
	m map[K]V
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func New[K comparable, V any]() *GoMap[K, V] {
	return &GoMap[K, V]{
		m: make(map[K]V),
	}
}

func (m *GoMap[K, V]) Put(key K, value V) {
	m.m[key] = value
}

func (m *GoMap[K, V]) Remove(key K) {
	delete(m.m, key)
}

func (m *GoMap[K, V]) Get(key K) (V, bool) {
	value, ok := m.m[key]
	return value, ok
}

func (m *GoMap[K, V]) Contains(key K) bool {
	_, ok := m.m[key]
	return ok
}

func (m *GoMap[K, V]) Size() (size uint32) {
	return uint32(len(m.m))
}

func (m *GoMap[K, V]) Empty() bool {
	return len(m.m) == 0
}

func (m *GoMap[K, V]) Clear() {
	clear(m.m)
}

func (m *GoMap[K, V]) Keys() (keys []K) {
	for key := range m.m {
		keys = append(keys, key)
	}
	return
}

func (m *GoMap[K, V]) Values() (values []V) {
	for _, value := range m.m {
		values = append(values, value)
	}
	return
}

func (m *GoMap[K, V]) Enumerate() (entries []Entry[K, V]) {
	for key, value := range m.m {
		entries = append(entries, Entry[K, V]{
			Key:   key,
			Value: value,
		})
	}
	return
}

func (m *GoMap[K, V]) ForEach(f func(K, V) bool) {
	for key, value := range m.m {
		if !f(key, value) {
			break
		}
	}
}

func (m *GoMap[K, V]) ForEachMutable(f func(K, *V) bool) {
	for key := range m.m {
		value := m.m[key]
		if !f(key, &value) {
			break
		}
		m.m[key] = value
	}
}
