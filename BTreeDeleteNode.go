package student

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root != nil {
		if root == node {
			right := root.Right
			root = root.Left
			root.Right = right
		}
		if root.Left != nil {
			root.Left = BTreeDeleteNode(root.Left, node)
		}
		if root.Right != nil {
			root.Right = BTreeDeleteNode(root.Right, node)
		}
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
