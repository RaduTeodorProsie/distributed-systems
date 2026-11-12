package main

import (
	"fmt"
	"maps"
	"math/rand/v2"
	"sync"
)

type Info = map[int]int
type Node struct{
	token int
	id int
	results Info
	senders map[int]chan <- Info
	receivers map[int] <- chan Info
}

func unite(a, b map[int]int){
	for k, v := range b{
		a[k] = v;
	}
}

func (p *Node) round(){
	for _, c := range p.senders{
		c <- maps.Clone(p.results)
	}

	for _, res := range p.receivers{
		other := <- res
		unite(p.results, other)		
	}
}

func solve(n int, edges[][2]int, iterations int){
	nodes := make([]Node, 0, n)

	for i := range n{
		nodes = append(nodes, Node{
			i,
			i,
			make(Info),
			make(map[int]chan <- Info),
			make(map[int] <- chan Info),
		})

		nodes[i].results[i] = i;
	}

	for _, edge := range edges{
		a, b := edge[0], edge[1]
		path := make(chan Info, 1)
		nodes[a].senders[b] = path
		nodes[b].receivers[a] = path
	}
	
	for range iterations{
		e := edges[rand.IntN(len(edges))]
		delete(nodes[e[0]].senders, e[1])
		delete(nodes[e[1]].receivers, e[0])

		var wg sync.WaitGroup
		for i := range n{
			wg.Go(func(){
				nodes[i].round()
			})
		}

		wg.Wait();
	}

	result := nodes[0].results

	ok := true
	for i := range n{
		if !maps.Equal(result, nodes[i].results){
			ok = false
		}
	}

	if !ok{
		fmt.Println("Did not converge")
	} else{
		fmt.Printf("Converged to %v\n", result)
	}
}

func main(){
	edges := [][2]int{
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 2},
		{2, 0},
		{2, 1},
	}

	solve(3, edges, 3)
}
