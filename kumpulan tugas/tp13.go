package main

import "fmt"

type student struct {
	name                            string
	math, indo, eng, sains, average float64
}

type arrStudent [2048]student

func entryStudent_1301223069(T *arrStudent, n *int) {
	/*I.S data mahasiswa sudah siap pada piranti masukan
		Proses : menerima masukan selama nama != stop
	F.S array t terisi sebanyak n bilangan bulat */
	
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
	/* I.S terdefinisi array T sejumlah n bilangan bulat
		F.S field average berisi nilai rata-rata dari 4 field mata pelajaran */
	for i := 0 ; i < n ; i ++ {
		T[i].average = (T[i].math + T[i].sains + T[i].indo + T[i].eng) / 4
	}
}

func max(T arrStudent, n int, flag string) int {
	/*{menerima string flag yang berisikan mata pelajaran dan mengembalikan 
	 indeks dari array T yang memiliki nilai maksimum)
	 */
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
	/* I.S terdefinisi array T sejumlah n bilangan bulat
		F.S array T terurut secara descending berdasarkan rata-rata nilai menggunakan selection sort */
	var pass,idx,i int
	var temp student
	
	pass = 1
	for pass <= n-1 {
		idx = pass-1 // 0
		i = pass // 1
		for i < n {
			if T[idx].average < T[i].average {
				idx = i
			}
			i++
		}
		temp = T[pass-1]
		T[pass-1] = T[idx]
		T[idx] = temp
		pass++
		fmt.Print(temp)
	}
}

func printTop3(T *arrStudent, n int) {
	/* I.S terdefinisi array T sejumlah n bilangan bulat yang terurut secara descending
		F.S menampilkan 3 data siswa dengan nilai rata-rata tertinggi*/
	for i:=0;i<3;i++{
		fmt.Println("Rangking",i+1,": ", T[i].name," rata-rata ", T[i].average)
	}
}
/*
func entryStudent_1301223069(T *arrStudent, n *int) {
	var name string
	var math, indo, eng, sains float64
	fmt.Scan(&name, &math, &indo, &eng, &sains)

	T[0].name = name
	T[0].math = math
	T[0].indo = indo
	T[0].eng = eng
	T[0].sains = sains

	*n = 1

	for *n < 2048 && T[*n].name != "STOP" {
		fmt.Scan(&T[*n].name, &T[*n].math, &T[*n].indo, &T[*n].eng, &T[*n].sains)
		*n++
	}
}
*/
func main() {
	var s arrStudent
	var n,j int
	var flag string
	
	entryStudent_1301223069(&s, &n)
	calculateAverage_1301223017(&s,n)
	flag = "math"
	j = max(s,n,flag)
	fmt.Println("Nilai Matematika tertinggi diraih oleh ",s[j].name)
	flag = "ind"
	j = max(s,n,flag)
	fmt.Println("Nilai Bahasa Indonesia tertinggi diraih oleh ",s[j].name)
	flag = "eng"
	j = max(s,n,flag)
	fmt.Println("Nilai Bahasa Inggris tertinggi diraih oleh ",s[j].name)
	flag = "sains"
	j = max(s,n,flag)
	fmt.Println("Nilai IPA/IPS tertinggi diraih oleh ",s[j].name)
	
	printTop3(&s,n)
	
}