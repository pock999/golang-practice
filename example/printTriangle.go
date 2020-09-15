package main

import "fmt"

func main() {
	var x uint16;
	fmt.Println("請輸入一個正整數，：")
	fmt.Scanln(&x);
	fmt.Printf("x的數值：%d\n",x)
	printTri(x);
}

func printTri(x uint16) {
	var i uint16;
	var j uint16;

	for i = 0 ;i<x ; i++ {
		str := "";
		for j=0 ; j<i ; j++ {
			str = str + "*";
		}
		fmt.Println(str);
	}
}