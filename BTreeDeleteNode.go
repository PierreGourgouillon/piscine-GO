package student

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Data == node.Data {
		if root.Right == nil {
			root = root.Left
		} else if root.Left == nil {
			root = root.Right
		} else if root.Left != nil && root.Right != nil {
			root.Data = BTreeMin(root.Right).Data
			root.Right = BTreeDeleteNode(root.Right, BTreeMin(root.Right))
			return root
		}
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	}
	return root
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root.Left != nil {
		return BTreeMin(root.Left)
	} else if root.Left == nil {
		return root
	}
	return root
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}
	f(root.Data)
	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	} else {
		return BTreeSearchItem(root.Left, elem)
	}
}

//FINISH
