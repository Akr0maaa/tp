package main

import "fmt"

func Ft_max_substring(s string) int {
	character := make(map[byte]int)
	debut := 0
	maxLength := 0

	for fin := 0; fin < len(s); fin++ {
		char := s[fin]

		if index, found := character[char]; found && index >= debut {
			debut = index + 1
		}

		character[char] = fin

		currentLength := fin - debut + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	fmt.Println(maxLength)
	return maxLength
}

func main() {
	Ft_max_substring("abcabcbb") // resultat : 3
	Ft_max_substring("bbbbb")    // resultat : 1
}
