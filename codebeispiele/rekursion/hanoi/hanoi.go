package hanoi

import "fmt"

// Bewegen einer einzelnen Platte von start nach ziel.
func BewegePlatte(s, z string) {
	fmt.Printf("%s ==> %s\n", s, z)
}

// Bewegen eines Turms der Höhe 1 von start über mitte nach ziel.
func Hanoi1(s, m, z string) {
	BewegePlatte(s, z)
}

// Bewegen eines Turms der Höhe 2 von start über mitte nach ziel.
func Hanoi2(s, m, z string) {
	Hanoi1(s, z, m)
	BewegePlatte(s, z)
	Hanoi1(m, s, z)
}

// Bewegen eines Turms der Höhe 3 von start über mitte nach ziel.
func Hanoi3(s, m, z string) {
	Hanoi2(s, z, m)
	BewegePlatte(s, z)
	Hanoi2(m, s, z)
}

// Bewegen eines Turms der Höhe 4 von start über mitte nach ziel.
func Hanoi4(s, m, z string) {
	Hanoi3(s, z, m)
	BewegePlatte(s, z)
	Hanoi3(m, s, z)
}

// Bewegen eines Turms der Höhe 5 von start über mitte nach ziel.
func Hanoi5(s, m, z string) {
	Hanoi4(s, z, m)
	BewegePlatte(s, z)
	Hanoi4(m, s, z)
}

// Bewegen eines Turms der Höhe 6 von start über mitte nach ziel.
func Hanoi6(s, m, z string) {
	Hanoi5(s, z, m)
	BewegePlatte(s, z)
	Hanoi5(m, s, z)
}

// Allgemeine Lösung:
// Bewegen eines Turms der Höhe n von start über mitte nach ziel.
func Hanoi(n int, s, m, z string) {
	if n == 0 {
		return
	}
	Hanoi(n-1, s, z, m)
	BewegePlatte(s, z)
	Hanoi(n-1, m, s, z)
}
