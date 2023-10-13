package main
import "fmt"
func main () {
	var a,b int
	fmt.Scan(&a)
	ganjil(a,b)
}

func ganjil(a,n int) int {
	if a >= 0 {
		if n % 2 != 0 {
			fmt.Print(n," ")
		}
		return ganjil(a-1,n+1)
	} else {
		return n
	}
}