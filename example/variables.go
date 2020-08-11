package main

import "fmt"

func main() {
	// Declare1 => var <name1, name2, ...> type = <value1, value2, ...>
	var str string = "First String";
	var num1, num2 int32 = 12, 24;
	fmt.Println("String => ", str);
	fmt.Println("num1, num2 => ", num1, num2);
	// -------------------------------------------------

	// Declare2 => var name = value => Determine `Type` According to `Value` 
	var test = "test";
	// test = 7;  ===> Error: because test is String
	fmt.Println("test => ", test);

	// not init => 0
	var i int;
	fmt.Println("i => ", i);
	// not init => false
	var b bool;
	fmt.Println("b => ", b);
	// -------------------------------------------------

	// Declare3 => name := type
	// := 左邊一定是要新宣告的
	cccc:="cccc";
	fmt.Println(cccc);
	// -------------------------------------------------
	fmt.Println("-------------------------------------------------");
	// =================================================
	
	// 多變數宣告
	// 1.
	var vn1, vn2, vn3, vn4 int;
	fmt.Printf("vn1, vn2, vn3, vn4 => %v %v %v %v\n", vn1, vn2, vn3, vn4);

	// 2. 舊的
	vn1, vn2 = 2, 3
	fmt.Printf("vn1, vn2 => %v %v\n", vn1, vn2);

	// 3. 
	var vnn1, vnn2 = 3, 4
	fmt.Printf("vnn1, vnn2 => %v %v\n", vnn1, vnn2);

	// 4. :=
	vvn1, vvn2 := 7, ":=";
	fmt.Printf("vvn1, vvn2 => %v %q\n", vvn1, vvn2);

	// 5. 因式分解式
	var (
		gv1 int = 1;
		gv2 string = "gv2";
	);
	fmt.Printf("gv1, gv2 => %v %q\n", gv1, gv2);
}