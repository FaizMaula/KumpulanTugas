package main

import "fmt"

const nmax = 1000

type himpunan struct {
	info    [nmax]string
	nelemen int
}

func main() {
	var a, b, c, d himpunan
	createset_1301220198(&a)
	createset_1301220198(&b)
	intersection_1301220198(a, b, &c)
	union_1301220198(a, b, &d)
	printset_1301220198(c)
	fmt.Println()
	a.nelemen = a.nelemen - 1
	printset_1301220198(a)
	printset_1301220198(d)
}

func createset_1301220198(set *himpunan) {
	var ketemu bool
	set.nelemen = 1
	fmt.Scan(&set.info[0])
	fmt.Scan(&set.info[1])
	ketemu = ismember_1301220198(*set, set.info[set.nelemen])
	for !ketemu {
		if !ketemu {
			set.nelemen++
			fmt.Scan(&set.info[set.nelemen])
		}
		ketemu = ismember_1301220198(*set, set.info[set.nelemen])
	}
}
func printset_1301220198(set himpunan) {
	for i := 0; i <= set.nelemen; i++ {
		fmt.Print(set.info[i], " ")
	}
}
func ismember_1301220198(set himpunan, s string) bool {
	var ketemu bool
	ketemu = false
	i := 0
	for !ketemu && i < set.nelemen {
		ketemu = set.info[i] == s
		i++
	}
	return ketemu
}
func intersection_1301220198(set1, set2 himpunan, set3 *himpunan) {
	var cek bool
	set3.nelemen = 0
	for i := 0; i < set1.nelemen; i++ {
		cek = ismember_1301220198(set2, set1.info[i])
		if cek {
			set3.info[set3.nelemen] = set1.info[i]
			set3.nelemen++
		}
	}
}
func union_1301220198(set1, set2 himpunan, set3 *himpunan) {
	var cek bool
	set3.nelemen = 0
	for i := 0; i < set2.nelemen; i++ {
		cek = ismember_1301220198(set1, set2.info[i])
		if !cek {
			set3.info[set3.nelemen] = set2.info[i]
			set3.nelemen++
		}
	}
}