package main

import "fmt"

func main() {
	var link *NodeI
	var link2 *NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(SortedListMerge(link2, link))
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1 = test(n1, n2)
	n1 = SortListInsert(n1)
	return n1
}

func test(n1 *NodeI, n2 *NodeI) *NodeI {

	for i := n2; i != nil; i = i.Next {
		n1 = listPushBack(n1, i.Data)
	}
	return n1
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI { //Mettre int dans list
	n := &NodeI{Data: data}

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

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type NodeI struct {
	Data int
	Next *NodeI
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func SortListInsert(l *NodeI) *NodeI { //Ranger dans l'ordre
	if l == nil {
		return l
	}
	var garage int

	for i := l; i.Next != nil; i = i.Next {
		for j := l; j.Next != nil; j = j.Next {
			if j.Data > j.Next.Data {
				garage = j.Next.Data
				j.Next.Data = j.Data
				j.Data = garage
			}
		}
	}
	return l
}

//FINI
