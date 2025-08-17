package bstree

import (
	"cmp"

	"github.com/archon42x/structo/utils"
)

type Node[K comparable, V any] struct {
	key    K
	value  V
	left   *Node[K, V]
	right  *Node[K, V]
	parent *Node[K, V]
}

type BSTree[K comparable, V any] struct {
	root       *Node[K, V]
	comparator utils.Comparator[K]
	size       uint32
}

func New[K cmp.Ordered, V any]() *BSTree[K, V] {
	return &BSTree[K, V]{
		comparator: cmp.Compare[K],
	}
}

func NewWith[K comparable, V any](comparator utils.Comparator[K]) *BSTree[K, V] {
	return &BSTree[K, V]{
		comparator: comparator,
	}
}

func (tree *BSTree[K, V]) Put(key K, value V) {
	node := tree.root
	if node == nil {
		tree.root = &Node[K, V]{
			key:   key,
			value: value,
		}
		tree.size++
		return
	}
	for {
		compare := tree.comparator(key, node.key)
		switch {
		case compare == 0:
			node.key = key
			node.value = value
			return
		case compare < 0:
			if node.left == nil {
				node.left = &Node[K, V]{
					key:    key,
					value:  value,
					parent: node,
				}
				tree.size++
				return
			} else {
				node = node.left
			}
		case compare > 0:
			if node.right == nil {
				node.right = &Node[K, V]{
					key:    key,
					value:  value,
					parent: node,
				}
				tree.size++
				return
			} else {
				node = node.right
			}
		}
	}
}

func (tree *BSTree[K, V]) Remove(key K) {
	node := tree.get(key)
	if node == nil {
		return
	}

	if node.left != nil && node.right != nil {
		next := tree.min(node.right)
		node.key = next.key
		node.value = next.value
		tree.replace(next, next.right)
	} else if node.left == nil && node.right == nil {
		tree.replace(node, nil)
	} else if node.left != nil {
		tree.replace(node, node.left)
	} else if node.right != nil {
		tree.replace(node, node.right)
	} else {
		panic("tree structure error")
	}

	tree.size--
}

func (tree *BSTree[K, V]) Get(key K) (value V, ok bool) {
	node := tree.get(key)
	if node == nil {
		return value, false
	} else {
		return node.value, true
	}
}

func (tree *BSTree[K, V]) replace(oldNode, newNode *Node[K, V]) {
	if oldNode.parent == nil {
		tree.root = newNode
		if newNode != nil {
			newNode.parent = nil
		}
		oldNode.left = nil
		oldNode.right = nil
	} else {
		switch oldNode {
		case oldNode.parent.left:
			oldNode.parent.left = newNode
			if newNode != nil {
				newNode.parent = oldNode.parent
			}
			oldNode.parent = nil
			oldNode.left = nil
			oldNode.right = nil
		case oldNode.parent.right:
			oldNode.parent.right = newNode
			if newNode != nil {
				newNode.parent = oldNode.parent
			}
			oldNode.parent = nil
			oldNode.left = nil
			oldNode.right = nil
		default:
			panic("tree structure error")
		}
	}
}

func (tree *BSTree[K, V]) get(key K) *Node[K, V] {
	node := tree.root
	for node != nil {
		compare := tree.comparator(key, node.key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.left
		case compare > 0:
			node = node.right
		}
	}
	return node
}

func (tree *BSTree[K, V]) min(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func (tree *BSTree[K, V]) max(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	for node.right != nil {
		node = node.right
	}
	return node
}
