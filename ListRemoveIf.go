package student

import "fmt"

func ListPushBack(l *List, data interface{}) { //Mettre string/int dans list
	k := &NodeL{Next: nil, Data: data}

	if l.Head == nil {
		l.Head = k
		l.Tail = k
	} else {
		l.Tail.Next = k
		l.Tail = k
	}
}

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
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

func ListRemoveIf(l *List, data_ref interface{}) { //Delete les doublons
	list := &List{}

	for l.Head != nil {
		if l.Head.Data != data_ref {
			ListPushBack(list, l.Head.Data)
		}
		l.Head = l.Head.Next
	}
	l.Head = list.Head
}

//FINI
