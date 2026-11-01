package main

import (
	"fmt"
	"math/rand"
	"time"
)

// stocheaza deviatia fata de timpul real
type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

// TODO:
// Calculeaza decalajul theta si round-trip-ul delta
func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	delta = t2.Sub(t1) + t4.Sub(t3);
	theta = t2.Sub(t1) + t3.Sub(t4); theta /= 2

	return
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	// Request
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	t2 := s_clock.Now()
	// Procesare intre 5 si 20 ms
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()

	// Response
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	c_clock := Clock{offset: -time.Duration(150) * time.Millisecond}
	s_clock := Clock{offset: 0}

	theta, delta := measure(&c_clock, &s_clock)

	thetamin, deltamin := theta, delta
	for range 10 {
		theta, delta = measure(&c_clock, &s_clock)
		thetamin = min(thetamin, theta)
		deltamin = min(deltamin, delta)
	}

	fmt.Println(c_clock.offset + theta, c_clock.offset + thetamin)
}
