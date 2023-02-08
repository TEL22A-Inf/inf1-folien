package countdown

import "fmt"

// Zählt auf der Konsole von n bis 1 herunter.
func Foo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	Foo(n - 1)
}
