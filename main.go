package main

import (
	"fmt"

	"github.com/joanitolopo/goml/datasets"
	selection "github.com/joanitolopo/goml/model_selection"
)

func main() {
	// get dataset
	data := datasets.Load("diabetes.csv")
	table, df := data.ReadCSV(5, 50)
	fmt.Println(table)
	fmt.Println(df)

	// split dataset
	split_data := selection.Split(df, 20, 42, false)
	fmt.Println(split_data.Y_train("label"))
}
