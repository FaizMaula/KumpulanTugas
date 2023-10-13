package main
import "fmt"

const N = 1000
type angka[N-1] int

func inputAngka(tab *angka, i *int) {
	fmt.Scan(&tab[*i])
	for *i<N-1 && tab[*i]!= 666 {
		*i++
		fmt.Scan(&tab[*i])
	}
}
func InsertionSort(T *angka, n int){
	var j int
	var temp int
	for i:=0;i<n;i++{
		j = i
		temp = T[j]
		for j > 0 && temp > T[j-1] {
			T[j] = T[j-1]
			j--
		}
		T[j]= temp
	}
}
func sequentialsearch(T angka, n , x int){
	i := 0
	ketemu := false
	for i<n && !ketemu {
		if x == T[i] {
			ketemu = true
		}
		i++
	}
	fmt.Println(ketemu)
}
func main (){
	var t angka
	var n,x int
	inputAngka(&t,&n)
	InsertionSort(&t,n)
	for i:=0;i<n;i++{
		fmt.Println(t[i])
	}
	fmt.Println("Cari Angka");fmt.Scan(&x)
	sequentialsearch(t,n,x)
}
