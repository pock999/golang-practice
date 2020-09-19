package main

import "fmt"

func main() {
	all := []uint8{
		15,
		11,
		12,
		16,
		1,
		3,
		9,
		5,
		7,
	};

	all = sort(all);
	fmt.Println(all);
	
}

// select sort
func sort(arr []uint8) []uint8 {
	fmt.Println("原始：", arr);

	var i, j, length int;

	i = 0;

	length = len(arr);

	for (i < length-1) {
		j = i;
		for (j<length) {
			if (arr[i] > arr[j]) {
				swap(&arr[i], &arr[j]);
			}
			j++;
		}
		i++;
	}
	fmt.Println("afete sort：", arr);
	return arr;
	
}

func swap(x, y *uint8) {
	temp := *x;
	*x = *y;
	*y = temp;
}