package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var path = "c:/data/rv_request_repairs.csv"

type myStruct struct {
	RowNum  int
	keyTerm string
	mydesc  string
	myacts  string
}

func main() {
	fmt.Println("Hello, you kicked this off")
	//var sliceIn = []string{}
	//var records [][]myStruct
	//var tstring [][]string
	//FileIn := os.Args[1]
	//FileIn := path
	//file, err := os.Open(filename)
	//records, err := ReadCsv("c:/data/rv_request_repairs.csv")
	records, err := ReadCsv()
	if isError(err) {
		panic(err)
	}

	for _, record := range records {
		rowInt, _ := strconv.Atoi(record[0])
		//fmt.Println((record))
		data := myStruct{
			RowNum:  rowInt,
			keyTerm: record[0],
			mydesc:  record[1],
			myacts:  record[2],
		}
		//	records = append(records, record)
		fmt.Printf("Row %d, %s, %s, %s\n", data.RowNum, data.keyTerm, data.mydesc, data.myacts)
	}
}

//fmt.Println(data.keyTerm)

//	reader := csv.NewReader(file)
//	records, _ := reader.ReadAll()

// fmt.Println(len(sliceIn)) //this is for the read files
// PrintSlice(sliceIn)  //this is to read the files
//PrintSlice(records)

//func ReadCsv(filename string) ([][]string, error) {
//func ReadCsv(filename string) ([][]string, error) {
func ReadCsv() ([][]string, error) {

	file, err := os.Open("c:/data/rv_request_repairs.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	fmt.Println(records)
	return records, nil
}

func ReadFullFile(filename string) ([]string, error) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return nil, err
	}
	//This part was to read  in a file line by line
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	//	for fileScanner.Scan() {
	// fmt.Println(fileScanner.Text())
	// tmp := strings.Split((fileScanner.Text())",")
	//		sliceIn = append(sliceIn, fileScanner.Text())
	//	}
	// end of this is part to read in a file

	// section 2 use a csv reader
	return nil, err
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func PrintSlice(tmp []string) {
	for _, v := range tmp {
		fmt.Println((v))
	}
	//fmt.Println(strings.Trim(fmt.Sprint(sliceStrings), "[]"))
}
