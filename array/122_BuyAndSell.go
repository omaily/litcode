package array

func BestTimetoBuyAndSell() {
	var size int = 10
	nums := generateRandomSlise(size)
	maxProfit(nums)
	printSlise(nums)
}

func maxProfit(prices []int) int {
	var summ int
	for i := 0; i < len(prices); i++ {

		var profit int
		for j := i + 1; j < len(prices); j++ {
			difference := prices[j] - prices[i]
			if profit <= difference {
				profit = difference
			} else {
				println(prices[i], " - ", prices[j-1], " = ", profit)
				summ += profit
				// если отнимим больше 1 попадем в петлю
				// еденицу отнимаем чтобы цикл №1 итерировался до нужного значения
				i = j - 1
				// без break цикл №2 не завершится => profit не обнулится
				break
			}
			if j == len(prices)-1 {
				println(prices[i], " - ", prices[j], " = ", profit)
				summ += profit
				i = j // если отнимим больше 1 попадем в петлю
			}
		}
	}
	return summ
}
