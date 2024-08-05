package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcomes to files in golang")
	content := "This need to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)

	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)

	}
	fmt.Println("Length of the content is ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")

}
func readFile(filename string) {
	datatype, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is  \n ",string(datatype))
}
