package main
import "fmt"
func main () {
	var a,b int
	fmt.Scan(&a)
	b = 1
	ganjil(a,a,b)
}

func ganjil(a,n,i int) int {
	
	if a != 0 {
		if n % i == 0 {
			fmt.Print(i," ")
		}
		return ganjil(a-1,n,i+1)
	} else {
		return n
	}
}