package main
import "fmt"
type kota struct {
	nama, kota1, kota2 string
	pop int
	besar, density float64
}
func main (){
	var a kota
	
	fmt.Scan(&a.nama,&a.pop,&a.besar,&a.kota1,&a.kota2)
	a.density = float64(a.pop)/a.besar
	fmt.Printf("%s %s %s %.2f %s" ,"Kerapatan penduduk kota", a.nama, "adalah", a.density , "orang per km2")
}