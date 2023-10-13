package main
import "fmt"
func main (){
	var c float64
	fmt.Scan(&c)
	temperatur(c)
}

func temperatur(suhu float64) {
	var r,f,k float64
	r = suhu * 0.8
	f = suhu * 1.8  + 32
	k = suhu + 273.15
	fmt.Print(r,"R ",f,"F ",k,"K")
}