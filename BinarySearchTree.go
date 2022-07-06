package tree

func binaryAddNode(root *Node, val int) {
	if root == nil {
		root = &Node{val: val}
		return
	}
	if val < root.val {
		if root.left == nil {
			root.left = &Node{val: val}
		} else {
			binaryAddNode(root.left, val)
		}
	} else {
		if root.right == nil {
			root.right = &Node{val: val}
		} else {
			binaryAddNode(root.right, val)
		}
	}
}

func search(root *Node, val int) *int {
	if root == nil {
		return nil
	}
	if root.val == val {
		return &root.val
	}
	if val < root.val {
		if root.left == nil {
			return nil
		}
		if root.left.val == val {
			return &root.left.val
		} else {
			search(root.left, val)
		}
	} else {
		if root.right == nil {
			return nil
		}
		if root.right.val == val {
			return &root.right.val
		} else {
			search(root.right, val)
		}
	}
	return nil //TODO: Proper way of fixing this? Given the possible values of int, it is impossible for this to be reached.
}
