package fakultaet

import "fmt"

func ExampleFactorialIt() {
	fmt.Println(FactorialIt(5))
	fmt.Println(FactorialIt(4))
	fmt.Println(FactorialIt(3))
	fmt.Println(FactorialIt(2))
	fmt.Println(FactorialIt(1))

	// Output:
	// 120
	// 24
	// 6
	// 2
	// 1
}

func ExampleFactorialRec() {
	fmt.Println(FactorialRec(5))
	fmt.Println(FactorialRec(4))
	fmt.Println(FactorialRec(3))
	fmt.Println(FactorialRec(2))
	fmt.Println(FactorialRec(1))

	// Output:
	// 120
	// 24
	// 6
	// 2
	// 1
}
