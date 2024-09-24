package main

import "fmt"

func Ft_profit(prices []int) int {
	Min := prices[0]
	ArgentObt := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < Min {
			Min = prices[i]
		} else if prices[i]-Min > ArgentObt {
			ArgentObt = prices[i] - Min
		}
	}
	fmt.Println(ArgentObt)
	return ArgentObt
}

func main() {
	Ft_profit([]int{7, 1, 5, 3, 6, 4}) // resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	Ft_profit([]int{7, 6, 4, 3, 1}) // resultat : 0
}
