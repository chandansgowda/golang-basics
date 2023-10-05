package main

import "fmt"

func main(){

	chandan := User{"Chandan", true, 21, "CSE"};

	fmt.Println(chandan);
	fmt.Println(chandan.Name, "is from", chandan.Department, "department");

}

type User struct {
	Name string
	IsAdmin bool
	Age int16
	Department string
}
