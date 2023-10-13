package main
import "fmt"
const NMAX = 12345

type prodi struct {
	nama, akreditasi string
	tahun, aktif, lulusan int
}
type fakultas struct {
	nama string
	arrProdi[NMAX-1] prodi
	N int
}
func main () {
	var fif fakultas
	var tahun,i,max int
	var nama,akret string
	
	buatFakultas_1301223017(&fif)
	for fif.N = 0 ; fif.N<=9; fif.N++ {
		fmt.Scan(&fif.arrProdi[fif.N].nama,&fif.arrProdi[fif.N].akreditasi,&fif.arrProdi[fif.N].tahun,&fif.arrProdi[fif.N].aktif,&fif.arrProdi[fif.N].lulusan)
	}
	daftarProdi_1301223017(&fif)
	nama = fif.arrProdi[i].nama
	max = fif.arrProdi[fif.N].aktif + fif.arrProdi[fif.N].lulusan	
	for i = 0 ; i<fif.N ;i++ {
		if max < fif.arrProdi[i].aktif + fif.arrProdi[i].lulusan{
			max = fif.arrProdi[i].aktif + fif.arrProdi[i].lulusan
			nama = fif.arrProdi[i].nama
		}
	}
	fmt.Println("Prodi" ,nama, "memiliki mahasiswa dan lulusan terbanyak yaitu", max)
	tahun = termuda_1301223017(fif)
	fmt.Println("Prodi termuda_1301223017 adalah", fif.arrProdi[tahun].nama)
	akret = prestasi_1301223017(fif)
	cetakProdi_1301223017(fif,akret)
}
func buatFakultas_1301223017(f *fakultas) {
	fmt.Scan(&f.nama)
	f.N = 0
}
func daftarProdi_1301223017(f *fakultas){
	var nama string
	var i int
	for f.N = 0;f.N <= 9;f.N++ {
		nama = f.arrProdi[f.N].nama
		i = (cekProdi_1301223017(nama,*f))
		if i != -1 {
			f.arrProdi[i].aktif = f.arrProdi[i].aktif + f.arrProdi[f.N].aktif
			f.arrProdi[i].lulusan = f.arrProdi[i].lulusan + f.arrProdi[f.N].lulusan
		}
	}
}
func cekProdi_1301223017(s string, f fakultas) int {
	var j int

	j = 0
	for j < f.N && f.arrProdi[j].nama != s {
		j++
	}
	if j == f.N {
		return -1
	} else {
		return j
	}
}
func termuda_1301223017(f fakultas) int {
	var tahun,j int
	tahun = f.arrProdi[0].tahun
	j = 0
	for i:= 0 ; i <= 9 ; i++ {
		if tahun <= f.arrProdi[i].tahun {
			tahun = f.arrProdi[i].tahun 
			j = i
		}
	}
	return j
}
func prestasi_1301223017(f fakultas) string {
	var ak,x string
	var u,b,c int
	for i:= 0 ; i < f.N ; i++ {
		ak = f.arrProdi[i].akreditasi
		if ak == "baik" {
			b++
		}else if ak == "unggul" {
			u++
		} else {
			c++
		}
	}
	if (u > b && u > c) || (u > c && u > b) {
		 x = "unggul"
	} else if (b > c && b > u) || (b > u && b > c)  {
		 x = "baik"
	} else {
		 x = "cukup"
	}
	return x
}

func cetakProdi_1301223017(f fakultas, x string) {
	var nama string
	var i int
	fmt.Printf("%s %s %s %s\n", "Akreditasi Prodi terbanyak di Fakultas ", f.nama, " adalah ", x)
	for f.N = 0;f.N <= 9;f.N++ {
		nama = f.arrProdi[f.N].nama
		i = (cekProdi_1301223017(nama,f))
		if i == -1 && x == f.arrProdi[f.N].akreditasi {
			fmt.Print(f.arrProdi[f.N].nama, " ")
		}
	}
}
