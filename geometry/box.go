package geometry

import (
	"../utils"
	"fmt"
)

// Box is 2d square, starting at {0,0} to {W-1,H-1}
type Box struct {
	W, H int
}

// Iterate left to right, top to bottom
func (b Box) Iterate() <-chan Point {
	ch := make(chan Point)
	go func() {
		for y := 0; y < b.H; y++ {
			for x := 0; x < b.W; x++ {
				ch <- Point{x, y, 0}
			}
		}
		close(ch)
	}()
	return ch
}

// IterateNearestFirst yields all points inside the box starting from `p` in a spiral
func (b Box) IterateNearestFirst(p Point) <-chan Point {
	ch := make(chan Point)
	vectors := [4]Vector{{1, 0, 0}, {0, 1, 0}, {-1, 0, 0}, {0, -1, 0}}
	go func() {
		maxr := utils.Max(utils.Abs(b.W-p.X), utils.Abs(b.H-p.Y)) + 1
		for r := 1; r <= maxr; r++ { // "radius"
			p = Point{r * -1, r * -1, 0}
			for _, v := range vectors { // sides
				for i := 0; i < r*2; i++ {
					p.Add(v)
					if p.in(b) {
						ch <- p
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}

// TestCast prints a map of the points returned by CastRaysUntil
func (b Box) TestCast(center Point, collides func(Point) bool) {
	board := make([][]bool, b.H)
	for i := range board {
		board[i] = make([]bool, b.W)
	}

	for p := range b.CastRaysUntil(center, collides) {
		if board[p.Y][p.X] {
			fmt.Println("Repeat:", p)
		}
		if !p.in(b) {
			fmt.Println("Out of bounds:", p)
		}
		board[p.Y][p.X] = true
	}

	for _, line := range board {
		for _, char := range line {
			if char {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func SlopeIterator(max int, reverse bool) <-chan Vector {
	ch := make(chan Vector)

	go func() {
		i := 0
		queue := make([]Vector, utils.Max(max*max, 2)) // for 1/1 < slope < 1/0
		for frac := range FareySequence(max, false) {
			p, q := Vector{frac.N, frac.D, 0}, Vector{frac.D, frac.N, 0}
			if reverse {
				p, q = q, p
			}
			ch <- p
			queue[i] = q
			i++
		}

		for i -= 2; i > 0; i-- {
			ch <- queue[i]
		}

		close(ch)
	}()

	return ch
}

// CastRaysUntil collides returns true, yielding each point to the channel
func (b Box) CastRaysUntil(center Point, collides func(Point) bool) <-chan Point {
	// Starting 'north', iterate by angle clockwise
	ch := make(chan Point)
	go func() {
		quads := [4]Vector{{1, -1, 0}, {1, 1, 0}, {-1, 1, 0}, {-1, -1, 0}}
		bounds := [4]Point{{b.W - 1, 0, 0}, {b.W - 1, b.H - 1, 0}, {0, b.H - 1, 0}, {0, 0, 0}}
		for i := 0; i < 4; i++ {
			toBound := center.directionTo(bounds[i]).Abs()
			//fmt.Println("quadrant", quads[i], ";", toBound, "from", center, "to corner", bounds[i])
			for slope := range SlopeIterator(toBound.max(), i%2 == 1) {
				slope.multiplyBy(quads[i])
				//fmt.Println(slope)
				for p := center.Plus(slope); p.X >= 0 && p.X < b.W && p.Y >= 0 && p.Y < b.H; p.Add(slope) {
					//fmt.Println(p)
					if collides(p) {
						break
					}
					ch <- p
				}
			}
		}
		close(ch)
	}()
	return ch
}

// GetRayCollisions is the inverse of CastRaysUntil; it yields only what stops each ray (other than the bounds)
func (b Box) GetRayCollisions(center Point, collides func(Point) bool) <-chan Point {
	ch := make(chan Point)
	onlyCollisions := func(p Point) bool {
		if collides(p) {
			ch <- p
			return true
		}
		return false
	}

	go func() {
		for p := range b.CastRaysUntil(center, onlyCollisions) {
			_ = p
		}
		close(ch)
	}()

	return ch
}
