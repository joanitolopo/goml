package modelselection

import (
	"fmt"
	"math/rand"

	"github.com/Kunde21/numgo"
)

type Train_test_split interface {
	X_train() *numgo.Array64
	X_test() *numgo.Array64
	Y_train() *numgo.Array64
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

func (c data_split) Y_train() *numgo.Array64 {
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
	n_arrays := array.Shape()[0]
	n_train, n_test := func() (n_train, n_test int) {
		rati := float32(test_size) / 100
		var test_size = ((rati) * float32(n_arrays))
		n_test = int(test_size)
		n_train = n_arrays - n_test
		return n_train, n_test
	}()
	// fmt.Println(n_arrays)
	// fmt.Println(n_train)
	// fmt.Println(n_test)

	var train_idx *numgo.Array64
	var test_idx *numgo.Array64

	if !shuffle[0] {
		// not use random shuffle
		train_idx = numgo.Arange(float64(n_train))
		test_idx = numgo.Arange(float64(n_train), float64(n_train+n_test))

	} else {
		rand.Seed(int64(random_state))
		// rand.Shuffle(n_arrays)
	}

	fmt.Println(train_idx)
	fmt.Println(test_idx)

	return data_split{
		x_train: array,
		x_test:  array,
		y_train: array,
		y_test:  array,
	}

}
