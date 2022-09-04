package main

import (
	"fmt"
	"log"
	"os"
)

//File.WriteString() writes the contents of a string to a file.

func main() {
	fmt.Println("first practice with file.WriteString()")
	//os.Create() creates the named file
	file1, err := os.Create("myfirstfile.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your file has been created")
	defer file1.Close()

	//file.WriteString() writes our string to the created file.
	_, err2 := file1.WriteString("The stabilo textmarkers are the best\n")

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("You have successfully written to myfile.txt\n Check it out with cat")
	fmt.Println("-------------------------------")

}
