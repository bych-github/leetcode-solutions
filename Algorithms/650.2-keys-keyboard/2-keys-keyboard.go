func minSteps(n int) int {
    res := 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			res += i
			n /= i
		}
	}
	return res
}