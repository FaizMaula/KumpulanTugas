package main
import "fmt"
func main (){
	var i,total int
	
	for i = 1 ; i<100 ; i++ {
		if i % 2 == 0 || i % 3 == 0 {
			fmt.Println(i)
			total++
		}
	}
	fmt.Print(total)
}