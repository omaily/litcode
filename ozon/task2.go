package ozon

import "fmt"

func contains(s *[]int, e int) bool {
	for _, a := range *s {
		if a == e {
			return true
		}
	}
	return false
}

func Task_2() {

	var qty int
	fmt.Scan(&qty)
	outputs := make([]string, qty)

	// var in [3]int
	var year int
	var month int
	var day int
	var februaryDays int

	shortMonth := []int{4, 6, 9, 11}
	for i := 0; i < qty; i++ {

		februaryDays = 28
		outputs[i] = "YES"

		fmt.Scanln(&day, &month, &year)

		if year%400 == 0 {
			februaryDays++
		} else if year%100 == 0 {

		} else if year%4 == 0 {
			februaryDays++
		}

		if month == 2 && day > februaryDays {
			outputs[i] = "NO"
		}
		if contains(&shortMonth, month) && day > 30 {
			outputs[i] = "NO"
		}

	}
	for i := range outputs {
		fmt.Println(outputs[i])
	}

}
