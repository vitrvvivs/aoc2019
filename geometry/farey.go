package geometry

type Fraction struct {
	N, D int
}

func (f *Fraction) flipped() Fraction {
	return Fraction{f.D, f.N}
}

func FareySequence(max int, reverse bool) <-chan Fraction {
	ch := make(chan Fraction)

	step := func(frac, next *Fraction) (*Fraction, *Fraction) {
		lcd := (max + frac.D) / next.D
		return next, &Fraction{next.N*lcd - frac.N, next.D*lcd - frac.D}
	}

	go func() {
		var frac, next *Fraction
		var test func() bool

		if !reverse {
			frac, next = &Fraction{0, 1}, &Fraction{1, max}
			test = func() bool { return next.N <= max }
		} else {
			frac, next = &Fraction{1, 1}, &Fraction{max - 1, max}
			test = func() bool { return frac.N > 0 }
		}

		ch <- *frac

		for test() {
			frac, next = step(frac, next)
			ch <- *frac
		}

		close(ch)
	}()

	return ch
}
