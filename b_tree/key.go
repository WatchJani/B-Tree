package b_tree

type Key struct {
	value []byte
	link  *Node
	// index []int
}

func NewKey(value []byte) Key {
	return Key{
		value: value,
		link:  nil,
		// index: make([]int, 0, 2), //capacity for index
	}
}

func (k *Key) UpdateLink(node *Node) {
	k.link = node
}

// func (k *Key) UpdateIndex(index int) {
// 	k.index = append(k.index, index)
// }
