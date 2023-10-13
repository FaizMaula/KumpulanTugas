package main
import "fmt"
func main() {
	var num, num1, num2 int
	
	fmt.Scanf("%d", &num)
	
	split_1301223226(num, &num1,&num2)
	
	fmt.Printf("%d %d \n%d", num1, num2, (num1+num2))
}

func len_1301223226(num int) int {
	var counter int
	for num > 0{
		counter++
		num = num / 10
	}
	
	return counter
}

func pangkat_1301223226(n int) int {
	if n == 0 {
		return 1
	} else {
		return 10 * pangkat_1301223226(n-1)
	}
}
func split_1301223226(num int, num1,num2 *int) {
	var mid int
	
	mid = len_1301223226(num) / 2
	
	*num1 = num/ pangkat_1301223226(mid)
	*num2 = num % pangkat_1301223226(mid)
}