package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin);

	var name string;
	fmt.Println("Hello User, Please Enter Your name: ");
	name, _ = reader.ReadString('\n');
	name = strings.TrimSpace(name);

	fmt.Println("Hello",name, "Please Enter Your age: ");
	age, _ := reader.ReadString('\n');

	ageWithoutSpaceOrNewline := strings.TrimSpace(age);

	ageAsInt, err := strconv.ParseInt(ageWithoutSpaceOrNewline, 10, 32);

	if (err!=nil){
		fmt.Println("Error Ocuured")
		return;
	} else{
		fmt.Println("Your name is", name);
		fmt.Println("Your age is", ageAsInt);
		fmt.Println("Your age after 10 years is:", ageAsInt+10);
	}

}