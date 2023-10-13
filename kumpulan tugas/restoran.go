package main

import "fmt"

func main (){
	var M, i, jumlahMenu, banyakOrang, total int
	var sisa bool
	fmt.Scan(&M)
	for i=0; i<M; i++ {
		fmt.Scan(&jumlahMenu,&banyakOrang, &sisa)
		hitungTarif_1301223226(jumlahMenu,sisa,banyakOrang, &total)
		fmt.Println(total)
	}
}

func tarif_1301223226(menu int) int {
	if menu <= 3 {
		return 10000
	} else if menu > 3 && menu <= 50 {
		return 10000 + (menu - 3) * 2500 // menu: 5 = 10000 + 2500 + 2500
	} else {
		return 100000
	}
}

func hitungTarif_1301223226(menu int, bersisa bool, n int, total *int) {
	if bersisa{
		*total = tarif_1301223226(menu) * n
	} else {
		*total = tarif_1301223226(menu)
	}
}