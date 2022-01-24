package dp

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	p2 := 0
	p3 := 0
	p5 := 0
	for i := 1; i < n; i++ {
		v2 := dp[p2] * 2
		v3 := dp[p3] * 3
		v5 := dp[p5] * 5
		x := minMultiInt(v2, v3, v5)
		dp[i] = x

		if v2 == x {
			p2++
		}
		if v3 == x {
			p3++
		}
		if v5 == x{
			p5++
		}
	}

	return dp[n-1]
}

func minMultiInt(ints ...int) int {
	min := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] < min {
			min = ints[i]
		}
	}
	return min
}
