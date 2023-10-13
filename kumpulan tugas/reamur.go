package main
import "fmt"
func main (){
	var celcius,reamur float64
	fmt.Scan(&celcius)
	reamur = (konversiSuhu(celcius))
	fmt.Print(reamur)
}


func konversiSuhu(celcius float64) float64 {
    /* fungsi menerima input suhu celsius dengan tipe real
       dan mengembalikan nilai suhu reamur dengan tipe real */
	var reamur float64
	reamur = 0.8 * celcius
	return reamur
}