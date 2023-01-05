package modelselection

import "fmt"

func Train_Test_split(array [][]string, test_size float64) (x_train, y_train, x_test, y_test [][]string) {
	fmt.Println(array)

	x_train = append(x_train, []string{"x_train"})
	y_train = append(y_train, []string{"y_train"})
	x_test = append(x_test, []string{"x_test"})
	y_test = append(y_test, []string{"y_test"})

	fmt.Println(test_size)

	return x_train, x_test, y_train, y_test
}
