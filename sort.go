package main

type By struct {
	a []rune
	b []int
}

func (a By) Len() int { return len(a.b) }
func (a By) Swap(i, j int) {
	a.a[a.b[i]], a.a[a.b[j]] = a.a[a.b[j]], a.a[a.b[i]]
}
func (a By) Less(i, j int) bool {
	return a.a[a.b[i]] < a.a[a.b[j]]
}
