package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var temperature float64
	reader := bufio.NewReader(os.Stdin);

	fmt.Print("Enter Temperature: ")
	input, _ := reader.ReadString('\n');
	temperature, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err==nil {
		if temperature > 30{
			fmt.Println("Its Hot");
		} else if temperature < 20 {
			fmt.Println("Its Cold");
		} else{
			fmt.Println("It's normal");
		}
	} else{
		fmt.Println("Error: Try again with a float input!")
	}

}