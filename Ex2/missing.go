package main

import "fmt"

func Ft_missing(nums []int) int {
	n := len(nums)

	expected := n * (n + 1) / 2

	actual := 0
	for _, num := range nums {
		actual += num
	}
	return expected - actual
}

func main() {
	fmt.Println(Ft_missing([]int{3, 1, 2}))                   // Résultat : 0
	fmt.Println(Ft_missing([]int{0, 1}))                      // Résultat : 2
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // Résultat : 8
}
