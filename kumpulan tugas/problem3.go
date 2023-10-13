package main
import "fmt"
func main (){
	var i,total int
	
	for i = 0 ; i <= 1000000000000 ; i++ {
		if i % 6 == 0 || i % 9 == 0 || i % 15 == 0 || i % 20 == 0 {
			total = total + 1
		}
	}
	fmt.Print(total) 
}