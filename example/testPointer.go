package main

import "fmt"

func main() {
	all := []uint8{
		1,
		3,
		5,
		7,
		9,
		11,
		12,
		15,
		16,
	};

	testPoint(&all);
	
}

func testPoint(arr *[]uint8) {
	fmt.Println(arr);
	fmt.Println(*arr);
}