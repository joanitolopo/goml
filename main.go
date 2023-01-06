package main

import (
	"fmt"

	"github.com/joanitolopo/goml/datasets"
)

func main() {
	// get dataset
	init := datasets.Config("diabetes.csv")
	table, value := init.ReadCSV(5, 50)
	fmt.Println(table)
	fmt.Println(value)
}
