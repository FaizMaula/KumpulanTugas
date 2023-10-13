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
	for A.nElemen > 0 {
		fmt.Print(A," ")
		A.nElemen--
	}
	
	
}

func createSet(set *himpunan) {
	var s string
	var ada bool = false
	for set.nElemen < 1000 && !ada {
		fmt.Scan(&set.info[set.nElemen])
		s = set.info[set.nElemen]
		ada = isMember(*set,s)
		set.nElemen++
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
	for set1.nElemen > 0 {
		for set1.info[set1.nElemen] != set2.info[set2.nElemen] {
			set2.nElemen--
			if set1.info[set1.nElemen] == set2.info[set2.nElemen]{
				set3.info[set3.nElemen] = set1.info[set1.nElemen] 
				set3.nElemen++
			}
		}
		set1.nElemen--
	}
} 
func union(set1, set2 himpunan, set3 *himpunan){
	var s string
	for set1.nElemen > 0 {
		s = set1.info[set1.nElemen]
		if isMember(*set3,s) == false {
			set3.info[set3.nElemen] = set1.info[set1.nElemen]
			set3.nElemen++
		}
		set1.nElemen--
	}	
	for set2.nElemen > 0 {
		s = set2.info[set1.nElemen]
		if isMember(*set3,s) == false {
			set3.info[set3.nElemen] = set2.info[set1.nElemen]
			set3.nElemen++
		}
		set2.nElemen--
	}	
}