package avltree

import (
	"cmp"
	"math/rand"
	"testing"
)

type User struct {
	name string
	age  uint32
}

func UserComparator(x, y User) int {
	if x.age == y.age {
		return cmp.Compare(x.name, y.name)
	} else {
		return cmp.Compare(x.age, y.age)
	}
}

func countNodes[K comparable, V any](tree *AVLTree[K, V]) uint32 {
	return _countNodes(tree, tree.root)
}

func _countNodes[K comparable, V any](tree *AVLTree[K, V], node *Node[K, V]) uint32 {
	if node == nil {
		return 0
	}
	return 1 + _countNodes(tree, node.left) + _countNodes(tree, node.right)
}

func validate[K comparable, V any](tree *AVLTree[K, V]) bool {
	return validateSize(tree) && validateBST(tree) && validateBalance(tree)
}

func validateSize[K comparable, V any](tree *AVLTree[K, V]) bool {
	return countNodes(tree) == tree.size
}

func validateBST[K comparable, V any](tree *AVLTree[K, V]) bool {
	return _validateBST(tree, tree.root)
}

func _validateBST[K comparable, V any](tree *AVLTree[K, V], node *Node[K, V]) bool {
	if node == nil {
		return true
	}
	if node.left != nil && tree.comparator(tree.max(node.left).key, node.key) > 0 {
		return false
	}
	if node.right != nil && tree.comparator(tree.min(node.right).key, node.key) < 0 {
		return false
	}
	return _validateBST(tree, node.left) && _validateBST(tree, node.right)
}

func validateBalance[K comparable, V any](tree *AVLTree[K, V]) bool {
	return _validateBalance(tree, tree.root)
}

func _validateBalance[K comparable, V any](tree *AVLTree[K, V], node *Node[K, V]) bool {
	if node == nil {
		return true
	}
	if tree.getBalance(node) > 1 || tree.getBalance(node) < -1 {
		return false
	}
	return _validateBalance(tree, node.left) && _validateBalance(tree, node.right)
}

func TestNew(t *testing.T) {
	tree := New[int, any]()
	if tree.root != nil {
		t.Error()
	}
	if tree.size != 0 {
		t.Error()
	}
	if !validate(tree) {
		t.Error()
	}
}

func TestNewWith(t *testing.T) {
	tree := NewWith[User, any](UserComparator)
	if tree.root != nil {
		t.Error()
	}
	if tree.size != 0 {
		t.Error()
	}
	if !validate(tree) {
		t.Error()
	}
}

func TestPut(t *testing.T) {
	tree := New[int, any]()
	nums := rand.Perm(10000)
	for _, num := range nums {
		tree.Put(num, struct{}{})
		if !validate(tree) {
			t.Error()
			return
		}
	}
}

func TestRemove(t *testing.T) {
	tree := New[int, any]()
	nums := rand.Perm(10000)
	for _, num := range nums {
		tree.Put(num, struct{}{})
	}

	nums = rand.Perm(10000)
	for _, num := range nums {
		tree.Remove(num)
		if !validate(tree) {
			t.Error()
			return
		}
	}
}

func TestGet(t *testing.T) {
	tree := New[int, any]()
	nums := rand.Perm(10000)
	for _, num := range nums {
		tree.Put(num, struct{}{})
		if _, ok := tree.Get(num); !ok {
			t.Error()
		}
	}

	nums = rand.Perm(10000)
	for _, num := range nums {
		tree.Remove(num)
		if _, ok := tree.Get(num); ok {
			t.Error()
		}
	}
}

func benchmarkPut(b *testing.B, tree *AVLTree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, tree *AVLTree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Remove(n)
		}
	}
}

func benchmarkGet(b *testing.B, tree *AVLTree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Get(n)
		}
	}
}

func BenchmarkPut1000(b *testing.B) {
	size := 1000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}

	b.ResetTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRemove1000(b *testing.B) {
	size := 1000
	tree := New[int, struct{}]()
	for i := 0; i < size; i++ {
		tree.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkGet1000(b *testing.B) {
	size := 1000
	tree := New[int, struct{}]()
	for i := 0; i < size; i++ {
		tree.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkPut10000(b *testing.B) {
	size := 10000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}

	b.ResetTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRemove10000(b *testing.B) {
	size := 10000
	tree := New[int, struct{}]()
	for i := 0; i < size; i++ {
		tree.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkGet10000(b *testing.B) {
	size := 10000
	tree := New[int, struct{}]()
	for i := 0; i < size; i++ {
		tree.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkPut100000(b *testing.B) {
	size := 100000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}

	b.ResetTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRemove100000(b *testing.B) {
	size := 100000
	tree := New[int, struct{}]()
	for i := 0; i < size; i++ {
		tree.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkGet100000(b *testing.B) {
	size := 100000
	tree := New[int, struct{}]()
	for i := 0; i < size; i++ {
		tree.Put(i, struct{}{})
	}

	b.ResetTimer()
	benchmarkGet(b, tree, size)
}
