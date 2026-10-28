package main

import(
	"fmt"
	"sync"
)

type Node struct{
	val int
	best int
	next chan <- int
	prev <- chan int
	leader bool
	Done bool
}

func (n *Node) operate(listen bool){
	for !n.Done {
		if listen{
			g, ok := <-n.prev
			if !ok{
				n.Done = true
				close(n.next)
			} else{
				n.best = max(n.best, g)
			}
		} else{
			if n.best < n.val{
				n.next <- n.val
			} else{
				n.next <- n.best
			}
		}

		if n.best == n.val  {
			n.leader = true
			close(n.next)
			n.Done = true
			for i := range n.prev{
				i = i + 1
			}
		}
		
		listen = !listen
	}
}

func main(){
	var nodes []Node
	var chans []chan int
	for range 10{
		chans = append(chans, make(chan int))
	}
	
	for i := range 10{
		nodes = append(nodes, Node{i, -1, chans[i], chans[((i - 1) + 10) % 10], false, false})
	}

	var wg sync.WaitGroup
	for i := range 10{
		wg.Go(func(){
			nodes[i].operate(i == 0)
		})
	}

	wg.Wait()
	for i, v := range nodes{
		if v.leader{
			fmt.Printf("the leader is %d\n", i)
		}
	}
}
