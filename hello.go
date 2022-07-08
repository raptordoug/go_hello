package main

import (
	"bufio"
	"fmt"
	"os"
)

var path = "c:/data/rv_request_repairs.csvb"
var myStruct struct {
	number  int
	keyTerm string
	mydesc  string
	myacts  string
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {
	fmt.Println("Hello, you kicked this off")
	var sliceIn = []string{}
	path := os.Args[1]

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		sliceIn = append(sliceIn, fileScanner.Text())
	}

	defer file.Close()
	fmt.Println(len(sliceIn))
	PrintSlice(sliceIn)
}

func PrintSlice(tmp []string) {
	for _, v := range tmp {
		fmt.Println((v))
	}
	//fmt.Println(strings.Trim(fmt.Sprint(sliceStrings), "[]"))
}
