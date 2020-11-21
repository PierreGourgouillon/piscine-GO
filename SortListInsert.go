package student

func listPushBack(l *NodeI, data int) *NodeI {
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

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		return l
	}

	l = listPushBack(l, data_ref)
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
