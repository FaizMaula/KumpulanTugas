package main
import "fmt"
const NMAX = 1000
type negara struct {
	nama string
	tahun,gdp int
	gdp21,gdp22,gdp23 int
}
type arrNegara [NMAX]negara  

func inputArr(T *arrNegara, n *int){
	var ase bool = true
	for *n < NMAX && ase {
		fmt.Scan(&T[*n].nama)
		ase = asean(T[*n].nama)
		if ase {
			fmt.Scan(&T[*n].tahun,&T[*n].gdp)
		}
		*n++
	}
}

func asean(x string) bool {
	return x == "Brunei" || x == "Singapore" || x == "Cambodia" || x == "Philippines" || x == "Malaysia" || x == "Indonesia" || x == "Vietnam" || x == "Thailand" || x == "Myanmar" || x == "Laos"
}

func itungGDP(T *arrNegara, n int ){
	for i :=n ;i >0 ; i--{
		for j := 0 ;j<n; j++ {
			if T[i].nama == T[j].nama {
				if T[j].tahun == 2021 {
					T[i].gdp21 += T[j].gdp
				} else if T[j].tahun == 2022 {
					T[i].gdp22 += T[j].gdp
				} else {
					T[i].gdp23 += T[j].gdp
				}
			}
		}
	}
}

func sort(T* arrNegara, n int, x int){
	if x == 2021 {
		for pass := 1 ; pass < n;pass++{
			idx := pass -1
			for i := pass;i<n;i++{
				if T[idx].gdp21 < T[i].gdp21 {
					idx = i	
				}
			}
			T[pass-1],T[idx] = T[idx], T[pass-1]
		}
	} else if x == 2022 {
		for pass := 1 ; pass < n;pass++{
			idx := pass -1
			for i := pass;i<n;i++{
				if T[idx].gdp22 < T[i].gdp22 {
					idx = i
				}
			}
			T[pass-1],T[idx] = T[idx], T[pass-1]
		}
	} else if x == 2023 {
		for pass := 1 ; pass < n;pass++{
			idx := pass -1
			for i := pass;i<n;i++{
				if T[idx].gdp23 < T[i].gdp23 {
					idx = i
				}
			}
			T[pass-1],T[idx] = T[idx], T[pass-1]
		}
	}
}
func sequential_search(T arrNegara, n,i int, x string) int {
	var j int
	j =  -1
	for i>0 && j == -1  {
		i--
		if T[i].nama == x {
			j = i
		}
	}
	return j
}
func printArr(T arrNegara, n int) {
	for i:=0 ;i <n-1;i++{
		idx := sequential_search(T,n,i,T[i].nama)
		if i == 0 || idx == -1 {
			fmt.Println(T[i].nama, T[i].gdp21,T[i].gdp22,T[i].gdp23)
		}
	}
}
func main (){
	var s arrNegara 
	var n,x int
	inputArr(&s,&n)
	fmt.Println("Urut berdasarkan GDP pada tahun :");fmt.Scan(&x)
	itungGDP(&s,n)
	fmt.Println(n)
	sort(&s,n,x)
	//delete(&s,n)
	fmt.Println("=====================")
	printArr(s,n)
}
