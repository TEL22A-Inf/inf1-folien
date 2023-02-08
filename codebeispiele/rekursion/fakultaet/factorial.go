package fakultaet

// Iterative Berechnung der Fakultät.
func FactorialIt(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Rekursive Berechnung der Fakultät.
func FactorialRec(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FactorialRec(n-1)
}
