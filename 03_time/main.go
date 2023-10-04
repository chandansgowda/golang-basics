package main

import (
	"fmt"
	"time"
)

func main(){
	presentTime := time.Now();

	fmt.Println("Present Time:", presentTime);

	// We have to use the following specific date and time for formatting any date
	formattedPresentTime := presentTime.Format("01-02-2006 15:04:05 Monday");

	fmt.Println("Formatted Present Time:", formattedPresentTime);

	createdTime := time.Date(2001,time.January, 1, 11, 11, 11, 0, time.Local);

	fmt.Println("Created Time:", createdTime);
}