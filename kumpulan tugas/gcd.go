package main
import "fmt"
func main (){
	for i:=1000000000 ;1000000000<10000000000;i++{
		for j :=1000000000;1000000000 < 10000000000;j++ {
			for i%j > 0 {
				selesai := i % j
				i = j
				j = selesai
				if selesai == 666 {
					fmt.Println(i,j)
				}
			}
		}
	}
}x