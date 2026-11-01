package main

import (
	"fmt"
	"sync"
)

type Node struct {
	x   float64
	in  []<-chan float64
	out []chan<- float64
}

func (n *Node) round() {
	for _, c := range n.out {
		c <- n.x
	}
	sum := n.x
	for _, c := range n.in {
		sum += <-c
	}
	n.x = sum / float64(len(n.in)+1)
}

func spread(nodes []*Node) float64 {
	lo, hi := nodes[0].x, nodes[0].x
	for _, node := range nodes {
		lo = min(lo, node.x)
		hi = max(hi, node.x)
	}
	return hi - lo
}

func MajorityVote(votes []int, edges [][2]int, eps float64) (int, float64, int) {
	nodes := make([]*Node, len(votes))
	for i, v := range votes {
		nodes[i] = &Node{x: float64(v)}
	}

	for _, e := range edges {
		c := make(chan float64, 1)
		nodes[e[0]].out = append(nodes[e[0]].out, c)
		nodes[e[1]].in = append(nodes[e[1]].in, c)
	}

	rounds := 0
	for spread(nodes) > eps {
		var wg sync.WaitGroup
		for _, node := range nodes {
			wg.Go(func() {
				node.round()
			})
		}
		wg.Wait()
		rounds++
	}

	avg := nodes[0].x
	result := 0
	if avg > 0.5 {
		result = 1
	}
	return result, avg, rounds
}

func main() {
	votes := []int{1, 0, 1, 1, 0, 0, 1}
	edges := [][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 0},
		{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}, {5, 0}, {6, 1},
	}
	r, avg, rounds := MajorityVote(votes, edges, 1e-9)
	fmt.Printf("media este %.4f dupa %d runde, rezultatul votului: %d\n", avg, rounds, r)

	votes2 := []int{1, 1, 0, 0, 0, 0}
	edges2 := [][2]int{
		{0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 3}, {3, 2},
		{3, 4}, {4, 3}, {4, 5}, {5, 4}, {5, 0}, {0, 5},
	}
	r2, avg2, rounds2 := MajorityVote(votes2, edges2, 1e-9)
	fmt.Printf("media este %.4f dupa %d runde, rezultatul votului: %d\n", avg2, rounds2, r2)
}
