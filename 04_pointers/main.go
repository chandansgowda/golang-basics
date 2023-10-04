package main

import (
	"fmt"
	"runtime"
)

func main(){
	cpuCount := runtime.NumCPU();
	fmt.Println("No. of CPUs:", cpuCount);

	var ptr *int = &cpuCount
	fmt.Println("Pointer address:", &ptr,);
	fmt.Println("\nAddress Pointed by Pointer:", ptr);
	fmt.Println("Value stored in address:", *ptr);

	*ptr *= 2 // changing the value stored in the address pointed to by the pointer variable

	fmt.Println("\nAddress Pointed by Pointer (After Change):", ptr);
	fmt.Println("Value stored in address (After Change):", *ptr);
}