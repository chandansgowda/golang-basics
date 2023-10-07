package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL string = "https://dog.ceo/api/breeds/image/random";

func main(){
	// Get Request
	response, err := http.Get(URL);
	handleError(err);

	databytes,err := io.ReadAll(response.Body);
	handleError(err);

	fmt.Println("Response:", string(databytes));

	defer response.Body.Close();
}

func handleError(err error){
	if err!=nil{
		panic(err);
	}
}