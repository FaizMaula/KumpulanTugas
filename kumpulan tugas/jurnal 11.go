package main
import "fmt"

const NMAX int = 1000

type himpunan struct {
	info [NMAX] string
	nElemen int
}

func main (){
	var A,B himpunan

	createSet(&A)
	createSet(&B)
	
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
