package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type dict map[string][]string

var dData dict

func findWord(s string, d dict) ([]string, error) {
	if def, ok := d[s]; ok {
		return def, nil
	} else if def, ok := d[strings.ToLower(s)]; ok {
		return def, nil
	}

	return nil, errors.New("could not find the word")
}

func read(fName string) {
	file, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatalf("Could not open file: %v\n", err)
	}

	jErr := json.Unmarshal(file, &dData)
	if jErr != nil {
		log.Fatalf("Could not unmarshal the file: %v\n", err)
	}
}

// Read the command line
func getInput() string {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a word to search for")
	}
	var words string
	for _, v := range os.Args[1:] {
		words += v + " "
	}
	words = strings.TrimSuffix(words, " ")
	return words

}

func outPut(ss string, s []string) {

	fmt.Printf("You searched for the meaning of (%s):\n", ss)
	fmt.Println("Possible meanings: ")
	for i, v := range s {
		fmt.Printf("%d: %s\n", i+1, v)
	}
}

func main() {
	fName := "data/data.json"
	read(fName)
	sWord := getInput()
	rWord, err := findWord(sWord, dData)
	if err != nil {
		log.Fatalf("%v %v ", err, sWord)
	}
	outPut(sWord, rWord)
}
