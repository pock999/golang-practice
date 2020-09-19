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
	fmt.Println(binarySearch(all, uint8(16), 0, len(all)));
}

func binarySearch(arr []uint8, n uint8, start int, end int) int {
	mid := (start + end) / 2;
	if (n > arr[mid]) {
		return binarySearch(arr, n, mid, end);
	} else if (n < arr[mid]) {
		return binarySearch(arr, n, start, mid);
	} else {
		return mid;
	}
}

// select sort
func sort(arr []uint8) []uint8 {
	fmt.Println("原始：", arr);

	var temp uint8;
	var i, j, length int;

	i = 0;

	length = len(arr);

	for (i < length-1) {
		j = i;
		for (j<length) {
			if (arr[i] > arr[j]) {
				temp = arr[i];
				arr[i] = arr[j];
				arr[j] = temp;
			}
			j++;
		}
		i++;
	}
	fmt.Println("afete sort：", arr);
	return arr;
}