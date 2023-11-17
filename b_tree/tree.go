package b_tree

// B-Tree
type Tree struct {
	memory []*Node
	degree int
}

func NewBTree(degree, capacity int) *Tree {
	memory := make([]*Node, 0, capacity)          //optimization pre load memory
	memory = append(memory, NewNode(degree, nil)) //create empty node

	return &Tree{
		degree: degree,
		memory: memory,
	}
}

func (t *Tree) Insert(key []byte) {
	// node := t.memory[0]
	node := Search(t.memory[0], key, t.degree) //find working node
	node.InsertKey(key)                        //add key to this node !fix double find position

	if node.capacity > t.degree-1 { //Check if node full
		node.Split(t)
	}
}

func (t Tree) GetMemory() []*Node {
	return t.memory
}

func (t *Tree) SwapRoot() {
	t.memory[len(t.memory)-1], t.memory[0] = t.memory[0], t.memory[len(t.memory)-1]
}
