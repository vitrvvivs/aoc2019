package utils

// Maximum
func Max(ar ...int) int {
	r := 0
	for _, v := range ar {
		if v > r {
			r = v
		}
	}
	return r
}

// Minimum
func Min(ar ...int) int {
	r := 0x7fffffffffffffff
	for _, v := range ar {
		if v < r {
			r = v
		}
	}
	return r
}

// Absolute value
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Ceiling
func Cieldiv(n, d int) int {
	r := n / d
	if n%d != 0 {
		r++
	}
	return r
}

func GCD(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

func Limit(i int, n int) int {
	if i <= -n {
		i = -n
	} else if i >= n {
		i = n
	}
	return i
}
