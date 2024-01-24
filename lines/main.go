package main

import (
	"bytes"
	"flag"
	"io"
	"os"
)

/*
	TODO
---- Ты берешь либу flag и настаиваешь конфиг. Конфиг это структура у которой поля файл1 и файл2 и с помощью
 хуйни типа в кавычках надо будет поставить теги
----Когда в мейне будешь функцию конфига вызывать, надо будет вызывать эту хуйню чтобы она считывала флаги и клала их
в эту структуру
----Дальше ты должен написать функцию, которая принимает имя файла и отдает весь текст или в идеале какой-то
построчный буфер, условно хуйню, которая выдает построчно
----Когда ты написал эту функцию, в мейне вызываешь ее для первого и второго файла
----Типа две переменные, и потом их в цикле поочереди вызываешь
----У цикла должны быть условия для выхода, когда файлы закончатся значит функция принта должна предполагать,
что когда она встретит символ EOF (конец файла), то надо выкинуть ошибку или !ok и в условиях цикла проверять
на эту ошибку или !ок
*/

/*Напишите программу, которая принимает на вход 2 имени файла и выводит их через строчку
(строчку из первого, строчку из второго, строчку из первого и т.д.).
Если один из файлов кончился, то выводится оставшийся подряд
*/


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