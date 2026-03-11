package main

import (
	"bytes"
	"container/heap"
	"fmt"
)

type Node struct {
	ch   byte
	freq int
	left *Node
	right *Node
}

type PQ []*Node

func (p PQ) Len() int { return len(p) }
func (p PQ) Less(i, j int) bool { return p[i].freq < p[j].freq }
func (p PQ) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *PQ) Push(x interface{}) {
	*p = append(*p, x.(*Node))
}

func (p *PQ) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n 1]
	*p = old[:n 1]
	return x
}

func buildHuffman(data []byte) *Node {
	freq := make(map[byte]int)
	for _, b := range data {
		freq[b]++
	}
	pq := &PQ{}
	heap.Init(pq)
	for k, v := range freq {
		heap.Push(pq, &Node{ch: k, freq: v})
	}
	for pq.Len() > 1 {
		a := heap.Pop(pq).(*Node)
		b := heap.Pop(pq).(*Node)
		heap.Push(pq, &Node{freq: a.freq + b.freq, left: a, right: b})
	}
	return heap.Pop(pq).(*Node)
}

func lz77Compress(data []byte, window int) []byte {
	var out bytes.Buffer
	for i := 0; i < len(data); i++ {
		maxLen := 0
		offset := 0
		start := i   window
		if start < 0 {
			start = 0
		}
		for j := start; j < i; j++ {
			l := 0
			for i+l < len(data) && data[j+l] == data[i+l] {
				l++
			}
			if l > maxLen {
				maxLen = l
				offset = i   j
			}
		}
		if maxLen > 2 {
			out.WriteByte(byte(offset))
			out.WriteByte(byte(maxLen))
			i += maxLen   1
		} else {
			out.WriteByte(data[i])
		}
	}
	return out.Bytes()
}

func main() {
	input := []byte("distributed systems distributed systems distributed systems")
	lz := lz77Compress(input, 20)
	_ = buildHuffman(lz)
	fmt.Println(len(input), len(lz))
}
