package main
import "fmt"
const nMax = 2022
type student struct {
	name,sid = string
	gpa = float64
}
type tabMhs [nMax-1] student
func main (){
	var t tabMhs
	for n := 0 ;n<9;n++{
		fmt.Scan(&t[n])
	}
	selSort(&t,n)
	fmt.Println(t)
}
func selSort(t *tabMhs,n int ) {
	for pass := 0 ;pass<n;pass++ {
		idx := pass-1
		for i:=pass;i<n;i++{
			if t[idx].gpa < t[i].gpa {
				idx = i
			}
		}
		t[pass-1],t[idx] = t[idx],t[pass-1]
	}
	
	

}