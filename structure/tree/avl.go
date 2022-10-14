package tree

import (
	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math/max"
)

type AVLTree[T constraints.Ordered] struct {
	*binaryTree[T]
}

// NewAVLTree create a novel AVL tree
func NewAVLTree[T constraints.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{
		binaryTree: &binaryTree[T]{
			Root: nil,
			NIL:  nil,
		},
	}
}

// Insert a chain of Node's into the AVL Tree
func (avl *AVLTree[T]) Insert(keys ...T) {
	for _, k := range keys {
		avl.Root = avl.insertHelper(avl.Root, k)
	}
}

// Delete a Node from the AVL Tree
func (avl *AVLTree[T]) Delete(key T) bool {
	tmp := avl.deleteHelper(avl.Root, key)
	if tmp == nil {
		return false
	}

	avl.Root = tmp
	return true
}

func (avl *AVLTree[T]) insertHelper(root *Node[T], key T) *Node[T] {
	if root == nil {
		return &Node[T]{
			Key:    key,
			Height: 1,
		}
	}

	switch {
	case key < root.Key:
		root.Left = avl.insertHelper(root.Left, key)
	case key > root.Key:
		root.Right = avl.insertHelper(root.Right, key)
	default:
		return root
	}

	// balance the tree
	root.Height = avl.height(root)
	bFactor := avl.balanceFactor(root)
	if bFactor > 1 {
		switch {
		case key < root.Left.Key:
			return avl.rightRotate(root)
		case key > root.Left.Key:
			root.Left = avl.leftRotate(root.Left)
			return avl.rightRotate(root)
		}
	}

	if bFactor < -1 {
		switch {
		case key > root.Right.Key:
			return avl.leftRotate(root)
		case key < root.Right.Key:
			root.Right = avl.rightRotate(root.Right)
			return avl.leftRotate(root)
		}
	}

	return root
}

func (avl *AVLTree[T]) deleteHelper(root *Node[T], key T) *Node[T] {
	if root == nil {
		return root
	}

	switch {
	case key < root.Key:
		root.Left = avl.deleteHelper(root.Left, key)
	case key > root.Key:
		root.Right = avl.deleteHelper(root.Right, key)
	default:
		if root.Left == nil || root.Right == nil {
			tmp := root.Left
			if root.Left != nil {
				tmp = root.Right
			}

			if tmp == nil {
				tmp = root
				root = nil
			} else {
				*root = *tmp
			}
		} else {
			tmp := avl.minimum(root.Right)
			root.Key = tmp.Key
			root.Right = avl.deleteHelper(root.Right, tmp.Key)
		}
	}

	if root == nil {
		return root
	}

	// balance the tree
	root.Height = avl.height(root)
	bFactor := avl.balanceFactor(root)
	switch {
	case bFactor > 1:
		switch {
		case avl.balanceFactor(root.Left) >= 0:
			return avl.rightRotate(root)
		default:
			root.Left = avl.leftRotate(root.Left)
			return avl.rightRotate(root)
		}
	case bFactor < -1:
		switch {
		case avl.balanceFactor(root.Right) <= 0:
			return avl.leftRotate(root)
		default:
			root.Right = avl.rightRotate(root.Right)
			return avl.leftRotate(root)
		}
	}

	return root
}

func (avl *AVLTree[T]) height(root *Node[T]) int {
	if root == nil {
		return 0
	}

	var leftHeight, rightHeight int
	if root.Left != nil {
		leftHeight = root.Left.Height
	}
	if root.Right != nil {
		rightHeight = root.Right.Height
	}
	return 1 + max.Int(leftHeight, rightHeight)
}

// balanceFactor : negative balance factor means subtree Root is heavy toward Left
// and positive balance factor means subtree Root is heavy toward Right side
func (avl *AVLTree[T]) balanceFactor(root *Node[T]) int {
	var leftHeight, rightHeight int
	if root.Left != nil {
		leftHeight = root.Left.Height
	}
	if root.Right != nil {
		rightHeight = root.Right.Height
	}
	return leftHeight - rightHeight
}

func (avl *AVLTree[T]) leftRotate(root *Node[T]) *Node[T] {
	y := root.Right
	yl := y.Left
	y.Left = root
	root.Right = yl

	root.Height = avl.height(root)
	y.Height = avl.height(y)
	return y
}

func (avl *AVLTree[T]) rightRotate(root *Node[T]) *Node[T] {
	y := root.Left
	yr := y.Right
	y.Right = root
	root.Left = yr

	root.Height = avl.height(root)
	y.Height = avl.height(y)
	return y
}
