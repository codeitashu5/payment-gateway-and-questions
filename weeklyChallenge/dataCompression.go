package weeklyChallenge

import (
	"container/heap"
	"fmt"
)

// Node Single information
type Node struct {
	Left      *Node
	Right     *Node
	Data      string
	Frequency int
}

func CreateNodeUsingData(data string) map[string]int {
	mapFrequency := map[string]int{}
	for i := 0; i < len(data); i++ {
		_, ok := mapFrequency[data[i:i+1]]
		if !ok {
			mapFrequency[data[i:i+1]] = 1
		} else {
			mapFrequency[data[i:i+1]] += 1
		}
	}
	return mapFrequency
}

type PriorityQueue []Node

func (p *PriorityQueue) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func (p *PriorityQueue) Push(node any) {
	*p = append(*p, node.(Node))
}

func (p *PriorityQueue) Len() int {
	return len(*p)
}

func (p *PriorityQueue) Less(i, j int) bool {
	arr := *p
	return arr[i].Frequency < arr[j].Frequency
}

func (p *PriorityQueue) Swap(i, j int) {
	temp := (*p)[i]
	(*p)[i] = (*p)[j]
	(*p)[j] = temp
}

func CreateCode(data string) {
	// we take the two lowest priority element and add there result to the heep
	minPriorityQueue := PriorityQueue{}
	mapCharFrequency := CreateNodeUsingData(data)

	// loop over the map and add the nodes in the priority queue
	for key, val := range mapCharFrequency {
		node := Node{nil, nil, key, val}
		heap.Push(&minPriorityQueue, node)
	}

	// use the heep function for consistent behaviour [creation of the huffman tree]
	for len(minPriorityQueue) > 1 {
		left := heap.Pop(&minPriorityQueue)
		right := heap.Pop(&minPriorityQueue)

		leftNode := left.(Node)
		rightNode := right.(Node)

		newNode := Node{Left: &leftNode, Right: &rightNode, Frequency: leftNode.Frequency + rightNode.Frequency}
		heap.Push(&minPriorityQueue, newNode)
	}

	root := heap.Pop(&minPriorityQueue)
	rootNode := root.(Node)
	mapCode := map[string][]byte{}
	// print and store the codes in the map
	printAndStoreInMap(&rootNode, &mapCode, []byte{})
	fmt.Println(mapCode)
	// compress the data using the mapCode
	compressedData := compressTheData(data, mapCode)
	fmt.Println(compressedData)

	// decompress the data
	decodedStr := decompression(&rootNode, compressedData)
	fmt.Println(decodedStr)

}

func decompression(root *Node, compressedData []byte) string {
	decodedStr := ""
	tempNode := root
	for i := 0; i < len(compressedData); i++ {
		if tempNode.Left == nil && tempNode.Right == nil {
			decodedStr = decodedStr + tempNode.Data
			tempNode = root
		}

		if compressedData[i] == 0 {
			tempNode = tempNode.Left
		}

		if compressedData[i] == 1 {
			tempNode = tempNode.Right
		}
	}
	return decodedStr + tempNode.Data
}

func compressTheData(data string, mapCode map[string][]byte) []byte {
	var compressedData []byte
	for i := 0; i < len(data); i++ {
		compressedData = append(compressedData, mapCode[data[i:i+1]]...)
	}
	return compressedData
}

func printAndStoreInMap(root *Node, mapCode *map[string][]byte, arrByte []byte) {
	if root.Left == nil && root.Right == nil {
		mapReq := *mapCode
		var newByteArr []byte
		newByteArr = append(newByteArr, arrByte...)
		mapReq[root.Data] = newByteArr
	}

	if root.Left != nil {
		arrByte = append(arrByte, 0)
		printAndStoreInMap(root.Left, mapCode, arrByte)
		arrByte = arrByte[0 : len(arrByte)-1]
	}

	if root.Right != nil {
		arrByte = append(arrByte, 1)
		printAndStoreInMap(root.Right, mapCode, arrByte)
		arrByte = arrByte[0 : len(arrByte)-1]
	}

}
