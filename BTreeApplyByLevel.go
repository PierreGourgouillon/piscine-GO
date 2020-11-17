package student

func Etage(root *TreeNode, i int, f func(...interface{}) (int, error)) {

	if root == nil {
		return
	}
	if i == 1 {
		f(root.Data)
	} else if i > 1 {
		Etage(root.Left, i-1, f)
		Etage(root.Right, i-1, f)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {

	for i := 1; i <= BTreeLevelCount(root); i++ {
		Etage(root, i, f)
	}
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

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeLevelCount(root *TreeNode) int {
	return levelCounter(root)
}

func levelCounter(root *TreeNode) int {
	count := 0
	if root == nil {
		return count
	}
	leftCount := levelCounter(root.Left)
	rightCount := levelCounter(root.Right)
	if leftCount > rightCount {
		return leftCount + 1
	}
	return rightCount + 1
}

//FINISH
