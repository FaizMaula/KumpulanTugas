package main
import "fmt"
func main (){

var huruf byte
var konsonan bool

fmt.Scanf("%c", &huruf)
konsonan = hurufKonsonan(huruf)
fmt.Println(konsonan)

}

func hurufKonsonan(huruf byte) bool {
	var konsonan bool
	if huruf != 'a' && huruf != 'i' && huruf != 'u' && huruf != 'e' && huruf != 'o' {
	    konsonan = true
	} else {
	    konsonan = false
	}
	return konsonan
}