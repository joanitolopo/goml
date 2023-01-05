package main

import (
	"fmt"

	"github.com/joanitolopo/goml/datasets"
)

func main() {
	// get dataset
	init := datasets.Config("diabetes.csv")
	df := init.ReadCSV(15)

	fmt.Println(df)
}
