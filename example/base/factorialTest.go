package main

import "fmt"

func main() {
	fmt.Println(fac(3));
}

func fac(n int) int {
	if(n == 1 || n == 0) {
		return 1;
	} else {
		return fac(n-1) * n;
	}
}