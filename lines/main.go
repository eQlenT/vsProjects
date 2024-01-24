package main

import (
	"bytes"
	"flag"
	"io"
	"os"
)



func readFile (fileName string) []byte {
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
	return fileData
}

func separator (data []byte) [][]byte {
	return bytes.Split(data, []byte("\n"))
}

type files struct {
	fileOneName string
	fileTwoName string
	fileOneData [][]byte
	fileTwoData [][]byte
	resFileData [][]byte
}

func (f *files) mixer (dataOne [][]byte, dataTwo [][]byte) {
//TODO Написать цикл, который переберёт все элементы dataOne и dataTwo
	for i:=0; i<len(dataOne); i++ {
			if i!=len(dataOne)-1 || i!= len(dataTwo) { 
			f.resFileData = append(f.resFileData, dataOne[i] , []byte("\n"), dataTwo[i], []byte("\n")) 
			} else if i == len(dataOne)-1 {
				f.resFileData = append(f.resFileData, dataOne[i], []byte("\n"), dataTwo, []byte("\n"))
			} 
		}	
}	
func main() {
	var myFiles files
	
	flag.StringVar(&myFiles.fileOneName, "file1", "", "first file name")
	flag.StringVar(&myFiles.fileOneName, "file2", "", "second file name")
	flag.Parse()
	flag.Arg(0)
	
	myFiles.fileOneData, myFiles.fileTwoData = separator(readFile(myFiles.fileOneName)), separator(readFile(myFiles.fileTwoName))

}