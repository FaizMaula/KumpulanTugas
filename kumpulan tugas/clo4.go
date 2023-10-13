package main
import "fmt"
func main() {
	for i:=0;i>-1000;i--{
		for j:=0 ; j<1000;j++{
			if (i*801) + (j*81) == 9 {
				fmt.Println(i,j)
			}
		}
	}
}