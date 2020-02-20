package utils

import "fmt"

type Fraction struct {
	n, d int
}

func (f *Fraction) Unpack() (int, int) {
	return f.n, f.d
}

func (a *Fraction) Multiply(b Fraction) {
	a.n *= b.n
	a.d *= b.d
}

func (a *Fraction) Add(b Fraction) {
	lcm := LCM(a.d, b.d)
	a.Multiply(Fraction{b.d / lcm, b.d / lcm})
	b.Multiply(Fraction{a.d / lcm, a.d / lcm})

	a.n += b.n
	a.d += b.d
}

func (f *Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.n, f.d)
}

func ReduceFraction(n, d int) (int, int) {
	i := 2
	for i <= n && i <= d {
		for n%i == 0 && d%i == 0 {
			n /= i
			d /= i
		}
		i++
	}
	return n, d
}
