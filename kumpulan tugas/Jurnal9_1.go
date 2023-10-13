package main
import "fmt"
func main (){
	var kota string
	var populasi int
	var besar float64
	var kota1,kota2 string
	
	fmt.Scan(&kota,&populasi,&besar,&kota1,&kota2)	
	fmt.Printf("%s %s %s %.2f %s" ,"Kerapatan penduduk kota", kota, "adalah", float64(populasi)/besar, "orang per km2")
}