package geometry

import "../utils"

// Point is a point
type Point struct {
	X, Y, Z int
}

func (p Point) in(b Box) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < b.W && p.Y < b.H
}
func (p *Point) Add(q Vector) {
	p.X += q.X
	p.Y += q.Y
	p.Z += q.Z
}
func (p *Point) Subtract(q Vector) {
	p.X -= q.X
	p.Y -= q.Y
}
func (p Point) Plus(q Vector) (r Point) {
	r = p
	r.Add(q)
	return r
}
func (p Point) distanceTo(q Point) (v Vector) {
	return Vector{X: utils.Abs(p.X - q.X), Y: utils.Abs(p.Y - q.Y), Z: utils.Abs(p.Z - q.Z)}
}
func (p Point) DistanceToOrigin() Vector {
	return p.distanceTo(Point{0, 0, 0})
}
func (p Point) directionTo(q Point) (v Vector) {
	return Vector{X: q.X - p.X, Y: q.Y - p.Y}
}
func (p Point) Abs() Point {
	return Point{utils.Abs(p.X), utils.Abs(p.Y), utils.Abs(p.Z)}
}

// Vector has the same data as a point, but symbolizes movement/distance
type Vector Point

func (v Vector) reduced() Vector {
	x, y := utils.ReduceFraction(v.X, v.Y)
	return Vector{X: x, Y: y}
}
func (v *Vector) inverted() Vector {
	return Vector{X: v.Y, Y: v.X}
}
func (v *Vector) scaled(n int) Vector {
	return Vector{X: v.X * n, Y: v.Y * n}
}
func (v *Vector) multiplyBy(w Vector) {
	v.X *= w.X
	v.Y *= w.Y
}
func (v Vector) multipliedBy(w Vector) Vector {
	return Vector{X: v.X * w.X, Y: v.Y * w.Y}
}
func (v Vector) dividedBy(w Vector) Vector {
	return Vector{X: v.X / w.X, Y: v.Y / w.Y}
}
func (v Vector) Abs() Vector {
	return Vector{X: utils.Abs(v.X), Y: utils.Abs(v.Y), Z: utils.Abs(v.Z)}
}
func (v Vector) min() int {
	if v.X < v.Y {
		return v.X
	}
	return v.Y
}
func (v Vector) max() int {
	if v.X > v.Y {
		return v.X
	}
	return v.Y
}
func (v Vector) Sum() int {
	return v.X + v.Y + v.Z
}
