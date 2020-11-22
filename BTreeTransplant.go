package student

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	if node.Data > root.Data {
		root.Right = BTreeTransplant(root.Right, node, rplc)
	} else if node.Data < root.Data {
		root.Left = BTreeTransplant(root.Left, node, rplc)
	}

	if node == root {
		root = rplc
	}
	return root
}

//FINI
