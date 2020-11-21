package student

import (
	"fmt"
)

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
		root.Data = rplc.Data
	}
	return root
}

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{
			Left:   nil,
			Right:  nil,
			Parent: nil,
			Data:   data,
		}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else if data >= root.Data {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
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

//FINI
