package main
import "fmt"

const NMAX = 12345

type prodi struct {
	nama, akreditasi string
	tahun, aktif, lulusan int
}
type fakultas struct {
	nama string
	
	N int
}
type arrProdi[NMAX-1] prodi
func main (){
	var p arrProdi
	var fif fakultas
	var j,k,max int
	var aktif int
	var nama string
	
	fmt.Scan(&fif.nama)
	for j = 0 ;j<9;j++ {
		fmt.Scan(&p[j].nama, &p[j].akreditasi, &p[j].tahun, &p[j].aktif, &p[j].lulusan)
	}
	
	for j = 0 ;j < 9; j++ {
		nama = p[j].nama
		for j = 0 ;j < 9; j++ {
			if p[j].nama == nama {
				aktif = aktif + (p[j].aktif + p[j].lulusan)
			}
		}
		for aktif > max {
			nama = p[j].nama
			for j = 0 ;j < 9; j++ {
				if p[j].nama == nama {
					max = max + (p[j].aktif + p[j].lulusan)
				}
			}
			j++
		}
	}
	fmt.Print("Prodi ",p[k].nama ," memiliki mahasiswa dan lulusan terbanyak yaitu ", aktif)
	
}