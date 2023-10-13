package main
import "fmt"
func main (){
	var x,x1,x2 int
	
	fmt.Scan(&x)
	split(x,&x1,&x2)
	fmt.Println(x1,x2)
	fmt.Print(x1+x2)
}

func len(num int) int{
	var i int
	i = 1
	for num > 0 {
		num = num / 10
		i++
	}
	return i
}

func pangkat(n int) int {
	if n== 0 {
		return 1
	} else {
		return 10 * pangkat(n-1)
	}
}

func split(num int, num1, num2 *int){
	var digit, bagi int
	digit = len(num) / 2
	bagi = pangkat(digit)
	if digit % 2 == 0 {
		*num1 = num / bagi
		*num2 = num % bagi
	} else {
		*num1 = num / bagi
		*num2 = num % bagi
	}
	
	
}