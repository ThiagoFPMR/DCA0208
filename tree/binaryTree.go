package tree

type node struct {
	value int
	left  *node
	right *node
}

type BinaryTree struct {
	root *node
}

// Initialize the BinaryTree.
func (tree *BinaryTree) Init() {
	tree.root = nil
}

// Internal function to calculate the height of the BinaryTree
func (tree *BinaryTree) recursiveHeight(current *node) int {
	if current == nil {
		return 0
	}
	leftHeight := tree.recursiveHeight(current.left)
	rightHeight := tree.recursiveHeight(current.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Return the height of the BinaryTree
func (tree *BinaryTree) Height() int {
	return tree.recursiveHeight(tree.root)
}

// Check if the BinaryTree is empty
func (tree *BinaryTree) IsEmpty() bool {
	return tree.root == nil
}

// Internal function to add a node to the BinaryTree
func (tree *BinaryTree) recursiveAdd(current *node, value int) {
	if value < current.value {
		if current.left == nil {
			current.left = &node{value: value, left: nil, right: nil}
		} else {
			tree.recursiveAdd(current.left, value)
		}
	} else {
		if current.right == nil {
			current.right = &node{value: value, left: nil, right: nil}
		} else {
			tree.recursiveAdd(current.right, value)
		}
	}
}

// Add an element to the BinaryTree
func (tree *BinaryTree) Add(value int) {
	if tree.IsEmpty() {
		tree.root = &node{value: value, left: nil, right: nil}
	} else {
		tree.recursiveAdd(tree.root, value)
	}
}
