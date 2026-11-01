package main

import (
	"fmt"
	"sync"
)

type Node struct {
	id     int
	maxID  int
	in     []<-chan int
	out    []chan<- int
	leader bool
}

func (n *Node) run(rounds int) {
	for range rounds {
		for _, c := range n.out {
			c <- n.maxID
		}
		for _, c := range n.in {
			n.maxID = max(n.maxID, <-c)
		}
	}
	n.leader = n.maxID == n.id
}

func GraphFlooding(ids []int, edges [][2]int) int {
	n := len(ids)
	nodes := make([]*Node, n)
	for i, id := range ids {
		nodes[i] = &Node{id: id, maxID: id}
	}

	for _, e := range edges {
		c := make(chan int, 1)
		nodes[e[0]].out = append(nodes[e[0]].out, c)
		nodes[e[1]].in = append(nodes[e[1]].in, c)
	}

	var wg sync.WaitGroup
	for _, node := range nodes {
		wg.Go(func() {
			node.run(n)
		})
	}
	wg.Wait()

	leader := -1
	for i, node := range nodes {
		if node.leader {
			leader = i
		}
	}
	return leader
}

func main() {
	ids := []int{7, 3, 12, 5, 9, 1}
	edges := [][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 0},
		{0, 3}, {4, 1}, {2, 5},
	}
	l := GraphFlooding(ids, edges)
	fmt.Printf("liderul este nodul %d cu id-ul %d\n", l, ids[l])

	ids2 := []int{4, 10, 2, 8, 6}
	edges2 := [][2]int{
		{0, 1}, {1, 0}, {1, 2}, {2, 3}, {3, 1}, {3, 4}, {4, 0},
	}
	l2 := GraphFlooding(ids2, edges2)
	fmt.Printf("liderul este nodul %d cu id-ul %d\n", l2, ids2[l2])
}
