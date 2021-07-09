package main

// x^y
func pow(x, y int) int {
	z := 1
	for i := 0; i < y; i++ {
		z *= x
	}
	return z
}

func qpow(x, n int) int {
	const mod = 1e9 + 7
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res

}

func greater(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func less(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func maxn(xs ...int) int {
	x := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > x {
			x = xs[i]
		}
	}
	return x
}

func min(xs ...int) int {
	x := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] < x {
			x = xs[i]
		}
	}
	return x
}

func equalSliceInt(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func absmod(x, y int) int {

	x = x % y
	if x >= 0 {
		return x
	}
	x += y
	return x
}

func search(l int, f func(int) bool) int {
	for i := 0; i < l; i++ {
		if f(i) {
			return i
		}
	}
	return l
}

func uniq(x []int) (y []int) {

	m := make(map[int]struct{}, len(x))
	y = make([]int, 0, len(x))

	for _, v := range x {
		m[v] = struct{}{}
	}

	for k := range m {
		y = append(y, k)
	}

	return
}

func appends(s []int, e ...int) []int {

	ns := make([]int, len(s), len(s)+1)
	copy(ns, s)
	ns = append(ns, e...)

	return ns
}

func permute(n, r int) int {
	ans := 1
	for i := 0; i < r; i++ {
		ans *= n
		n--
	}

	return ans
}

func combine(n, r int) int {
	return permute(n, r) / permute(r, r)
}

func add(x int) int {
	ans := 0
	for i := 1; i <= x; i++ {
		ans += i
	}
	return ans
}
