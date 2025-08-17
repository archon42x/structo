package bstree

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

func countNodes[K comparable, V any](tree *BSTree[K, V]) uint32 {
	return _countNodes(tree, tree.root)
}

func _countNodes[K comparable, V any](tree *BSTree[K, V], node *Node[K, V]) uint32 {
	if node == nil {
		return 0
	}
	return 1 + _countNodes(tree, node.left) + _countNodes(tree, node.right)
}

func validate[K comparable, V any](tree *BSTree[K, V]) bool {
	return validateSize(tree) && validateBST(tree)
}

func validateSize[K comparable, V any](tree *BSTree[K, V]) bool {
	return countNodes(tree) == tree.size
}

func validateBST[K comparable, V any](tree *BSTree[K, V]) bool {
	return _validateBST(tree, tree.root)
}

func _validateBST[K comparable, V any](tree *BSTree[K, V], node *Node[K, V]) bool {
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
		}
	}
}

func TestPutUser(t *testing.T) {
	tree := NewWith[User, any](UserComparator)
	names := []string{"Jack", "Bob", "Helen"}
	ages := make([]int, 100)
	for i := 0; i < len(ages); i++ {
		ages[i] = i
	}

	for i := 0; i < 300; i++ {
		tree.Put(User{
			name: names[rand.Intn(len(names))],
			age:  uint32(ages[rand.Intn(len(ages))]),
		}, struct{}{})
		if !validate(tree) {
			t.Error()
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
		}
	}
}

func TestRemoveUser(t *testing.T) {
	tree := NewWith[User, any](UserComparator)
	names := []string{"Jack", "Bob", "Helen"}
	ages := make([]int, 100)
	for i := 0; i < len(ages); i++ {
		ages[i] = i
	}

	for i := 0; i < 300; i++ {
		tree.Put(User{
			name: names[rand.Intn(len(names))],
			age:  uint32(ages[rand.Intn(len(ages))]),
		}, struct{}{})
	}

	for i := 0; i < 300; i++ {
		tree.Remove(User{
			name: names[rand.Intn(len(names))],
			age:  uint32(ages[rand.Intn(len(ages))]),
		})
		if !validate(tree) {
			t.Error()
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

func BenchmarkPut(b *testing.B) {
	tree := New[int, any]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := rand.Int()
		tree.Put(key, struct{}{})
	}
}

func BenchmarkGet(b *testing.B) {
	tree := New[int, struct{}]()

	keys := make([]int, b.N)
	for i := 0; i < len(keys); i++ {
		keys[i] = rand.Int()
		tree.Put(keys[i], struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := keys[rand.Intn(len(keys))]
		tree.Get(key)
	}
}

func BenchmarkRemove(b *testing.B) {
	tree := New[int, struct{}]()

	keys := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		tree.Put(keys[i], struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Remove(keys[i])
	}
}
