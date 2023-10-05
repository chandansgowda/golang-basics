package main

import "fmt"

func main(){
	var (
		a int = 1
		b int = 2
	)
	fmt.Println(addTwoNumbers(a,b));
	fmt.Println(addManyNumbers(1,2,3,4,5,6,7,8,9));

	swapTwoVariables(&a,&b);
	fmt.Println("After swapping\na:", a, "b:", b);
}

func addTwoNumbers(a int, b int) int {
	return a+b;
}

func addManyNumbers(values ...int) int {
	var res int = 0;
	for _, value := range values {
		res += value;
	}
	return res;
}

func swapTwoVariables(a *int, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}