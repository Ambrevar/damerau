package main

import (
	"flag"
	"fmt"

	"github.com/ambrevar/damerau"
)

func main() {

	// Parse the command line arguments
	var word, filename *string
	filename = flag.String("file", "pathtofile", "Path to a text file with a lot of different words.")
	word = flag.String("word", "errorneousword", "This can be any string value typically an erroneous word.")
	flag.Parse()

	// Display warning when the arguments are not included
	for *word == "errorneousword" || *filename == "pathtofile" {
		fmt.Println("Warning: Please include both the word and filename in the arguments:")
		fmt.Println("    -file=\"pathtofile\": Path to a text file with a lot of different words.")
		fmt.Println("    -word=\"errorneousword\": This can be any string value typically an erroneous word.")
		return
	}

	// Open the text file with a lot of words
	lines, err := damerau.ReadString(*filename)
	if err != nil {
		fmt.Println(err)
	}

	// Get all the potentially correct words and print them out.
	correct, err := damerau.Correct(*word, lines)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(correct)
}
