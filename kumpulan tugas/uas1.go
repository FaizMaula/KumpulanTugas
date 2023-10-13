package main
import "fmt"
func main(){
	var s string
	var nom,x,sep,goc,ser int
	
	fmt.Print("Jenis Valas: ");fmt.Scan(&s)
	fmt.Print("Nominal Valas: ");fmt.Scan(&nom)
	if s == "USD" {
		x = nom * 14000
	} else if s == "GBP" {
		x = nom * 18500
	}
	fmt.Println	("Nominal IDR: ",x)
	ser = x / 100000
	x = x - (100000*ser)
	sep = x / 50000
	x = x - (50000*sep)
	goc = x / 10000
	fmt.Println("RP. 100.000.- : ",ser)
	fmt.Println("RP. 50.000.- : ",sep)
	fmt.Println("RP. 10.000.- : ",goc)
}