package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://github.com/AOSSIE-Org/Resonate/issues?q=is%3Aopen+is%3Aissue+label%3Aaip-2023"

func main(){

	//Parsing URL
	result, _ := url.Parse(URL);

	fmt.Println(result.Scheme);
	fmt.Println(result.Host);
	fmt.Println(result.Path);
	fmt.Println(result.RawQuery);

	// Getting query parameters in a map
	queryParams := result.Query();
	fmt.Println(queryParams);

	// URL Construction
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "github.com",
		Path: "/chandansgowda",
	}
	fmt.Println(partsOfUrl.String());

}