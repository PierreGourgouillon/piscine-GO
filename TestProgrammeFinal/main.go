package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	student ".."
	"github.com/01-edu/z01"
)

func main() {
	fmt.Println("----------------Atoi----------------")
	fmt.Println(student.Atoi(""))
	fmt.Println(student.Atoi("-"))
	fmt.Println(student.Atoi("--123"))
	fmt.Println(student.Atoi("1"))
	fmt.Println(student.Atoi("-3"))
	fmt.Println(student.Atoi("8292"))
	fmt.Println(student.Atoi("9223372036854775807"))
	fmt.Println(student.Atoi("-9223372036854775808"))
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------RecursivePower----------------")
	fmt.Println(student.RecursivePower(-7, -2))
	fmt.Println(student.RecursivePower(-8, -7))
	fmt.Println(student.RecursivePower(4, 8))
	fmt.Println(student.RecursivePower(1, 3))
	fmt.Println(student.RecursivePower(-1, 1))
	fmt.Println(student.RecursivePower(-6, 5))
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------PrintCombN----------------")
	student.PrintCombN(1)
	student.PrintCombN(2)
	student.PrintCombN(3)
	student.PrintCombN(4)
	student.PrintCombN(5)
	student.PrintCombN(6)
	student.PrintCombN(7)
	student.PrintCombN(8)
	student.PrintCombN(9)
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------PrintNbrBase----------------")
	student.PrintNbrBase(919617, "01")
	z01.PrintRune('\n')
	student.PrintNbrBase(753639, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	student.PrintNbrBase(-174336, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	student.PrintNbrBase(-661165, "1")
	z01.PrintRune('\n')
	student.PrintNbrBase(-861737, "Zone01Zone01")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "-ab")
	z01.PrintRune('\n')
	student.PrintNbrBase(-9223372036854775808, "0123456789")
	z01.PrintRune('\n')
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------Doop----------------")
	fmt.Println("Test à part")
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------AtoiBase----------------")
	fmt.Println(student.AtoiBase("bcbbbbaab", "abc"))
	fmt.Println(student.AtoiBase("0001", "01"))
	fmt.Println(student.AtoiBase("00", "01"))
	fmt.Println(student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	fmt.Println(student.AtoiBase("AAho?Ao", "WhoAmI?"))
	fmt.Println(student.AtoiBase("thisinputshouldnotmatter", "abca"))
	fmt.Println(student.AtoiBase("125", "0123456789"))
	fmt.Println(student.AtoiBase("uoi", "choumi"))
	fmt.Println(student.AtoiBase("bbbbbab", "-ab"))
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------SplitWhiteSpaces----------------")
	fmt.Println(student.SplitWhiteSpaces("The earliest foundations of what would become computer science predate the invention of the modern digital computer"))
	fmt.Println(student.SplitWhiteSpaces("Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"))
	fmt.Println(student.SplitWhiteSpaces("aiding in computations such as multiplication and division ."))
	fmt.Println(student.SplitWhiteSpaces("Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."))
	fmt.Println(student.SplitWhiteSpaces("Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"))
	fmt.Println(student.SplitWhiteSpaces("In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"))
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------Split----------------")
	fmt.Println(student.Split("|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,", "|=choumi=|"))
	fmt.Println(student.Split("!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,", "!==!"))
	fmt.Println(student.Split("AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,", "AFJ"))
	fmt.Println(student.Split("<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,", "<<==123==>>"))
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------ConvertBase----------------")
	fmt.Println(student.ConvertBase("4506C", "0123456789ABCDEF", "choumi"))
	fmt.Println(student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF"))
	fmt.Println(student.ConvertBase("256850", "0123456789", "01"))
	fmt.Println(student.ConvertBase("uuhoumo", "choumi", "Zone01"))
	fmt.Println(student.ConvertBase("683241", "0123456789", "0123456789"))
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------Rotatevowels---------------")
	fmt.Println("Test à part")
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------AdvancedSortWordArr---------------")
	result1 := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	student.AdvancedSortWordArr(result1, strings.Compare)
	fmt.Println(result1)
	result1 = []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	student.AdvancedSortWordArr(result1, strings.Compare)
	fmt.Println(result1)
	result1 = []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	student.AdvancedSortWordArr(result1, strings.Compare)
	fmt.Println(result1)
	result1 = []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	student.AdvancedSortWordArr(result1, func(a, b string) int { return strings.Compare(b, a) })
	fmt.Println(result1)
	result1 = []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	student.AdvancedSortWordArr(result1, func(a, b string) int { return strings.Compare(b, a) })
	fmt.Println(result1)
	result1 = []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	student.AdvancedSortWordArr(result1, func(a, b string) int { return strings.Compare(b, a) })
	fmt.Println(result1)
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------Cat---------------")
	fmt.Println("Test à part")
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------Ztail---------------")
	fmt.Println("Test à part")
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------ActiveBits---------------")
	fmt.Println(student.ActiveBits(15))
	fmt.Println(student.ActiveBits(17))
	fmt.Println(student.ActiveBits(4))
	fmt.Println(student.ActiveBits(11))
	fmt.Println(student.ActiveBits(9))
	fmt.Println(student.ActiveBits(12))
	fmt.Println(student.ActiveBits(2))
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------SortListInsert---------------")
	var lunk *student.NodeI
	lunk = listPushBack(lunk, 0)
	lunk = student.SortListInsert(lunk, 39)
	PrintList(lunk)

	var link *student.NodeI
	link = listPushBack(link, 0)
	link = listPushBack(link, 1)
	link = listPushBack(link, 2)
	link = listPushBack(link, 3)
	link = listPushBack(link, 4)
	link = listPushBack(link, 5)
	link = listPushBack(link, 24)
	link = listPushBack(link, 25)
	link = listPushBack(link, 54)
	link = student.SortListInsert(link, 33)
	PrintList(link)

	var link2 *student.NodeI
	link2 = listPushBack(link2, 0)
	link2 = listPushBack(link2, 2)
	link2 = listPushBack(link2, 18)
	link2 = listPushBack(link2, 33)
	link2 = listPushBack(link2, 37)
	link2 = listPushBack(link2, 39)
	link2 = listPushBack(link2, 52)
	link2 = listPushBack(link2, 53)
	link2 = listPushBack(link2, 57)
	link2 = student.SortListInsert(link2, 53)
	PrintList(link2)

	var link3 *student.NodeI
	link3 = listPushBack(link3, 0)
	link3 = listPushBack(link3, 5)
	link3 = listPushBack(link3, 18)
	link3 = listPushBack(link3, 24)
	link3 = listPushBack(link3, 28)
	link3 = listPushBack(link3, 35)
	link3 = listPushBack(link3, 42)
	link3 = listPushBack(link3, 45)
	link3 = listPushBack(link3, 52)
	link3 = student.SortListInsert(link3, 52)
	PrintList(link3)

	var link4 *student.NodeI
	link4 = listPushBack(link4, 0)
	link4 = listPushBack(link4, 12)
	link4 = listPushBack(link4, 20)
	link4 = listPushBack(link4, 23)
	link4 = listPushBack(link4, 23)
	link4 = listPushBack(link4, 24)
	link4 = listPushBack(link4, 30)
	link4 = listPushBack(link4, 41)
	link4 = listPushBack(link4, 53)
	link4 = listPushBack(link4, 57)
	link4 = listPushBack(link4, 59)
	link4 = student.SortListInsert(link4, 38)
	PrintList(link4)
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------SortListMerge---------------")
	var link5 *student.NodeI
	var link6 *student.NodeI
	PrintList(student.SortedListMerge(link5, link6))

	var link7 *student.NodeI
	link7 = listPushBack(link7, 2)
	link7 = listPushBack(link7, 2)
	link7 = listPushBack(link7, 4)
	link7 = listPushBack(link7, 9)
	link7 = listPushBack(link7, 12)
	link7 = listPushBack(link7, 12)
	link7 = listPushBack(link7, 19)
	link7 = listPushBack(link7, 20)
	PrintList(student.SortedListMerge(link5, link7))

	var link8 *student.NodeI
	link8 = listPushBack(link8, 4)
	link8 = listPushBack(link8, 4)
	link8 = listPushBack(link8, 6)
	link8 = listPushBack(link8, 9)
	link8 = listPushBack(link8, 13)
	link8 = listPushBack(link8, 18)
	link8 = listPushBack(link8, 20)
	link8 = listPushBack(link8, 20)
	PrintList(student.SortedListMerge(link8, link5))

	var link9 *student.NodeI
	link9 = listPushBack(link9, 0)
	link9 = listPushBack(link9, 7)
	link9 = listPushBack(link9, 39)
	link9 = listPushBack(link9, 92)
	link9 = listPushBack(link9, 97)
	link9 = listPushBack(link9, 93)
	link9 = listPushBack(link9, 91)
	link9 = listPushBack(link9, 28)
	link9 = listPushBack(link9, 64)
	var link10 *student.NodeI
	link10 = listPushBack(link10, 80)
	link10 = listPushBack(link10, 23)
	link10 = listPushBack(link10, 27)
	link10 = listPushBack(link10, 30)
	link10 = listPushBack(link10, 85)
	link10 = listPushBack(link10, 81)
	link10 = listPushBack(link10, 75)
	link10 = listPushBack(link10, 70)
	PrintList(student.SortedListMerge(link9, link10))

	var link11 *student.NodeI
	link11 = listPushBack(link11, 0)
	link11 = listPushBack(link11, 2)
	link11 = listPushBack(link11, 11)
	link11 = listPushBack(link11, 30)
	link11 = listPushBack(link11, 54)
	link11 = listPushBack(link11, 56)
	link11 = listPushBack(link11, 70)
	link11 = listPushBack(link11, 79)
	link11 = listPushBack(link11, 99)
	var link12 *student.NodeI
	link12 = listPushBack(link12, 23)
	link12 = listPushBack(link12, 28)
	link12 = listPushBack(link12, 38)
	link12 = listPushBack(link12, 67)
	link12 = listPushBack(link12, 67)
	link12 = listPushBack(link12, 79)
	link12 = listPushBack(link12, 95)
	link12 = listPushBack(link12, 97)
	PrintList(student.SortedListMerge(link11, link12))

	var link13 *student.NodeI
	link13 = listPushBack(link13, 0)
	link13 = listPushBack(link13, 3)
	link13 = listPushBack(link13, 8)
	link13 = listPushBack(link13, 8)
	link13 = listPushBack(link13, 13)
	link13 = listPushBack(link13, 19)
	link13 = listPushBack(link13, 34)
	link13 = listPushBack(link13, 38)
	link13 = listPushBack(link13, 46)
	var link14 *student.NodeI
	link14 = listPushBack(link14, 7)
	link14 = listPushBack(link14, 39)
	link14 = listPushBack(link14, 45)
	link14 = listPushBack(link14, 53)
	link14 = listPushBack(link14, 59)
	link14 = listPushBack(link14, 70)
	link14 = listPushBack(link14, 76)
	link14 = listPushBack(link14, 79)
	PrintList(student.SortedListMerge(link13, link14))
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------ListRemoveIf---------------")
	list := &student.List{}
	student.ListRemoveIf(list, 1)
	PrintList2(list)

	lust := &student.List{}
	student.ListRemoveIf(lust, 96)
	PrintList2(lust)

	list2 := &student.List{}
	student.ListPushBack(list2, 98)
	student.ListPushBack(list2, 98)
	student.ListPushBack(list2, 33)
	student.ListPushBack(list2, 34)
	student.ListPushBack(list2, 33)
	student.ListPushBack(list2, 34)
	student.ListPushBack(list2, 33)
	student.ListPushBack(list2, 89)
	student.ListPushBack(list2, 33)
	student.ListRemoveIf(list2, 34)
	PrintList2(list2)

	list3 := &student.List{}
	student.ListPushBack(list3, 79)
	student.ListPushBack(list3, 74)
	student.ListPushBack(list3, 99)
	student.ListPushBack(list3, 79)
	student.ListPushBack(list3, 7)
	student.ListRemoveIf(list3, 99)
	PrintList2(list3)

	list4 := &student.List{}
	student.ListPushBack(list4, 56)
	student.ListPushBack(list4, 93)
	student.ListPushBack(list4, 68)
	student.ListPushBack(list4, 56)
	student.ListPushBack(list4, 87)
	student.ListPushBack(list4, 68)
	student.ListPushBack(list4, 56)
	student.ListPushBack(list4, 68)
	student.ListRemoveIf(list4, 68)
	PrintList2(list4)

	list5 := &student.List{}
	student.ListPushBack(list5, "mvkUxbqhQve4l")
	student.ListPushBack(list5, "4Zc4t hnf SQ")
	student.ListPushBack(list5, "q2If E8BPuX")
	student.ListRemoveIf(list5, "4Zc4t hnf SQ")
	PrintList2(list5)
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------BTreeTransplant---------------")
	root := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root, "07")
	student.BTreeInsertData(root, "05")
	student.BTreeInsertData(root, "12")
	student.BTreeInsertData(root, "02")
	student.BTreeInsertData(root, "10")
	student.BTreeInsertData(root, "03")
	node := student.BTreeSearchItem(root, "12")
	replacement := &student.TreeNode{Data: "55"}
	student.BTreeInsertData(replacement, "33")
	student.BTreeInsertData(replacement, "60")
	student.BTreeInsertData(replacement, "12")
	student.BTreeInsertData(replacement, "15")
	root = student.BTreeTransplant(root, node, replacement)
	print(os.Stdout, root, 0, 'M')

	fmt.Println()

	root2 := &student.TreeNode{Data: "33"}
	student.BTreeInsertData(root2, "20")
	student.BTreeInsertData(root2, "5")
	student.BTreeInsertData(root2, "13")
	student.BTreeInsertData(root2, "52")
	student.BTreeInsertData(root2, "11")
	node2 := student.BTreeSearchItem(root2, "20")
	replacement2 := &student.TreeNode{Data: "55"}
	student.BTreeInsertData(replacement2, "33")
	student.BTreeInsertData(replacement2, "60")
	student.BTreeInsertData(replacement2, "12")
	student.BTreeInsertData(replacement2, "15")
	root2 = student.BTreeTransplant(root2, node2, replacement2)
	print(os.Stdout, root2, 0, 'M')

	fmt.Println()

	root3 := &student.TreeNode{Data: "03"}
	student.BTreeInsertData(root3, "39")
	student.BTreeInsertData(root3, "11")
	student.BTreeInsertData(root3, "99")
	student.BTreeInsertData(root3, "14")
	student.BTreeInsertData(root3, "44")
	student.BTreeInsertData(root3, "11")
	node3 := student.BTreeSearchItem(root3, "11")
	replacement3 := &student.TreeNode{Data: "55"}
	student.BTreeInsertData(replacement3, "33")
	student.BTreeInsertData(replacement3, "60")
	student.BTreeInsertData(replacement3, "12")
	student.BTreeInsertData(replacement3, "15")
	root3 = student.BTreeTransplant(root3, node3, replacement3)
	print(os.Stdout, root3, 0, 'M')
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------BTreeApplyByLevel---------------")
	root4 := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root4, "07")
	student.BTreeInsertData(root4, "05")
	student.BTreeInsertData(root4, "12")
	student.BTreeInsertData(root4, "02")
	student.BTreeInsertData(root4, "10")
	student.BTreeInsertData(root4, "03")
	student.BTreeApplyByLevel(root4, fmt.Println)

	fmt.Println()

	root5 := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root5, "07")
	student.BTreeInsertData(root5, "05")
	student.BTreeInsertData(root5, "12")
	student.BTreeInsertData(root5, "02")
	student.BTreeInsertData(root5, "10")
	student.BTreeInsertData(root5, "03")
	student.BTreeApplyByLevel(root5, fmt.Print)
	z01.PrintRune('\n')
	fmt.Println("------------------------------------")

	fmt.Println()

	fmt.Println("----------------BTreeDeleteNode---------------")
	root6 := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root6, "07")
	student.BTreeInsertData(root6, "05")
	student.BTreeInsertData(root6, "12")
	student.BTreeInsertData(root6, "02")
	student.BTreeInsertData(root6, "10")
	student.BTreeInsertData(root6, "03")
	node6 := student.BTreeSearchItem(root6, "02")
	root6 = student.BTreeDeleteNode(root6, node6)
	print(os.Stdout, root6, 0, 'M')

	fmt.Println()

	root7 := &student.TreeNode{Data: "33"}
	student.BTreeInsertData(root7, "20")
	student.BTreeInsertData(root7, "5")
	student.BTreeInsertData(root7, "13")
	student.BTreeInsertData(root7, "31")
	student.BTreeInsertData(root7, "52")
	student.BTreeInsertData(root7, "11")
	node7 := student.BTreeSearchItem(root7, "20")
	root7 = student.BTreeDeleteNode(root7, node7)
	print(os.Stdout, root7, 0, 'M')

	fmt.Println()

	root8 := &student.TreeNode{Data: "03"}
	student.BTreeInsertData(root8, "39")
	student.BTreeInsertData(root8, "11")
	student.BTreeInsertData(root8, "99")
	student.BTreeInsertData(root8, "14")
	student.BTreeInsertData(root8, "44")
	student.BTreeInsertData(root8, "11")
	node8 := student.BTreeSearchItem(root8, "03")
	root8 = student.BTreeDeleteNode(root8, node8)
	print(os.Stdout, root8, 0, 'M')

	fmt.Println()

	root9 := &student.TreeNode{Data: "03"}
	student.BTreeInsertData(root9, "01")
	student.BTreeInsertData(root9, "03")
	student.BTreeInsertData(root9, "94")
	student.BTreeInsertData(root9, "19")
	student.BTreeInsertData(root9, "111")
	student.BTreeInsertData(root9, "24")
	node9 := student.BTreeSearchItem(root9, "03")
	root9 = student.BTreeDeleteNode(root9, node9)
	print(os.Stdout, root9, 0, 'M')
	fmt.Println("------------------------------------")

}

func print(w io.Writer, node *student.TreeNode, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	print(w, node.Right, ns+2, 'R')
	print(w, node.Left, ns+2, 'L')

}

func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, "-> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func PrintList2(l *student.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, "-> ")
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
