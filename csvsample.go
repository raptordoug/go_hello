package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	ID  string
	Age int
}

func main() {
	records, err := ReadCsv("example.csv")
	if err != nil {
		panic(err)
	}

	var employee []Employee
	for _, record := range records {
		ageInt, _ := strconv.Atoi(record[1])
		data := Employee{
			ID:  record[0],
			Age: ageInt,
		}
		employee = append(employee, data)
	}
	fmt.Println(employee)
}

func ReadCsv(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
