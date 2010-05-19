package primes

// Copyright 2009 Anh Hai Trinh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This sieve is Eratosthenesque and only considers odd candidates.

// Print all primes <= n, where n := flag.Arg(0).
// If the flag -n is given, it will print the nth prime only.

import (
	"container/heap"
	"container/ring"
	"container/vector"
	"flag"
	"math"
)

var nth = flag.Bool("n", false, "print the nth prime only")
var nCPU = flag.Int("ncpu", 1, "number of CPUs to use")

// Return a chan of odd numbers, starting from 5.
func odds() chan int {
	out := make(chan int, 500)
	go func() {
		n := 5
		for {
			out <- n
			n += 2
		}
	}()
	return out
}

// Return a chan of odd multiples of the prime number p, starting from p*p.
func multiples(p int) chan int {
	out := make(chan int, 50)
	go func() {
		n := p * p
		for {
			out <- n
			n += 2 * p
		}
	}()
	return out
}

type PeekCh struct {
	head int
	ch   chan int
}

// Heap of PeekCh, sorting by head values.
type PeekChHeap struct {
	*vector.Vector
}

func (h *PeekChHeap) Less(i, j int) bool {
	return h.At(i).(*PeekCh).head < h.At(j).(*PeekCh).head
}

// Return a channel which serves as a sending proxy to `out`.
// Use a goroutine to receive values from `out` and store them
// in an expanding buffer, so that sending to `out` never blocks.
// See this discussion:
// <http://rogpeppe.wordpress.com/2010/02/10/unlimited-buffering-with-low-overhead>
func sendproxy(out chan<- int) chan<- int {
	proxy := make(chan int, 100)
	go func() {
		n := 128 // the allocated size of the circular queue
		first := ring.New(n)
		last := first
		var c chan<- int
		var e int
		for {
			c = out
			if first == last {
				// buffer empty: disable output
				c = nil
			} else {
				e = first.Value.(int)
			}
			select {
			case e = <-proxy:
				last.Value = e
				if last.Next() == first {
					// buffer full: expand it
					last.Link(ring.New(n))
					n *= 2
				}
				last = last.Next()
			case c <- e:
				first = first.Next()
			}
		}
	}()
	return proxy
}

// Return a chan int of primes.
func FastSieve() chan int {
	// The output values.
	out := make(chan int, 100)
	out <- 2
	out <- 3

	// The channel of all composites to be eliminated in increasing order.
	composites := make(chan int, 500)

	// The feedback loop.
	primes := make(chan int, 100)
	primes <- 3

	// Merge channels of multiples of `primes` into `composites`.
	go func() {
		h := &PeekChHeap{new(vector.Vector)}
		min := 15
		for {
			m := multiples(<-primes)
			head := <-m
			for min < head {
				composites <- min
				minchan := heap.Pop(h).(*PeekCh)
				min = minchan.head
				minchan.head = <-minchan.ch
				heap.Push(h, minchan)
			}
			for min == head {
				minchan := heap.Pop(h).(*PeekCh)
				min = minchan.head
				minchan.head = <-minchan.ch
				heap.Push(h, minchan)
			}
			composites <- head
			heap.Push(h, &PeekCh{<-m, m})
		}
	}()

	// Sieve out `composites` from `candidates`.
	go func() {
		// In order to generate the nth prime we only need multiples of
		// primes ≤ sqrt(nth prime).  Thus, the merging goroutine will
		// receive from this channel much slower than this goroutine
		// will send to it, making the buffer accumulates and blocks this
		// goroutine from sending to `primes`, causing a deadlock.  The
		// solution is to use a proxy goroutine to do automatic buffering.
		primes := sendproxy(primes)

		candidates := odds()
		p := <-candidates

		for {
			c := <-composites
			for p < c {
				primes <- p
				out <- p
				p = <-candidates
			}
			if p == c {
				p = <-candidates
			}
		}
	}()

	return out
}

//Faster version of the sieve found, this is left because it can handle uint64
//example of the sieve, it's easier to read.
func generate() chan uint64 {
	ch := make(chan uint64)
	go func(){
		for i := uint64(2); ; i++ {
			ch <- i
		}
	}()
	return ch
}
// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in chan uint64, prime uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		for {
			if i := <-in; i % prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

// The prime sieve: Daisy-chain Filter processes together.
func Sieve() <-chan uint64{
	out := make(chan uint64)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			out <- prime
			ch = filter(ch, prime)
		}
	}()
	return out
}

func bruteSieve() <-chan uint64{
	out := make(chan uint64,100)
	go func() {
		for i:=uint64(2);;i++{
			prime:=true
			for j:=uint64(2);j<uint64(math.Sqrt(float64(i)));j++{
				if i%j==0{
					prime=false
					break
				}
			}
			if prime{
				out <- i
			}
		}
	}()
	return out
}

func Factors(n uint64) []uint64{
	factors := make ([]uint64,0,100)//Hope 100 is enough
	sieve:=Sieve()
	for i:=<-sieve;;i=<-sieve{
		if n==1{
			break
		}
		for ;n%i==0; {
			n=n/i
			factors = factors[0:len(factors)+1]
			factors[len(factors)-1]=i
		}
	}
	return factors
}

func NumDivisors(n uint64) uint64{
	//Buggy implementation, if n is a perfect square then the number of divisors will necessary be one too much
	divisors:=uint64(1)
	for i:=uint64(2);i<=uint64(math.Sqrt(float64(n)));i++{
		if n%i==0{
			divisors++
		}
	}
	divisors*=2
	return divisors

}

