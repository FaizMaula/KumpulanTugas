package main
import "fmt"
type arr[0..n-1] int

func main() {
	var i int
	var s arr
	s[0] = 1
	s[1] = 2
	
	for i < 2 {
		s[i]++
		i++
	}
	fmt.Print(s	)
}