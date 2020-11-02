package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const pageSize int64 = 4096
const lineSize int64 = 256

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	files, err := ioutil.ReadDir("lib")
	check(err)

	fmt.Println("Давайте погадаем!")

	fileNames := make(map[int]string)

	for key, file := range files {
		bookName := file.Name()
		fileNames[key+1] = bookName
		fmt.Println(key+1, bookName)
	}

	fmt.Println("Выберите книгу: ")
	var bookNumber float64
	fmt.Scanf("%f", &bookNumber)

	filePath := "lib/" + fileNames[int(bookNumber)]

	fileStat, err := os.Stat(filePath)
	check(err)

	size := fileStat.Size()
	fmt.Println(size)

	fmt.Print("Введите номер страницы от 1 до ", size/pageSize-1, ": ")
	var pageNumber float64
	fmt.Scanf("%f", &pageNumber)

	fmt.Print("Введите номер строки от 1 до ", pageSize/lineSize, ": ")
	var lineNumber float64
	fmt.Scanf("%f", &lineNumber)

	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	_, err = f.Seek(pageSize*int64(pageNumber)+lineSize*int64(lineNumber), 0)
	check(err)

	chunk := make([]byte, lineSize)
	text, err := f.Read(chunk)
	check(err)

	newString := strings.ToValidUTF8(string(chunk[:text]), "")
	fmt.Printf("%v\n", strings.TrimSpace(newString))
}
