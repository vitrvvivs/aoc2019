package main

import (
	"fmt"
)

const DEBUG = false

func test(x int, exact_repeat bool) bool {
	/* It is a six-digit number
	 * Two adjacent digits are the same (like 22 in 122345).
	 * Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
	 */
	if x < 100000 || x > 999999 { return false }

	j := 10
	repeat := 1
	repeat_matches := false
	for ; x > 0; x /= 10 {
		if x % 10 > j { return false }
		if x % 10 == j {
			repeat++
		} else if exact_repeat && repeat == 2 {
			repeat_matches = true
			repeat = 1
		} else {
			repeat = 1
		}
		if (! exact_repeat) && repeat == 2 {
			repeat_matches = true
		}
		j = x % 10
	}
	return repeat_matches || repeat == 2
}

func main() {
	min, max := 134792, 675810
	counter := 0

	if DEBUG {
		fmt.Println(111111, test(111111, false), true)
		fmt.Println(122349, test(122349, false), true)
		fmt.Println(223450, test(223450, false), false)
		fmt.Println(123789, test(123789, false), false)
		fmt.Println(113789, test(123789, false), true)
	}

	for i := min; i <= max; i++ {
		if test(i, false) {counter++}
	}
	fmt.Println(counter)

	if DEBUG {
		fmt.Println(111111, test(111111, true), false)
		fmt.Println(112233, test(112233, true), true)
		fmt.Println(123456, test(123456, true), false)
		fmt.Println(123444, test(123444, true), false)
		fmt.Println(111122, test(111122, true), true)
		fmt.Println(111223, test(111223, true), true)
		fmt.Println(223450, test(223450, true), false)
		fmt.Println(113456, test(223450, true), true)
	}

	counter = 0
	for i := min; i <= max; i++ {
		if test(i, true) {counter++}
	}
	fmt.Println(counter)
}
