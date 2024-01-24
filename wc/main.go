package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	output := ""

	linesPtr := flag.Bool("l", false, "measure the number of lines in the file.")
	wordsPtr := flag.Bool("w", false, "measure the number of words in the file.")
	bytesPtr := flag.Bool("m", false, "measure the number of bytes in the file.")
	charPtr := flag.Bool("c", false, "measure the number of characters in the file.")

	flag.Parse()
	fileName := flag.Arg(0)

	var fileData []byte

	if fileName != "" {
		file, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		fileData = file
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		fileData = stdin
	}

	switch {
	case *charPtr:
		output += fmt.Sprintf("number of characters: %d ", countChars(fileData))
	case *wordsPtr:
		output += fmt.Sprintf("number of words: %d ", countWords(fileData))
	case *linesPtr:
		output += fmt.Sprintf("number of lines: %d ", countLines(fileData))
	case *bytesPtr:
		output += fmt.Sprintf("number of bytes: %d ", countBytes(fileData))
	default:
		output = fmt.Sprintf("words  %d\nbytes  %d\nlines  %d\nchars %d\n", countWords(fileData), countBytes(fileData), countLines(fileData), countChars(fileData))
	}

	if fileName != "" {
		output += fileName
	}

	fmt.Println(output)
}
