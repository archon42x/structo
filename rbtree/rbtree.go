package rbtree

import (
	"cmp"

	"github.com/archon42x/structo/utils"
)

type Color bool

const (
	black, red Color = true, false
)

type Node[K comparable, V any] struct {
	key    K
	value  V
	left   *Node[K, V]
	right  *Node[K, V]
	parent *Node[K, V]
	color  Color
}

type RBTree[K comparable, V any] struct {
	root       *Node[K, V]
	comparator utils.Comparator[K]
	size       uint32
}

func New[K cmp.Ordered, V any]() *RBTree[K, V] {
	return &RBTree[K, V]{
		comparator: cmp.Compare[K],
	}
}

func NewWith[K comparable, V any](comparator utils.Comparator[K]) *RBTree[K, V] {
	return &RBTree[K, V]{
		comparator: comparator,
	}
}

func (tree *RBTree[K, V]) Put(key K, value V) {
	node := tree.root
	if node == nil {
		tree.root = &Node[K, V]{
			key:   key,
			value: value,
			color: black,
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
					color:  red,
				}
				tree.putFix(node.left)
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
					color:  red,
				}
				tree.putFix(node.right)
				tree.size++
				return
			} else {
				node = node.right
			}
		}
	}
}

func (tree *RBTree[K, V]) Remove(key K) {
	node := tree.get(key)
	if node == nil {
		return
	}

	if node.left != nil && node.right != nil {
		next := tree.min(node.right)
		tree.removeFix(next)
		node.key = next.key
		node.value = next.value
		tree.replace(next, next.right)
	} else if node.left == nil && node.right == nil {
		tree.removeFix(node)
		tree.replace(node, nil)
	} else if node.left != nil {
		tree.removeFix(node)
		tree.replace(node, node.left)
	} else if node.right != nil {
		tree.removeFix(node)
		tree.replace(node, node.right)
	} else {
		panic("tree structure error")
	}
	if tree.root != nil {
		tree.root.color = black
	}
	tree.size--
}

func (tree *RBTree[K, V]) Get(key K) (value V, ok bool) {
	node := tree.get(key)
	if node == nil {
		return value, false
	} else {
		return node.value, true
	}
}

func (tree *RBTree[K, V]) putFix(node *Node[K, V]) {
	// node is root
	if node.parent == nil {
		node.color = black
		return
	}

	parent := node.parent
	// parent is black, dont need fix
	if tree.getColor(parent) == black {
		return
	}

	// parent is red
	if parent.parent == nil {
		panic("tree structure error")
	}
	// grandparent is black
	grandparent := parent.parent

	switch parent {
	case grandparent.left:
		uncle := grandparent.right
		if tree.getColor(uncle) == red {
			// uncle is red
			parent.color = black
			grandparent.color = red
			if uncle != nil {
				uncle.color = black
			}
			tree.putFix(grandparent)
		} else {
			// uncle is balck
			if node == parent.right {
				tree.rotateLeft(parent)
				node, parent = parent, node
			}
			parent.color = black
			grandparent.color = red
			tree.rotateRight(grandparent)
		}
	case grandparent.right:
		uncle := grandparent.left
		if tree.getColor(uncle) == red {
			// uncle is red
			parent.color = black
			grandparent.color = red
			if uncle != nil {
				uncle.color = black
			}
			tree.putFix(grandparent)
		} else {
			// uncle is balck
			if node == parent.left {
				tree.rotateRight(parent)
				node, parent = parent, node
			}
			parent.color = black
			grandparent.color = red
			tree.rotateLeft(grandparent)
		}
	default:
		panic("tree structure error")
	}
}

func (tree *RBTree[K, V]) removeFix(node *Node[K, V]) {
	if tree.getColor(node) == red {
		return
	}
	parent := node.parent
	if parent == nil {
		return
	}
	switch node {
	case parent.left:
		sibling := parent.right
		if sibling == nil {
			panic("tree structure error")
		}
		if tree.getColor(sibling) == red {
			// sibling is red
			parent.color = red
			sibling.color = black
			tree.rotateLeft(parent)
			tree.removeFix(node)
			return
		} else {
			// sibling is black
			inner := sibling.left
			outer := sibling.right
			if tree.getColor(inner) == black && tree.getColor(outer) == black {
				sibling.color = red
				if tree.getColor(parent) == red {
					parent.color = black
				} else {
					tree.removeFix(parent)
				}
			} else {
				if tree.getColor(outer) == black {
					// inner is red
					inner.color = black
					sibling.color = red
					tree.rotateRight(sibling)
					sibling, outer = inner, sibling
				}
				// outer is red
				sibling.color = parent.color
				parent.color = black
				outer.color = black
				tree.rotateLeft(parent)
			}
		}
	case parent.right:
		sibling := parent.left
		if sibling == nil {
			panic("tree structure error")
		}
		if tree.getColor(sibling) == red {
			parent.color = red
			sibling.color = black
			tree.rotateRight(parent)
			tree.removeFix(node)
			return
		} else {
			// sibling is black
			inner := sibling.right
			outer := sibling.left
			if tree.getColor(inner) == black && tree.getColor(outer) == black {
				sibling.color = red
				if tree.getColor(parent) == red {
					parent.color = black
				} else {
					tree.removeFix(parent)
				}
			} else {
				if tree.getColor(outer) == black {
					// inner is red
					inner.color = black
					sibling.color = red
					tree.rotateLeft(sibling)
					sibling, outer = inner, sibling
				}
				// outer is red
				sibling.color = parent.color
				parent.color = black
				outer.color = black
				tree.rotateRight(parent)
			}
		}
	default:
		panic("tree structure error")
	}
}

func (tree *RBTree[K, V]) getColor(node *Node[K, V]) Color {
	if node == nil {
		return black
	}
	return node.color
}

func (tree *RBTree[K, V]) get(key K) *Node[K, V] {
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

func (tree *RBTree[K, V]) replace(oldNode, newNode *Node[K, V]) {
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

func (tree *RBTree[K, V]) min(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func (tree *RBTree[K, V]) max(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	for node.right != nil {
		node = node.right
	}
	return node
}

func (tree *RBTree[K, V]) rotateLeft(node *Node[K, V]) {
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
}

func (tree *RBTree[K, V]) rotateRight(node *Node[K, V]) {
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
}
