package main

import (
	"fmt"
	"sort"
)

func main(){

	// Def and Init

	var myNumbers = []int{3,4,2,1};
	fmt.Println("Slice:", myNumbers);

	//OR

	numbers2 := make([]int,5);
	numbers2[0] = 1
	numbers2[1] = 2
	fmt.Println("Slice using make():", numbers2, "Length:", len(numbers2));

	// Sorting
	sort.Ints(myNumbers);
	fmt.Println("Sorted Slice:", myNumbers);

	// Adding elements
	myNumbers  = append(myNumbers, 100, 200, 300);
	fmt.Println("Slice after adding 100, 200, 300:", myNumbers, "Length:", len(myNumbers), "Capacity:", cap(myNumbers));

	// Removing elements
	myNumbers = append(myNumbers[:2], myNumbers[3:]...)
	fmt.Println("Slice after removing element at index 2:", myNumbers)

	// Check if slice is sorted
	fmt.Println(sort.IntsAreSorted(myNumbers));
}