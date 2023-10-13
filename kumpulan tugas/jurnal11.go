package main
import "fmt"

const NMAX int = 1000

type himpunan struct {
	info [NMAX] string
	nElemen int
}

func main (){
	var A,B,C,D himpunan

	createSet(&A)
	createSet(&B)
	intersection(A,B,&C)
	union(A,B,&D)
	printSet(C)
	fmt.Println()
	for i := 1 ; i < A.nElemen  ; i++ {
		fmt.Print(A.info[i]," ")
	}
	printSet(D)
	
}

func createSet(set *himpunan) {
	
	var s string
	var ada bool = false
	for set.nElemen < 1000 && !ada {
		set.nElemen++
		fmt.Scan(&set.info[set.nElemen])
		s = set.info[set.nElemen]
		ada = isMember(*set,s)
		
	}
}
func printSet(set himpunan) {
	for i := set.nElemen-1 ; i >= 0  ; i-- {
		fmt.Print(set.info[i]," ")
	}
}
func isMember(set himpunan, s string) bool {
	var ada bool = false
	for set.nElemen > 0 && !ada {
		set.nElemen--
		ada = set.info[set.nElemen] == s
	}
	return ada
}
func intersection(set1, set2 himpunan, set3 *himpunan) {
	var i int
	set1.nElemen--
	set2.nElemen--
	for set1.nElemen > 0 {
		i = set2.nElemen
			for i > 0 {
				if set1.info[set1.nElemen] == set2.info[i]{
					set3.info[set3.nElemen] = set1.info[set1.nElemen] 
					set3.nElemen++
				}
				i--
			}
		set1.nElemen--
	}
}
func union(set1, set2 himpunan, set3 *himpunan) {
	var cek,cek2 bool
	set3.nElemen = 0
	for i := set2.nElemen ; 0 < i; i-- {
		cek = isMember(set1, set2.info[i])
		if !cek  && !cek2 {
			set3.info[set3.nElemen] = set2.info[i]
			set3.nElemen++
		}	
	}
	if set3.nElemen > 1 {
		set3.info[0] = " "
	}
}
