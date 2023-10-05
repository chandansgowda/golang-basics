package main

import "fmt"

func main(){
	names := [5]string{"Chandan", "Darshan", "Narendra"}

	// Type 1
	for i:=0; i<5; i++ {
		fmt.Println(names[i])
	}

	// Type 2
	for i := range names{
		fmt.Println(names[i])
	}

	// Type 3
	for i, name := range names{
		fmt.Printf("Index %v, Value %v\n", i, name);
	}

	// Type 4 - like while loop
	var x int8 = 0;
	for x < 5 {
		fmt.Print(x);
		x++
	}
	fmt.Println()
}