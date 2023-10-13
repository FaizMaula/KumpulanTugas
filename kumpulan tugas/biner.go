package main
import "fmt"

type descend[1000] int

func main() {	
	var a descend
	var i,bil,x int
	var hasil bool
	var n int
	
	fmt.Print("Jumlah Angka :");fmt.Scan(&n)
	i = 1
	for i <= n {
		fmt.Scan(&a[i])
		i++
	}
	fmt.Print("Cari Angka :");fmt.Scan(&bil)
	fmt.Println(cek(a,n,bil))
	fmt.Println(urutan(a,n,bil))
	x = urutan(a,n,bil)
	hasil = cek(a,n,bil)
	if hasil {
		fmt.Print("Angka ",a[x]," ditemukan di urutan ke-",x)
	} else {
		fmt.Print("Angka tidak ditemukan")
	}
}

func cek(a descend,n,x int) bool {
	var l,r,m int
	
	r = n
	m = (l + r)/2
	for l <= r && a[m] != x {
		if x < a[m] {
			l = m + 1
		} else if x > a[m] {
			r = m - 1
		}
		m = (l + r) / 2
	}
	return m >= 0 && a[m] == x
}

func urutan(a descend, n, x int) int {
	var l,r,m,array int
	
	r = n
	m = (l+r)/2
	array = 555
	for l <= r && array == 555{
		if x == a[m] {
			array = m
		} else if x > a[m] {
			r = m - 1
		} else if x < a[m] {
			l = m + 1
		}
		m = (r+l) / 2
	}
	return array
}	