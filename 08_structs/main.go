package main

import "fmt"

func main(){

	chandan := User{"Chandan", true, 21, "CSE", "c@c.com"};

	fmt.Println(chandan);
	fmt.Println(chandan.Name, "is from", chandan.Department, "department");

	chandan.GetPermissionStatus();

	chandan.UpdateEmail("c@new.com");
	fmt.Println(chandan);

}

type User struct {
	Name string
	IsAdmin bool
	Age int16
	Department string
	Email string
}

func (u User) GetPermissionStatus() {
	if u.IsAdmin {
		fmt.Println("User is Admin")
	} else{
		fmt.Println("User is not an admin")
	}
}

func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail;
}