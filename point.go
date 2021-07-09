package main

import "strconv"

type point struct {
	x, y int
}

func (p *point) belong(a, b *point) bool {
	return (a.x-p.x)*(b.y-p.y) == (b.x-p.x)*(a.y-p.y)
}

func (p *point) manhattan(p2 *point) int {
	return abs(p.x-p2.x) + abs(p.y-p2.y)
}

type line struct {
	x1, x2 int
	v      int
}

func (l *line) intersected(l2 *line) bool {
	return l.x1 <= l2.x2 && l2.x1 <= l.x2
}

func (l *line) String() string {
	// return strconv.Itoa(l.v)
	return strconv.Itoa(l.x1) + " " + strconv.Itoa(l.x2)
}
