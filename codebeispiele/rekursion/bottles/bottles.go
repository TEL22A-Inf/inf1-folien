package bottles

import "fmt"

// Schreibt den Text des Liedes "99 Bottles of Beer"
// für Startwerte >= 1 auf die Konsole.
func Bottles(n int) {
	if n == 1 {
		fmt.Println("One last bottle of beer on the wall,")
		fmt.Println("One last bottle of beer.")
		return
	}
	fmt.Printf("%d bottles of beer on the wall,\n", n)
	fmt.Printf("%d bottles of beer.\n", n)
	fmt.Println("Take one down, pass it around,")
	fmt.Printf("%d bottles of beer on the wall.\n", n-1)
	Bottles(n - 1)
}
