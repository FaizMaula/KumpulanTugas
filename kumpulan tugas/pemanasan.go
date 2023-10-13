package main
import "fmt"

type member struct {
	nama string
	nomor int
}

type tNama [10] member
func input(t *tNama, x *int) {
	t[0].nama,t[0].nomor = "Hae",100
	t[1].nama,t[1].nomor = "Haegg",34
	t[2].nama,t[2].nomor = "Haefd",3213
	t[3].nama,t[3].nomor = "Haes",535
	t[4].nama,t[4].nomor = "Haec",3314
	*x = 5

}
func search(t tNama,x int,cari string) int {
	var c int = 69
	for i:=0;i<x;i++ {
		if cari == t[i].nama {
			c = i
		}
	}
	return c
	
}
func main (){
	var tab tNama
	var x int
	var nama string
	input(&tab,&x)
	fmt.Scan(&nama)
	fmt.Print(search(tab,x,nama))
}