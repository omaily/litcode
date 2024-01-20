package ozon

import "fmt"

func Task_1() error {

	var qty int
	fmt.Scan(&qty)
	outputs := make([]string, qty)

	for i := 0; i < qty; i++ {
		var in [10]int
		fmt.Scanln(&in[0], &in[1], &in[2], &in[3], &in[4], &in[5], &in[6], &in[7], &in[8], &in[9])

		ships := map[int]int{
			1: 4,
			2: 3,
			3: 2,
			4: 1,
		}

		outputs[i] = "YES"

		for _, value := range in {
			ships[value]--
			if ships[value] < 0 {
				outputs[i] = "NO"
			}
		}
	}

	for i := range outputs {
		fmt.Println(outputs[i])
	}

	return nil
}
