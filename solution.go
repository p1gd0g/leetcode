package main

func numSubarraysWithSum(nums []int, goal int) int {

	l := len(nums)
	inds := make([]int, 0, l)

	for i, v := range nums {
		if v == 1 {
			inds = append(inds, i)
		}
	}

	res := 0

	il := len(inds)
	if goal == 0 {
		if il == 0 {
			return add(l)
		}

		res += add(inds[0])

		for i := 0; i < il; i++ {

			if i == il-1 {
				res += add(l - inds[i] - 1)
				continue
			}

			if inds[i+1] == inds[i]+1 {
				continue
			}

			res += add(inds[i+1] - inds[i] - 1)

		}

		return res
	}

	for tail := goal - 1; tail < il; tail++ {
		head := tail - (goal - 1)

		pre0 := func() int {
			n := 0
			ind := inds[head]
			if head == 0 {
				n = ind + 1
			} else {
				n = ind - inds[head-1]
			}

			return n
		}()

		suf0 := func() int {
			n := 0
			ind := inds[tail]
			if tail == il-1 {
				n = l - ind
			} else {
				n = inds[tail+1] - ind
			}
			return n
		}()

		res += pre0 * suf0
	}
	return res

}
