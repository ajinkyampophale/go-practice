package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Create("./myfile.txt")
	checkNilError(err)

	content := "Writing content to the file."

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is", length)

	defer file.Close()

	readFile("./myfile.txt")

}

func readFile(fileName string){
	databytes, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Content: ", string(databytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}