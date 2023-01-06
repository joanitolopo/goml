package datasets

import (
	"fmt"
	"log"
	"strconv"

	"github.com/frictionlessdata/tableschema-go/csv"
	"github.com/frictionlessdata/tableschema-go/schema"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Method interface {
	ReadCSV(row_length int, use_data ...int) (string, []map[string]interface{})
}

type MethodConfig struct {
	filepath string
}

func Config(filepath string) Method {
	return MethodConfig{
		filepath: filepath,
	}
}

func (m MethodConfig) ReadCSV(row_length int, use_data ...int) (string, []map[string]interface{}) {

	// read dataset
	tab, err := csv.NewTable(csv.FromFile(m.filepath), csv.LoadHeaders())
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+m.filepath, err)
	}

	// read the data type of each columns
	sch, err := schema.Infer(tab)
	if err != nil {
		log.Fatal("Unable to detect data type for "+m.filepath, err)
	}

	field, index := sch.GetField(tab.Headers()[0])
	fmt.Printf("%+v \n", field)
	fmt.Printf("%+v \n", index)

	// show data in interval number tells the total rows
	columns, err := tab.ReadAll()
	if err != nil {
		log.Fatal("Unable to read the data for "+m.filepath, err)
	}

	// save value with limit, if limit == 0, use all
	if use_data == nil {
		use_data = make([]int, 1)
		use_data[0] = len(columns)
	}

	tabel, value := interactive_table(tab.Headers(), columns, row_length, use_data[0])

	return tabel, value
}

func interactive_table(headers []string, columns [][]string, row_length int, use_data int) (string, []map[string]interface{}) {
	t := table.NewWriter()

	//converting header []string to a []interface{}
	head_interface := make([]interface{}, len(headers))
	for i, v := range headers {
		head_interface[i] = v
	}
	t.AppendHeader(head_interface)

	operand := []map[string]interface{}{}
	for index, row := range columns {
		column_ith := make([]interface{}, len(row))
		for idx, value := range row {
			column_ith[idx] = value
		}

		// operand = append(operand, column_ith)
		dict_value := map[string]interface{}{
			strconv.Itoa(index): column_ith,
		}
		operand = append(operand, dict_value)

		if index <= row_length {
			t.AppendRow(column_ith)
		}

		if index == use_data {
			break
		}

	}

	t.AppendSeparator()
	return t.Render(), operand

}

// mp3 := make(map[string]interface{})
//     for k, v := range mp1 {
//         if _, ok := mp1[k]; ok {
//             mp3[k] = v
//         }
//     }

//     for k, v := range mp2 {
//         if _, ok := mp2[k]; ok {
//             mp3[k] = v
//         }
//     }
