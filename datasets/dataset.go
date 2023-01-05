package datasets

import (
	"fmt"
	"log"

	"github.com/frictionlessdata/tableschema-go/csv"
	"github.com/frictionlessdata/tableschema-go/schema"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Method interface {
	ReadCSV(limit int) string
}

type MethodConfig struct {
	filepath string
}

func Config(filepath string) Method {
	return MethodConfig{
		filepath: filepath,
	}
}

func (m MethodConfig) ReadCSV(limit int) string {

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

	tabel, _ := m.interactive_table(tab.Headers(), columns, limit)

	return tabel
}

func (m MethodConfig) interactive_table(headers []string, columns [][]string, limit int) (string, []interface{}) {
	t := table.NewWriter()
	// t.SetOutputMirror(os.Stdout)

	//converting header []string to a []interface{}
	head_interface := make([]interface{}, len(headers))
	for i, v := range headers {
		head_interface[i] = v
	}
	t.AppendHeader(head_interface)

	var operand []interface{}
	for index, row := range columns {
		column_ith := make([]interface{}, len(row))
		for idx, value := range row {
			column_ith[idx] = value
		}
		operand = append(operand, column_ith)
		t.AppendRow(column_ith)
		if index == limit {
			break
		}

	}

	// t.AppendRow(table)
	t.AppendSeparator()

	return t.Render(), operand

}
