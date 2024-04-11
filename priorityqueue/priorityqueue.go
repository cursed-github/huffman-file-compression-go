package priorityqueue

import "container/heap"

type HuffmanNode struct {
	Frequency int
	Char rune
	Left *HuffmanNode
	Right *HuffmanNode
	CharEncoding string
}

// PriorityQueue implements heap.Interface and holds HoffmanNodes
type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { 
	return len(pq) 
}

func (pq PriorityQueue) Less(i, j int) bool {
	// Min-Heap based on frequency, change > to < for Max-Heap
	return pq[i].Frequency < pq[j].Frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node:= x.(*HuffmanNode)
	*pq=append(*pq,node)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
    node := (*pq)[n-1] 
    *pq = (*pq)[0 : n-1] 
    return node
}

func (pq *PriorityQueue) Init() {
	heap.Init(pq)
}

func (pq *PriorityQueue) InsertNode(node *HuffmanNode){
	pq.Push(node)
}

func (pq *PriorityQueue) ExtractMinimum() *HuffmanNode {
	if len(*pq)==0 {
	return nil
	}

	node := heap.Pop(pq).(*HuffmanNode)
	return node
}

