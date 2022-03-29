package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//opening a text file
func openfile() {
	f, err := os.Create("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

}

func repetition(st string) map[string]int {

	// using strings.Field Function=it used to splits the given string around each instance of one or more consecuitive white space characters,as defind by unicode.isspace,returning a slice of subracting of str or an empty slice
	//using word count
	input := strings.Fields(st)
	wc := make(map[string]int) //here map is used for collectio of unordered pairs of key-value
	for _, word := range input {
		_, matched := wc[word]
		if matched { //using if statement
			wc[word] += 1 //if word matched
		} else { //else if word does not matched
			wc[word] = 1
		}
	}
	return wc
}

func main() {

	openfile()
	b, err := ioutil.ReadFile("temp.txt") // reading file
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	fmt.Println(str)
	for index, element := range repetition(str) {
		fmt.Println(index, "=", element)
	}

}
