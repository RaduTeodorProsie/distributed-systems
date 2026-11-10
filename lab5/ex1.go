package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

const LIM_ROUNDS = 1000
const EPS = 1e-6

type Point [2]float64

type Robot struct {
	pos   Point
	inbox chan Point
	out   []chan<- Point
}

func (r *Robot) step() {
	for _, c := range r.out {
		c <- r.pos
	}
	sum := r.pos
	for range r.out {
		p := <-r.inbox
		sum[0] += p[0]
		sum[1] += p[1]
	}
	k := float64(1 + len(r.out))
	r.pos = Point{sum[0] / k, sum[1] / k}
}

func dist(a, b *Robot) float64 {
	dx := a.pos[0] - b.pos[0]
	dy := a.pos[1] - b.pos[1]
	return math.Sqrt(dx*dx + dy*dy)
}

func connect(robots []*Robot, radius float64) {
	for _, r := range robots {
		r.inbox = make(chan Point, len(robots))
		r.out = r.out[:0]
	}
	for i := range robots {
		for j := i + 1; j < len(robots); j++ {
			if dist(robots[i], robots[j]) < radius {
				robots[i].out = append(robots[i].out, robots[j].inbox)
				robots[j].out = append(robots[j].out, robots[i].inbox)
			}
		}
	}
}

func converged(robots []*Robot) bool {
	for _, r := range robots {
		if dist(r, robots[0]) >= EPS {
			return false
		}
	}
	return true
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var n int
	fmt.Fscan(f, &n)
	robots := make([]*Robot, n)
	for i := range robots {
		robots[i] = &Robot{}
		fmt.Fscan(f, &robots[i].pos[0], &robots[i].pos[1])
	}
	var radius float64
	fmt.Fscan(f, &radius)

	rounds := 0
	for !converged(robots) && rounds < LIM_ROUNDS {
		connect(robots, radius)
		var wg sync.WaitGroup
		for _, r := range robots {
			wg.Go(r.step)
		}
		wg.Wait()
		rounds++
	}

	if converged(robots) {
		fmt.Printf("s-a ajuns la convergenta dupa %d runde\n", rounds)
	} else {
		fmt.Printf("nu s-a ajuns la convergenta dupa %d runde\n", LIM_ROUNDS)
	}
}
