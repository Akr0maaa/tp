package main

import "fmt"

func Ft_coin(coins []int, amount int) int {
	difference := make([]int, amount+1)
	for i := range difference {
		difference[i] = amount + 1
	}

	difference[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				if difference[i-coin]+1 < difference[i] {
					difference[i] = difference[i-coin] + 1
				}
			}
		}
	}

	if difference[amount] > amount {
		return -1
	}

	return difference[amount]
}

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3 car (11 == 5 + 5 + 1)
	fmt.Println(Ft_coin([]int{2}, 3))        // resultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // resultat : 0
}
