package ozon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func hopeless(scanner *bufio.Reader, numerser []int) {
	numerser[0] = -1
	for i := 1; i <= len(numerser); i++ {
		numerser[i] = -1
		if _, _, err := scanner.ReadLine(); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func Task_4() {

	scanner := bufio.NewReader(os.Stdin)

	qteRaft, _, err := scanner.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	qte, err := strconv.Atoi(string(qteRaft))
	if err != nil {
		fmt.Println("Can not convert this []byte to int")
		return
	}

	outputs := make([][]int, qte)

	for i := 0; i < qte; i++ {

		workerRaft, _, err := scanner.ReadLine()
		if err != nil {
			fmt.Println(err)
			return
		}

		worker, _ := strconv.Atoi(string(workerRaft))
		outputs[i] = make([]int, worker)
		tempInRom := [2]int{15, 30}

		for j := 0; j < worker; j++ {
			wish, _ := scanner.ReadBytes('\n')
			wishSign := wish[0]
			wishStep, _ := strconv.Atoi(string(wish[3:5]))

			if wishSign == '<' && tempInRom[0] > wishStep {
				hopeless(scanner, outputs[i][j:])
				break
			}

			if wishSign == '<' && tempInRom[0] <= wishStep {
				tempInRom[1] = wishStep
				outputs[i][j] = tempInRom[0]
			}

			if wishSign == '>' && wishStep > tempInRom[1] {
				hopeless(scanner, outputs[i][j:])
				break
			}

			if wishSign == '>' && wishStep <= tempInRom[1] {
				tempInRom[0] = wishStep
				outputs[i][j] = tempInRom[0]
			}
		}
	}

	for i := range outputs {
		for j := range outputs[i] {
			fmt.Println(outputs[i][j])
		}
		fmt.Println()
	}
}
