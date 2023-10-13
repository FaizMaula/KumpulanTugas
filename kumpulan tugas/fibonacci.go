package main
import "fmt"

func main () {
	var x int
	fmt.Scan(&x)
	fmt.Println(fibonnacci(x))
}

func fibonnacci(n int) int {
	var x1,x2,x3 int
	x1 = 0
	x2 = 1
	if n == 1 {
		return 0
	} else {
		for i := 1 ; i < n ;i++{
			x3 = x1 + x2
			x1 = x2
			x2 = x3
		}
		return x3
	}
}