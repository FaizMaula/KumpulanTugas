package main
import "fmt"
const NMAX = 100
type pokemon struct {
	name string
	HP,Att,Def,SP_Att,SP_Def int
	AttSpeed float64
}
type tPokemon[NMAX-1] pokemon

func inputPokemon(T *tPokemon, n *int){
	var i int
	var found bool = false
	for i<NMAX-1 && !found{
		fmt.Scan(&T[i].name)
		if T[i].name != "STOP" {
			fmt.Scan(&T[i].HP,&T[i].Att,&T[i].Def,&T[i].SP_Att,&T[i].SP_Def,&T[i].AttSpeed)
		} else {
			found = true
		}
		*n++
		i++
	}
}
func PrintArr(T tPokemon,n int){
	for i:=0;i<n-1;i++{
		fmt.Println(T[i])
	}
}

func selSort(T *tPokemon, n int) {
	var i,pass,idx int
	var temp pokemon
	
	for pass = 1;pass < n;pass++ {
		idx = pass-1
		for i = pass; i < n-1;i++ {
			if T[idx].SP_Att < T[i].SP_Att {
				idx = i
			}
		}
		temp = T[pass-1]
		T[pass-1] = T[idx]
		T[idx] = temp
	}
}
func insSort(T* tPokemon, n int) {
	var j int
	var temp pokemon
	for i := 0;i<n-1;i++ {
		j = i
		temp = T[j]
		for j > 0 && temp.SP_Att > T[j-1].SP_Att {
			T[j] = T[j-1]
			j--
		}
		T[j] = temp
	}
}

func tinggirendah(T tPokemon, n int){
	var j,s,max,min int
	max = T[0].Att
	min = T[0].Def
	for i:=0;i<n-1;i++ {
		if max < T[i].Att {
			max = T[i].Att
			j = i
		} else if min > T[i].Def {
			min = T[i].Def
			s = i
		}
	}
	
	fmt.Println("Attack tertinggi ", T[j].name)
	fmt.Println("Defense terendah ", T[s].name)
}

func cariSP(T tPokemon, n int, s int) {
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

func main () {
	var s tPokemon
	var n int
	var x int
	var apa string
	inputPokemon(&s,&n)
	fmt.Println("==========")
	PrintArr(s,n)
	fmt.Scan(&apa)
	if apa == "ins"{
		fmt.Println("INS")
		insSort(&s,n)
	} else {
		fmt.Println("SEL")
		selSort(&s,n)
	}
		
	fmt.Println("==========")
	PrintArr(s,n)
	fmt.Println("==========")
	tinggirendah(s,n)
	fmt.Println("==========")
	
	fmt.Print("Cari SP");fmt.Scan(&x)
	cariSP(s,n,x)
}