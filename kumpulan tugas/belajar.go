package main
import "fmt"
type asu struct {
	name string
	x int
}
type tAsu [10]asu

func main (){
	var t tAsu
	var n int
	var x string
	var y int
	inputArr(&t,&n)
	selSort(&t,n)
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Print(binsearch(t,n,x))
	fmt.Print(t)
}

func inputArr(t *tAsu, n *int){
	t[0].name,t[0].x = "Luis", 10
	t[2].name,t[2].x = "Jose",25
	t[3].name,t[3].x = "Hector",35
	t[4].name,t[4].x = "Javier",45
	t[5].name,t[5].x = "Carlos",52
	t[6].name,t[6].x = "Juan",65
	t[7].name,t[7].x = "Aranguiz",76
	t[8].name,t[8].x = "Rodriguez",98
	t[9].name,t[9].x = "Gilbert",34
	*n = 10
} 

func selSort(T *tAsu, n int) {
	 for pass := 1 ;pass<n;pass++{
		idx:= pass-1
		for i:=pass;i<n;i++{
			if T[idx].name > T[i].name {
				idx = i
			}
		}
		T[pass-1],T[idx] = T[idx],T[pass-1]
	 }
}
func binsearch(t tAsu,n int, c string) bool {
	found := false
	kr := 0
	kn := n-1
	for kr <= kn && !found {
		mid := (kr+kn)/2
		if c > t[mid].name  {
			kr = mid + 1
		} else if c < t[mid].name {
			kn = mid - 1
		} else {
			found = true
		}
	}
	return found
}
/*func cariSP(T tPokemon, n int, s int) {
	var kr,kn,mid int
	var found bool = false
	kr = 0
	kn = n-1
	for kr <= kn && !found {
		mid = (kr+kn)/2
		if T[mid].SP_Att > s {
			kr = mid + 1
		} else if T[mid].SP_Att < s {
			kn = mid - 1
		} else {
			found = true
		}
	}
	
	if found {
		fmt.Println(T[mid].name)
	} 
	
}
*/

func search(t tAsu,n int,x string) bool {
	var found bool = false
	i := 0
	for i < n && !found {
		if t[i].name == x {
			found = true
		}	
		i++
	}
	return found
}