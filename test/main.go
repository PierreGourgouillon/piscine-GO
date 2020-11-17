package main

import (
	"fmt"

	student ".."
	"github.com/01-edu/z01"
)

func main() {
	fmt.Println("----------------Exo1----------------")
	fmt.Println(student.Atoi("12345"))
	fmt.Println(student.Atoi("0000000012345"))
	fmt.Println(student.Atoi("012 345"))
	fmt.Println(student.Atoi("Hello World!"))
	fmt.Println(student.Atoi("+1234")
	fmt.Println(student.Atoi("-1234"))
	fmt.Println(student.Atoi("++1234"))
	fmt.Println(student.Atoi("--1234"))
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo2----------------")
	arg1 := 4
	arg2 := 3
	fmt.Println(student.RecursivePower(arg1, arg2))
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo3----------------")
	fmt.Println("PrintCombN")
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo4----------------")
	student.PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo5----------------")
	fmt.Println("doop")
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo6----------------")
	fmt.Println(student.AtoiBase("125", "0123456789"))
	fmt.Println(student.AtoiBase("1111101", "01"))
	fmt.Println(student.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(student.AtoiBase("uoi", "choumi"))
	fmt.Println(student.AtoiBase("bbbbbab", "-ab"))
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo7----------------")
	fmt.Println(student.SplitWhiteSpaces("Hello how are you?"))
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo8----------------")
	s := "HelloHAhowHAareHAyou?"
	fmt.Println(student.Split(s, "HA"))
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo9----------------")
	result := student.ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo10---------------")
	fmt.Println("RotateVowels")
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo11---------------")
	result1 := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	student.AdvancedSortWordArr(result1, student.Compare)
	fmt.Println(result1)
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo12---------------")
	fmt.Println("cat")
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo13---------------")
	fmt.Println("ztail")
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo14---------------")
	nbits := student.ActiveBits(7)
	fmt.Println(nbits)
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo15---------------")
	var link *student.NodeI
	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)
	PrintList(link)
	link = student.SortListInsert(link, -2)
	link = student.SortListInsert(link, 2)
	PrintList(link)
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo16---------------")
	var link3 *student.NodeI
	var link4 *student.NodeI
	link3 = listPushBack(link3, 3)
	link3 = listPushBack(link3, 5)
	link3 = listPushBack(link3, 7)
	link4 = listPushBack(link4, -2)
	link4 = listPushBack(link4, 9)
	PrintList(student.SortedListMerge(link4, link3))
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo17---------------")
	link5 := &student.List{}
	link6 := &student.List{}
	fmt.Println("----normal state----")
	student.ListPushBack(link6, 1)
	PrintList2(link6)
	student.ListRemoveIf(link6, 1)
	fmt.Println("------answer-----")
	PrintList2(link6)
	fmt.Println()
	fmt.Println("----normal state----")
	student.ListPushBack(link5, 1)
	student.ListPushBack(link5, "Hello")
	student.ListPushBack(link5, 1)
	student.ListPushBack(link5, "There")
	student.ListPushBack(link5, 1)
	student.ListPushBack(link5, 1)
	student.ListPushBack(link5, "How")
	student.ListPushBack(link5, 1)
	student.ListPushBack(link5, "are")
	student.ListPushBack(link5, "you")
	student.ListPushBack(link5, 1)
	PrintList2(link5)
	student.ListRemoveIf(link5, 1)
	fmt.Println("------answer-----")
	PrintList2(link5)
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo18---------------")
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node := student.BTreeSearchItem(root, "1")
	replacement := &student.TreeNode{Data: "3"}
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo19---------------")
	root2 := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root2, "1")
	student.BTreeInsertData(root2, "7")
	student.BTreeInsertData(root2, "5")
	student.BTreeApplyByLevel(root2, fmt.Println)
	fmt.Println("------------------------------------")

	fmt.Println("----------------Exo20---------------")
	root3 := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root3, "1")
	student.BTreeInsertData(root3, "7")
	student.BTreeInsertData(root3, "5")
	node2 := student.BTreeSearchItem(root3, "4")
	fmt.Println("Before delete:")
	student.BTreeApplyInorder(root3, fmt.Println)
	root = student.BTreeDeleteNode(root3, node2)
	fmt.Println("After delete:")
	student.BTreeApplyInorder(root3, fmt.Println)
	fmt.Println("------------------------------------")

}

func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func PrintList2(l *student.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
