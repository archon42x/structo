package avltree

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
	height uint32
}

type AVLTree[K comparable, V any] struct {
	root       *Node[K, V]
	comparator utils.Comparator[K]
	size       uint32
}

func New[K cmp.Ordered, V any]() *AVLTree[K, V] {
	return &AVLTree[K, V]{
		comparator: cmp.Compare[K],
	}
}

func NewWith[K comparable, V any](comparator utils.Comparator[K]) *AVLTree[K, V] {
	return &AVLTree[K, V]{
		comparator: comparator,
	}
}

func (tree *AVLTree[K, V]) Put(key K, value V) {
	node := tree.root
	if node == nil {
		tree.root = &Node[K, V]{
			key:    key,
			value:  value,
			height: 1,
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
					height: 1,
				}
				tree.size++
				tree.putFix(node)
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
					height: 1,
				}
				tree.size++
				tree.putFix(node)
				return
			} else {
				node = node.right
			}
		}
	}
}

func (tree *AVLTree[K, V]) Remove(key K) {
	node := tree.get(key)
	if node == nil {
		return
	}

	if node.left != nil && node.right != nil {
		next := tree.min(node.right)
		node.key = next.key
		node.value = next.value
		node = tree.replace(next, next.right)
	} else if node.left == nil && node.right == nil {
		node = tree.replace(node, nil)
	} else if node.left != nil {
		node = tree.replace(node, node.left)
	} else if node.right != nil {
		node = tree.replace(node, node.right)
	} else {
		panic("tree structure error")
	}
	tree.removeFix(node)
	tree.size--
}

func (tree *AVLTree[K, V]) Get(key K) (value V, ok bool) {
	node := tree.get(key)
	if node == nil {
		return value, false
	} else {
		return node.value, true
	}
}

func (tree *AVLTree[K, V]) putFix(node *Node[K, V]) {
	for node != nil {
		b := tree.getBalance(node)
		if b > 1 {
			b := tree.getBalance(node.left)
			if b > 0 {
				// LL
				if node.left == nil || node.left.left == nil {
					panic("tree structure error")
				}
				node = tree.rotateRight(node)
			} else if b < 0 {
				// LR
				if node.left == nil || node.left.right == nil {
					panic("tree structure error")
				}
				tree.rotateLeft(node.left)
				node = tree.rotateRight(node)
			} else {
				panic("tree structure error")
			}
		} else if b < -1 {
			b := tree.getBalance(node.right)
			if b > 0 {
				// RL
				if node.right == nil || node.right.left == nil {
					panic("tree structure error")
				}
				tree.rotateRight(node.right)
				node = tree.rotateLeft(node)
			} else if b < 0 {
				// RR
				if node.right == nil || node.right.right == nil {
					panic("tree structure error")
				}
				node = tree.rotateLeft(node)
			} else {
				panic("tree structure error")
			}
		} else {
			tree.calcHeight(node)
		}
		node = node.parent
	}
}

func (tree *AVLTree[K, V]) removeFix(node *Node[K, V]) {
	for node != nil {
		b := tree.getBalance(node)
		if b > 1 {
			b := tree.getBalance(node.left)
			if b >= 0 {
				// LL
				if node.left == nil || node.left.left == nil {
					panic("tree structure error")
				}
				node = tree.rotateRight(node)
			} else {
				// LR
				if node.left == nil || node.left.right == nil {
					panic("tree structure error")
				}
				tree.rotateLeft(node.left)
				node = tree.rotateRight(node)
			}
		} else if b < -1 {
			b := tree.getBalance(node.right)
			if b > 0 {
				// RL
				if node.right == nil || node.right.left == nil {
					panic("tree structure error")
				}
				tree.rotateRight(node.right)
				node = tree.rotateLeft(node)
			} else {
				// RR
				if node.right == nil || node.right.right == nil {
					panic("tree structure error")
				}
				node = tree.rotateLeft(node)
			}
		} else {
			tree.calcHeight(node)
		}
		node = node.parent
	}
}

func (tree *AVLTree[K, V]) get(key K) *Node[K, V] {
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
	return nil
}

func (tree *AVLTree[K, V]) replace(oldNode, newNode *Node[K, V]) *Node[K, V] {
	parent := oldNode.parent
	if parent == nil {
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
	return parent
}

func (tree *AVLTree[K, V]) getHeight(node *Node[K, V]) uint32 {
	if node == nil {
		return 0
	}
	return node.height
}

func (tree *AVLTree[K, V]) getBalance(node *Node[K, V]) int8 {
	return int8(tree.getHeight(node.left) - tree.getHeight(node.right))
}

func (tree *AVLTree[K, V]) calcHeight(node *Node[K, V]) {
	node.height = max(tree.getHeight(node.left), tree.getHeight(node.right)) + 1
}

func (tree *AVLTree[K, V]) min(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func (tree *AVLTree[K, V]) max(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	for node.right != nil {
		node = node.right
	}
	return node
}

func (tree *AVLTree[K, V]) rotateLeft(node *Node[K, V]) *Node[K, V] {
	if node.right == nil {
		panic("tree structure error")
	}
	child := node.right
	if node.parent == nil {
		// node is root
		child.parent = nil
		tree.root = child
	} else {
		parent := node.parent
		switch node {
		case parent.left:
			parent.left = child
		case parent.right:
			parent.right = child
		default:
			panic("tree structure error")
		}
		child.parent = parent
	}

	node.parent = child
	node.right = child.left
	if child.left != nil {
		child.left.parent = node
	}
	child.left = node

	tree.calcHeight(node)
	tree.calcHeight(node.parent)

	return child
}

func (tree *AVLTree[K, V]) rotateRight(node *Node[K, V]) *Node[K, V] {
	if node.left == nil {
		panic("tree structure error")
	}
	child := node.left
	if node.parent == nil {
		// node is root
		child.parent = nil
		tree.root = child
	} else {
		parent := node.parent
		switch node {
		case parent.left:
			parent.left = child
		case parent.right:
			parent.right = child
		default:
			panic("tree structure error")
		}
		child.parent = parent
	}

	node.parent = child
	node.left = child.right
	if child.right != nil {
		child.right.parent = node
	}
	child.right = node

	tree.calcHeight(node)
	tree.calcHeight(node.parent)

	return child
}
