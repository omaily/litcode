package ozon

import (
	"fmt"
	"regexp"
)

func Task_3() {

	var qty int
	fmt.Scan(&qty)
	var numbersRow string
	outputs := make([][]string, qty)

	for i := 0; i < qty; i++ {

		fmt.Scanln(&numbersRow)

		validData := true
		zp := regexp.MustCompile(`[A-Z]\d{1,2}[A-Z]{2}`)
		temp := zp.Split(numbersRow, -1)
		for i := range temp {
			if temp[i] != "" {
				validData = false
				break
			}
		}

		outputs[i] = zp.FindAllString(numbersRow, -1)

		if validData {
			outputs[i] = zp.FindAllString(numbersRow, -1)
		} else {
			outputs[i] = []string{"-"}
		}
	}

	for _, line := range outputs {
		for _, part := range line {
			fmt.Printf("%s ", part)
		}
		fmt.Println()
	}

}
