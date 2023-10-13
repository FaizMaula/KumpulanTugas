package main
import "fmt"
func main () {
	var kartu_1301223017[51] int 
	var i int
	
	fmt.Scan(&kartu_1301223017[i])
	for i<= 51 && kartu_1301223017[i] != 0 {
		i++
		fmt.Scan(&kartu_1301223017[i])
	}
	for i != 0{
		i--
		fmt.Print(kartu_1301223017[i]," ")
	}
}