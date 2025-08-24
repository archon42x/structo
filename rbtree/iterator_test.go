package rbtree

import (
	"math/rand"
	"testing"
)

func TestNext(t *testing.T) {
	tree := New[int, any]()
	nums := rand.Perm(10000)
	for _, num := range nums {
		tree.Put(num, struct{}{})
	}

	iter := tree.Iter()
	iter.Begin()
	prev_nil := true
	var prev int
	for iter.Next() {
		key, ok := iter.Key()
		if !ok {
			t.Error()
		}

		if prev_nil {
			prev_nil = false
		} else {
			if prev > key {
				t.Error()
			}
		}

		prev = key
	}
}

func TestPrev(t *testing.T) {
	tree := New[int, any]()
	nums := rand.Perm(10000)
	for _, num := range nums {
		tree.Put(num, struct{}{})
	}

	iter := tree.Iter()
	iter.End()
	prev_nil := true
	var prev int
	for iter.Prev() {
		key, ok := iter.Key()
		if !ok {
			t.Error()
		}

		if prev_nil {
			prev_nil = false
		} else {
			if prev < key {
				t.Error()
			}
		}

		prev = key
	}
}
