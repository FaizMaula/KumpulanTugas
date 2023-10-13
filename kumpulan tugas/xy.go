package main
import "fmt"
func main () {
	var a,b int
	fmt.Scan(&a,&b)
	(ganjil(&a,&b))
}

func ganjil(x,y *int) {
	if *x > *y {
		fmt.Print(*y,*x)
	} else if *x < *y {
		fmt.Print(*x,*y)
	}
}