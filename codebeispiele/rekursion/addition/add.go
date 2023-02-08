package addition

// Berechnet rekursiv die Summe m+n.
func Add(m, n int) int {
	if n == 0 {
		return m
	}
	return 1 + Add(m, n-1)
}
