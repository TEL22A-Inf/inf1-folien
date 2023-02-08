package addition

import "fmt"

func ExampleAdd() {
	fmt.Println(Add(3, 4))
	fmt.Println(Add(2, 5))
	fmt.Println(Add(0, 2))
	fmt.Println(Add(2, 0))

	// Output:
	// 7
	// 7
	// 2
	// 2
}
