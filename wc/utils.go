package main

import "bytes"

func countLines(fileData []byte) int {
	return bytes.Count(fileData, []byte("\n"))
}

func countWords(fileData []byte) int {
	return len(bytes.Fields(fileData))
}

func countChars(fileData []byte) int {
	return len(bytes.Runes(fileData))
}

func countBytes(fileData []byte) int {
	return len(fileData)
}