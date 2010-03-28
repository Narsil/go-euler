package primes

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

func fixedgenerate(limit uint64) chan uint64 {
    ch := make(chan uint64)
    go func(){
        for i := uint64(2);i<limit ; i++ {
            ch <- i
        }
    }()
    return ch
}

func FixedSieve(limit uint64) <-chan uint64{
	out := make(chan uint64)
    go func() {
        ch := fixedgenerate(limit)
        for {
            prime := <-ch
            out <- prime
            ch = filter(ch, prime)
        }
    }()
    return out
}

func brutePrime(number uint64) bool {
    for i := uint64(2); i < number; i++ {
        if number%i == 0 {
            return false
        }
    }
    return true;
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

func Prime(number uint64) bool{
	return brutePrime(number)
}
