package fibonacci

// Serie func returns a fibonacci serie of size n
func Serie(n int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{1}
	}
	const first = 0
	const second = 1
	serie := make([]int, n+1)
	serie[0] = first
	serie[1] = second
	for i := 2; i < len(serie); i++ {
		serie[i] = serie[i-1] + serie[i-2]
	}
	return serie[1:]
}
