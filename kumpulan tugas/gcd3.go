package main
import "fmt"
func main (){
	for i := 1000 ;i<1104;i++ {
		total := i + 10
		if total % 105 == 0 {
			fmt.Println(i)
		}
	}
}