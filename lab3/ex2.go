package main

import(
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

type LamportClock struct{ counter int }

// incrementeaza contorul
func (l *LamportClock) Tick() { l.counter++ }
func (l *LamportClock) Send() { l.Tick() }
func (l *LamportClock) Receive(ts int) { l.counter = max(l.counter, ts); l.Tick() }

type Message struct{ sender, ts int }
type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

// TODO: retuneaza true daca a precede b
func less(a, b Event) bool {
	if a.internal_counter == b.internal_counter{
		return a.process < b.process
	}
	return a.internal_counter < b.internal_counter
}

type Process struct{
	id int
	clock LamportClock
	next chan <- Message
	prev <- chan Message
	events []Event
}

func (p *Process) operate(rounds int){
	for range rounds{
		for range rand.Intn(4){
			p.clock.Tick()
			p.events = append(p.events, Event{p.id, p.clock.counter, "internal"})
		}

		p.clock.Send()
		p.next <- Message{p.id, p.clock.counter}
		p.events = append(p.events, Event{p.id, p.clock.counter, "send"})

		m := <-p.prev
		p.clock.Receive(m.ts)
		p.events = append(p.events, Event{p.id, p.clock.counter, "recv"})
	}
}

func main(){
	var procs []Process
	var chans []chan Message
	for range 3{
		chans = append(chans, make(chan Message, 5))
	}

	for i := range 3{
		procs = append(procs, Process{i + 1, LamportClock{0}, chans[(i + 1) % 3], chans[i], nil})
	}

	var wg sync.WaitGroup
	for i := range 3{
		wg.Go(func(){
			procs[i].operate(5)
		})
	}

	wg.Wait()
	var events []Event
	for _, v := range procs{
		events = append(events, v.events...)
	}

	sort.Slice(events, func(i, j int) bool{
		return less(events[i], events[j])
	})

	for _, v := range events{
		fmt.Printf("P%d %s %d\n", v.process, v.kind, v.internal_counter)
	}
}
