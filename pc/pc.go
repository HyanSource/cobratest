package pc

//it's a permutations and combination package

//factorial
func F(n int) int {
	t := 1
	for i := 1; i <= n; i++ {
		t *= i
	}
	return t
}

//permutations
func P(n int, k int) int {
	return F(n) / F(n-k)
}

//combination
func C(n int, k int) int {
	return F(n) / (F(k) * F(n-k))
}

//combination with repetition
func H(n int, k int) int {
	return C(n+k-1, k)
}
