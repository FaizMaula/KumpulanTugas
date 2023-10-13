package main
import "fmt"
type kota struct {
	nama, kota1, kota2 string
	pop int
	besar, density float64
}

type jKota[100] kota

func main (){
	var a jKota
	var i,n int
	var total float64
	
	fmt.Scan(&a[i].nama,&a[i].pop,&a[i].besar,&a[i].kota1,&a[i].kota2)
	a[i].density = float64(a[i].pop)/a[i].besar
	for a[i].nama != "AKHIR" && a[i].pop != 0 && a[i].besar != 0.0 && a[i].kota1 != "AKHIR" && a[i].kota2 != "AKHIR"{
		i++
		n++
		fmt.Scan(&a[i].nama,&a[i].pop,&a[i].besar,&a[i].kota1,&a[i].kota2)
		a[i].density = float64(a[i].pop)/a[i].besar
	}	
	for i = 0 ; i < n ;i++ {
		total += a[i].density
	}
	fmt.Printf("%s %.2f %s\n", "Kerapatan rata-rata penduduk kota adalah", total/float64(n),"orang per km2") 
	for i=0 ;i <n ; i++ {
		if a[i].density < total/float64(n) {
			fmt.Print(a[i].nama)
			
			if a[i].density < total/float64(n) && i != n-1 {
				fmt.Print(", ")
			}
		}
		
	}
}
