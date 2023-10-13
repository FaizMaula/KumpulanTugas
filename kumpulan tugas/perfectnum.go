package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	display_1301223226(a, b)
}

func isFaktor_1301223226(num, x int) bool {
	return num%x == 0
}
func jumlahFaktor_1301223226(num int, total *int) {
	var i int
	for i = 1; i < num; i++ {
		if isFaktor_1301223226(num, i) {
			*total = *total + i
		}
	}
}

func perfect_1301223226(num int) bool {
	var hasilJumlahFaktor int

	jumlahFaktor_1301223226(num, &hasilJumlahFaktor)

	return num == hasilJumlahFaktor
}

func display_1301223226(x, y int) {
	var i int

	fmt.Print("Barisan Perfect Number : ")
c
	for i = x; i <= y; i++ {
		if perfect_1301223226(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}