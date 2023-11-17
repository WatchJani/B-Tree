package b_tree

import (
	"bytes"
)

// return index of value or best index position in slice
func SearchInsert(keys []Key, target []byte, capacity int) int { //capacity -> [[1][20][41][][][]] -> capacity = 3 element in array
	start := 0

	for start <= capacity {
		mid := (start + capacity) / 2
		compareResult := bytes.Compare(keys[mid].value, target)

		if compareResult == 0 {
			return -1
		} else if compareResult < 0 {
			start = mid + 1
		} else {
			capacity = mid - 1
		}
	}

	return start
}

// insert element, how can be sorted
// !try to write better
func ShiftAndInsert(arr *[]Key, newNumbers []Key, insertIndex int) {
	for i := len(*arr) - 1; i > insertIndex; i-- {
		(*arr)[i] = (*arr)[i-len(newNumbers)]
	}

	for i := 0; i < len(newNumbers); i++ {
		(*arr)[insertIndex+i] = newNumbers[i]
	}
}

func Search(node *Node, target []byte, degree int) *Node {
	for node != nil {
		index := SearchInsert(node.keys, target, node.capacity-1)
		if index < degree && node.keys[index].link != nil {
			node = node.keys[index].link
			// fmt.Println("key")
		} else if index == node.capacity && node.link != nil {
			node = node.link
			// fmt.Println("node")
		} else {
			break
		}
	}

	return node
}

func Append(keys []Key) []Key {
	key := make([]Key, 5)

	for index, value := range keys {
		key[index] = value
	}

	return key
}
