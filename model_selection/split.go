package modelselection

import (
	"math/rand"

	"github.com/Kunde21/numgo"
)

type Train_test_split interface {
	X_train() *numgo.Array64
	X_test() *numgo.Array64
	Y_train(name string) *numgo.Array64
	Y_test() *numgo.Array64
}

type data_split struct {
	x_train *numgo.Array64
	x_test  *numgo.Array64
	y_train *numgo.Array64
	y_test  *numgo.Array64
}

func (c data_split) X_train() *numgo.Array64 {
	return c.x_train
}

func (c data_split) Y_train(name string) *numgo.Array64 {
	return c.y_train
}

func (c data_split) X_test() *numgo.Array64 {
	return c.x_test
}

func (c data_split) Y_test() *numgo.Array64 {
	return c.y_test
}

func Split(array *numgo.Array64, test_size int, random_state int, shuffle ...bool) Train_test_split {

	if shuffle == nil {
		shuffle = append(shuffle, false)
	}

	// split data based on sklearn process
	n_columns := array.Shape()[1]
	n_rows := array.Shape()[0]

	n_train, n_test := validation(test_size, n_rows)

	var train_idx []int
	var test_idx []int
	var train_data *numgo.Array64
	var test_data *numgo.Array64

	if !shuffle[0] {
		// not use random shuffle
		train_idx = Arange(n_train)
		test_idx = Arange(n_train, n_train+n_test)

		// get data
		train_data = get_data(array, train_idx, n_train, n_columns)
		test_data = get_data(array, test_idx, n_test, n_columns)

	} else {
		rand.Seed(int64(random_state))
		// rand.Shuffle(n_arrays)
	}

	return data_split{
		x_train: train_data,
		x_test:  test_data,
		y_train: array,
		y_test:  array,
	}

}

func get_data(array *numgo.Array64, l_index []int, n_rows, n_columns int) *numgo.Array64 {
	result := []float64{}
	for _, idx := range l_index {
		row := array.SliceElement(idx)
		result = append(result, row...)
	}

	data := numgo.NewArray64(result, n_rows, n_columns)

	return data

}

func validation(test_size, n_arrays int) (n_train, n_test int) {
	rati := float32(test_size) / 100
	var test_size2 = ((rati) * float32(n_arrays))
	n_test = int(test_size2)
	n_train = n_arrays - n_test
	return n_train, n_test
}
