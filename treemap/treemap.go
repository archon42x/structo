package treemap

import (
	"cmp"

	"github.com/archon42x/structo/rbtree"
	"github.com/archon42x/structo/utils"
)

type TreeMap[K comparable, V any] struct {
	tree *rbtree.RBTree[K, V]
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func New[K cmp.Ordered, V any]() *TreeMap[K, V] {
	return &TreeMap[K, V]{
		tree: rbtree.New[K, V](),
	}
}

func NewWith[K comparable, V any](comparator utils.Comparator[K]) *TreeMap[K, V] {
	return &TreeMap[K, V]{
		tree: rbtree.NewWith[K, V](comparator),
	}
}

func (m *TreeMap[K, V]) Put(key K, value V) {
	m.tree.Put(key, value)
}

func (m *TreeMap[K, V]) Remove(key K) {
	m.tree.Remove(key)
}

func (m *TreeMap[K, V]) Get(key K) (V, bool) {
	return m.tree.Get(key)
}

func (m *TreeMap[K, V]) Contains(key K) bool {
	_, ok := m.tree.Get(key)
	return ok
}

func (m *TreeMap[K, V]) Size() (size uint32) {
	return m.tree.Size()
}

func (m *TreeMap[K, V]) Empty() bool {
	return m.tree.Size() == 0
}

func (m *TreeMap[K, V]) Clear() {
	m.tree.Clear()
}

func (m *TreeMap[K, V]) Keys() (keys []K) {
	iter := m.tree.Iter()
	for iter.Next() {
		key, ok := iter.Key()
		if !ok {
			panic("treemap's tree structure error")
		}
		keys = append(keys, key)
	}
	return
}

func (m *TreeMap[K, V]) Values() (values []V) {
	iter := m.tree.Iter()
	for iter.Next() {
		value, ok := iter.Value()
		if !ok {
			panic("treemap's tree structure error")
		}
		values = append(values, value)
	}
	return
}

func (m *TreeMap[K, V]) Enumerate() (entries []Entry[K, V]) {
	iter := m.tree.Iter()
	for iter.Next() {
		key, ok := iter.Key()
		if !ok {
			panic("treemap's tree structure error")
		}
		value, ok := iter.Value()
		if !ok {
			panic("treemap's tree structure error")
		}
		entries = append(entries, Entry[K, V]{
			Key:   key,
			Value: value,
		})
	}
	return
}

func (m *TreeMap[K, V]) ForEach(f func(K, V) bool) {
	iter := m.tree.Iter()
	for iter.Next() {
		key, ok := iter.Key()
		if !ok {
			panic("treemap's tree structure error")
		}
		value, ok := iter.Value()
		if !ok {
			panic("treemap's tree structure error")
		}
		if !f(key, value) {
			break
		}
	}
}

func (m *TreeMap[K, V]) ForEachMutable(f func(K, *V) bool) {
	iter := m.tree.Iter()
	for iter.Next() {
		key, ok := iter.Key()
		if !ok {
			panic("treemap's tree structure error")
		}
		value, ok := iter.ValuePtr()
		if !ok {
			panic("treemap's tree structure error")
		}
		if !f(key, value) {
			break
		}
	}
}
