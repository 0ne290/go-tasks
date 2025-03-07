package internal

type BinarySearchTree struct {
	root *node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{nil}
}

func (tree *BinarySearchTree) IsEmpty() bool {
	return tree.root == nil
}

func (tree *BinarySearchTree) NodeLeftRight(callback func(int)) {
	var execute func(current *node)
	execute = func(current *node) {
		for current != nil {
			callback(current.value)

			execute(current.left)

			current = current.right
		}
	}

	execute(tree.root)
}

func (tree *BinarySearchTree) LeftNodeRight(callback func(int)) {
	var execute func(current *node)
	execute = func(current *node) {
		for current != nil {
			execute(current.left)

			callback(current.value)

			current = current.right
		}
	}

	execute(tree.root)
}

func (tree *BinarySearchTree) LeftRightNode(callback func(int)) {
	var execute func(current *node)
	execute = func(current *node) {
		if current == nil {
			return
		}

		execute(current.left)

		execute(current.right)

		callback(current.value)
	}

	execute(tree.root)
}

func (tree *BinarySearchTree) Add(value int) {
	if tree.IsEmpty() {
		tree.root = &node{value, nil, nil}

		return
	}

	current := tree.root
	for {
		if value > current.value {
			if current.right == nil {
				current.right = &node{value, nil, nil}

				return
			}

			current = current.right
		} else {
			if current.left == nil {
				current.left = &node{value, nil, nil}

				return
			}

			current = current.left
		}
	}
}

func (tree *BinarySearchTree) Remove(value int) bool {
	current := tree.root
	var prev *node
	for {
		if current == nil {
			return false
		}

		if value == current.value {
			if current.left == nil && current.right == nil {
				if current == tree.root {
					tree.root = nil
				} else if prev.left == current {
					prev.left = nil
				} else {
					prev.right = nil
				}
			} else if current.left != nil && current.right == nil {
				if current == tree.root {
					tree.root = current.left
				} else if prev.left == current {
					prev.left = current.left
				} else {
					prev.right = current.left
				}
			} else if current.left == nil && current.right != nil {
				if current == tree.root {
					tree.root = current.right
				} else if prev.left == current {
					prev.left = current.right
				} else {
					prev.right = current.right
				}
			} else {
				remove := current
				prev = current
				current = current.right
				if current.left == nil {
					prev.right = current.right
				} else {
					for current.left != nil {
						prev = current
						current = current.left
					}

					prev.left = current.right
				}

				remove.value = current.value
			}

			return true
		}

		prev = current
		if value > current.value {
			current = current.right
		} else {
			current = current.left
		}
	}
}

func (tree *BinarySearchTree) Contains(value int) bool {
	current := tree.root
	for current != nil {
		if value == current.value {
			return true
		}

		if value > current.value {
			current = current.right
		} else {
			current = current.left
		}
	}

	return false
}