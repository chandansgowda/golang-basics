package main

import "fmt"

func main(){
	var students [5]string
	students[0] = "Chandan"
	students[1] = "Darshan"
	students[4] = "Narendra"
	fmt.Println("\nStudent Names are:", students);
	// Will print 5 even though we have added 3 students only - because we have declared the array with size 5
	fmt.Println("Total Students:", len(students));

	var teachers = [3]string{"John", "Mary", "Ravan"}
	fmt.Println("Teacher Names:", teachers);
	fmt.Println("Total Teachers:", len(teachers));
}