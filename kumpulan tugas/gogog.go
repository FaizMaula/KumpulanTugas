package main
import "fmt"
func main (){
	for i := 0 ;i<1104;i++ {
		total := (45*i + 50) - 90
		if total % 977 == 0 {
			fmt.Println(i)
		}
	}
}