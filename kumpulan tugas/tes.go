package main

import "fmt"

type student struct {
	name                            string
	math, indo, eng, sains, average float64
}

type arrStudent [2048]student

func entryStudent_1301223069(T *arrStudent, n *int) {
	var found bool = false
	for !found {
		fmt.Scan(&T[*n].name)
		
		if T[*n].name == "STOP" {
			found = true
		} else {
			fmt.Scan(&T[*n].math, &T[*n].indo, &T[*n].eng, &T[*n].sains)
		}
		*n++
	}
}

func calculateAverage_1301223017(T *arrStudent,n int) {
	for i := 0 ; i < n ; i ++ {
		T[i].average = (T[i].math + T[i].sains + T[i].indo + T[i].eng) / 4
	}
}

func max(T arrStudent, n int, flag string) int {
	var max float64
	var j int
	if flag == "math" {
		max = T[0].math
		for i := 1 ;i<n-1;i++ {
			if max < T[i].math {
				max = T[i].math
				j = i
			}
		}
	} else if flag == "sains" {
		max = T[0].sains
		for i := 1 ;i<n-1;i++ {
			if max < T[i].sains {
				max = T[i].sains
				j = i
			}
		}
	} else if flag == "eng" {
		max = T[0].eng
		for i := 1 ;i<n-1;i++ {
			if max < T[i].eng {
				max = T[i].eng
				j = i
			}
		}
	} else {
		max = T[0].indo
		for i := 1 ;i<n-1;i++ {
			if max < T[i].indo {
				max = T[i].indo
				j = i
			}
		}
	}
	return j
}

func rangking(T *arrStudent, n int){
	var pass,idx,i int
	var temp student
	
	pass = 1
	for pass <= n-1 {
		idx = pass-1 // 0
		i = pass // 1
		for i < n {
			if T[idx].average > T[i].average {
				idx = i
			}
			i++
		}
		temp = T[pass-1]
		T[pass-1] = T[idx]
		T[idx] = temp
		pass++
	
	}
}

func printTop3(T *arrStudent, n int) {
	for i:=1;i<4;i++{
		fmt.Println("Rangking",i+1,": ", T[i].name," rata-rata ", T[i].average)
	}
}

func printmax(T arrStudent, n int) {
	fmt.Println("Nilai Matematika tertinggi diraih oleh ",T[max(T,n,"math")].name)
	fmt.Println("Nilai Bahasa Indonesia tertinggi diraih oleh ",T[max(T,n,"ind")].name)
	fmt.Println("Nilai Bahasa Inggris tertinggi diraih oleh ",T[max(T,n,"eng")].name)
	fmt.Println("Nilai IPA/IPS tertinggi diraih oleh ",T[max(T,n,"sains")].name)
}
func main() {
	var s arrStudent
	var n int	
	
	entryStudent_1301223069(&s, &n)
	calculateAverage_1301223017(&s,n)
	rangking(&s,n)
	
	printTop3(&s,n)
	printmax(s,n)
	
	
	
}

