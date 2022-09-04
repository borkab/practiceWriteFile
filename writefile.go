package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//File.WriteString() writes the contents of a string to a file.
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
	fmt.Println("-------------------------------")

	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")

	fmt.Println("second practice with File.Write() and File.WriteAt()")
	//the File.Write() writes n bytes to a file.
	//the File.WriteAt() writes n bytes to a file starting at the specified byte offset.
	file2, err3 := os.Create("mysecondfile.txt")

	if err3 != nil {
		log.Fatal(err3)
	}
	defer file2.Close()

	fmt.Println("we created to our file mysecondfile.txt")
	//first we write our string to the file with File.Write()
	//then we write our second string next to it with File.WriteAt()

	val1 := "Gopher and Octocat"
	//we transform the string to bytes
	data := []byte(val1)

	//we write the bytes with Write()
	_, err4 := file2.Write(data)

	if err4 != nil {
		log.Fatal(err4)
	}
	fmt.Println("we wrote to our file mysecondfile.txt")

	val2 := " are best friends"
	data2 := []byte(val2)
	//we calculate the length of the previously written string
	var length int64 = int64(len(data))

	//we write the string and the specified index with WriteAt()
	_, err5 := file2.WriteAt(data2, length)

	if err5 != nil {
		log.Fatal(err5)
	}
	fmt.Println("we wrote our second string to our file mysecondfile.txt ")

	val3 := " forever <3 awww"
	data3 := []byte(val3)

	_, err6 := file2.WriteAt(data3, length)

	if err6 != nil {
		log.Fatal(err6)
	}
	fmt.Println("we wrote our third string to our file mysecondfile.txt\nCheck it out with cat")
	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")

	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")

	fmt.Println("third practice with ioutil.WriteFile()")

	val4 := "Adam: the turning red movie is only to buy for 13.99EUR on prime\n"
	data4 := []byte(val4)

	err7 := ioutil.WriteFile("redonprime.txt", data4, 0644)

	if err7 != nil {
		log.Fatal(err7)
	}

	fmt.Println("you have successfully written to your new txt file")

	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")

	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")

}
