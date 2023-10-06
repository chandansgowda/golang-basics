package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin);

	fmt.Print("Enter filename: ");
	filename,_ := reader.ReadString('\n');
	filename = strings.TrimSpace(filename);

	fmt.Print("File content: ");
	content,_ := reader.ReadString('\n');
	content = strings.TrimSpace(content);

	// Creating a new file
	file, err := os.Create("./"+filename+".txt");
	handleError(err);

	// Writing to the file
	contentLength, err := io.WriteString(file, content);
	handleError(err);
	fmt.Println("File Was Created Successfully", "\nFile length:", contentLength);

	// Reading a file
	databytes,err := os.ReadFile("./"+filename+".txt");
	handleError(err);
	fmt.Println("Databyte:", databytes, "\nString:", string(databytes));

	// Deleting a file
	err = os.Remove("./"+filename+".txt");
	handleError(err);

	// Closing a file
	defer file.Close();

}

func handleError(err error){
	if err!=nil{
		panic(err);
	}
}