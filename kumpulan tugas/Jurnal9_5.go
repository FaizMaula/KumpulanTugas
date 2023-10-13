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
	var x string
	
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
	
	fmt.Scan(x)
	cetakTerdekat_1301223017(a,n,x)
}
 
func cetakTerdekat_1301223017(a jKota, n int, x string){
	
	var i int
	
	fmt.Scan(&x)
	for i=0 ;i <n ; i++ {
		if a[i].kota1 == x || a[i].kota2 == x {
			fmt.Print("Kota yang berdekatan dengan ", x, " adalah ", a[i].nama)
			for i=i+1 ;i <n ; i++{
				if a[i].kota1 == x || a[i].kota2 == x {
					fmt.Print(" dan ", a[i].nama)
				}
			}
		}
	}
}