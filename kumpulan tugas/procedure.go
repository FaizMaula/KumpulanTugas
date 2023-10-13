package main
import "fmt"

func main(){
	var x int
	fmt.Scan(&x)
	cariTTL(&x)
	fmt.Print(x)
}

func cariTTL(x *int){
	*x = *x / 10000
	*x = *x % 1000000
}