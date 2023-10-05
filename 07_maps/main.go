package main

import "fmt"

func main(){
	// Creating maps
	employee := make(map[string]string);

	// Adding/updating
	employee["name"] = "Chandan";
	employee["department"] = "CSE"
	employee["city"] = "Mysore";
	fmt.Println("Employee details:", employee);

	// Removing
	delete(employee,"department");
	fmt.Println("Employee details (after removing):", employee);

	// Looping over a map
	for key,value := range employee {
		fmt.Printf("For key %v, the value stored is %v\n", key, value);
	}

}