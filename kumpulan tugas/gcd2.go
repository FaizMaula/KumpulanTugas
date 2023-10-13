package main
import "fmt"
func main (){
	for i:=100 ;i<10000;i++{
		for j := 101 ;j<10000;j++ {
			for j%i > 0 {
				selesai := j%i
				j = i
				i = selesai
				if selesai == 145 {
					fmt.Println(i,j)
				}
			}
		}
	}
}