package hanoi

import "fmt"

func ExampleBewegePlatte() {
	BewegePlatte("A", "B")
	BewegePlatte("A", "C")
	BewegePlatte("B", "C")

	// Output:
	// A ==> B
	// A ==> C
	// B ==> C
}

func ExampleHanoi1() {
	Hanoi1("A", "B", "C")

	// Output:
	// A ==> C
}

func ExampleHanoi2() {
	Hanoi2("A", "B", "C")

	// Output:
	// A ==> B
	// A ==> C
	// B ==> C
}

func ExampleHanoi3() {
	Hanoi3("A", "B", "C")

	// Output:
	// A ==> C
	// A ==> B
	// C ==> B
	// A ==> C
	// B ==> A
	// B ==> C
	// A ==> C
}

func ExampleHanoi4() {
	Hanoi4("A", "B", "C")

	// Output:
	// A ==> B
	// A ==> C
	// B ==> C
	// A ==> B
	// C ==> A
	// C ==> B
	// A ==> B
	// A ==> C
	// B ==> C
	// B ==> A
	// C ==> A
	// B ==> C
	// A ==> B
	// A ==> C
	// B ==> C
}

func ExampleHanoi5() {
	Hanoi5("A", "B", "C")

	// Output:
	// A ==> C
	// A ==> B
	// C ==> B
	// A ==> C
	// B ==> A
	// B ==> C
	// A ==> C
	// A ==> B
	// C ==> B
	// C ==> A
	// B ==> A
	// C ==> B
	// A ==> C
	// A ==> B
	// C ==> B
	// A ==> C
	// B ==> A
	// B ==> C
	// A ==> C
	// B ==> A
	// C ==> B
	// C ==> A
	// B ==> A
	// B ==> C
	// A ==> C
	// A ==> B
	// C ==> B
	// A ==> C
	// B ==> A
	// B ==> C
	// A ==> C
}

func ExampleHanoi6() {
	Hanoi6("A", "B", "C")

	// Output:
	// A ==> B
	// A ==> C
	// B ==> C
	// A ==> B
	// C ==> A
	// C ==> B
	// A ==> B
	// A ==> C
	// B ==> C
	// B ==> A
	// C ==> A
	// B ==> C
	// A ==> B
	// A ==> C
	// B ==> C
	// A ==> B
	// C ==> A
	// C ==> B
	// A ==> B
	// C ==> A
	// B ==> C
	// B ==> A
	// C ==> A
	// C ==> B
	// A ==> B
	// A ==> C
	// B ==> C
	// A ==> B
	// C ==> A
	// C ==> B
	// A ==> B
	// A ==> C
	// B ==> C
	// B ==> A
	// C ==> A
	// B ==> C
	// A ==> B
	// A ==> C
	// B ==> C
	// B ==> A
	// C ==> A
	// C ==> B
	// A ==> B
	// C ==> A
	// B ==> C
	// B ==> A
	// C ==> A
	// B ==> C
	// A ==> B
	// A ==> C
	// B ==> C
	// A ==> B
	// C ==> A
	// C ==> B
	// A ==> B
	// A ==> C
	// B ==> C
	// B ==> A
	// C ==> A
	// B ==> C
	// A ==> B
	// A ==> C
	// B ==> C
}

func ExampleHanoi() {
	Hanoi(1, "A", "B", "C")
	fmt.Println()
	Hanoi(2, "A", "B", "C")
	fmt.Println()
	Hanoi(3, "A", "B", "C")
	fmt.Println()
	Hanoi(4, "A", "B", "C")

	// Output:
	// A ==> C
	//
	// A ==> B
	// A ==> C
	// B ==> C
	//
	// A ==> C
	// A ==> B
	// C ==> B
	// A ==> C
	// B ==> A
	// B ==> C
	// A ==> C
	//
	// A ==> B
	// A ==> C
	// B ==> C
	// A ==> B
	// C ==> A
	// C ==> B
	// A ==> B
	// A ==> C
	// B ==> C
	// B ==> A
	// C ==> A
	// B ==> C
	// A ==> B
	// A ==> C
	// B ==> C
}
