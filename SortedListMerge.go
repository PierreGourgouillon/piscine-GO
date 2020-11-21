package student

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1 = test(n1, n2)
	n1 = SortList(n1)
	return n1
}

func test(n1 *NodeI, n2 *NodeI) *NodeI {

	for i := n2; i != nil; i = i.Next {
		n1 = listPushBack(n1, i.Data)
	}
	return n1
}

func SortList(l *NodeI) *NodeI {
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
