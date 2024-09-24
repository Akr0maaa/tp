package main

func Ft_non_overlap(intervals [][]int) int {

}

func main() {
	Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}) // resultat : 1
	// pour que les intervalles soient tous des intervalles qui ne se superpose pas,
	// il suffit de d'enlever qu'un seul intervalle, [1,3]
	Ft_non_overlap([][]int{{1, 2}, {2, 3}})         // resultat : 0
	Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}) // resultat : 2
}
