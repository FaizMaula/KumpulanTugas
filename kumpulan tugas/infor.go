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
	var i,max int
	
	buatFakultas(&fif)
	daftarProdi(&fif)
	
	max = fif.arrProdi[i].aktif + fif.arrProdi[i].lulusan
	for i = 0 ; i<fif.N ;i++ {
		if max < fif.arrProdi[i].aktif + fif.arrProdi[i].lulusan{
			max = fif.arrProdi[i].aktif + fif.arrProdi[i].lulusan
		}
	}
	fmt.Println(fif.arrProdi[0].nama,fif.arrProdi[0].aktif, fif.arrProdi[0].lulusan)
	fmt.Println(fif.arrProdi[1].nama,fif.arrProdi[1].aktif, fif.arrProdi[1].lulusan)
}
func buatFakultas(f *fakultas) {
	fmt.Scan(&f.nama)
	fmt.Scan(&f.N)
}

func daftarProdi(f *fakultas){
	var nama string
	var i int
	for j := 0;j < f.N;j++ {
		fmt.Scan(&f.arrProdi[j].nama,&f.arrProdi[j].akreditasi,&f.arrProdi[j].tahun,&f.arrProdi[j].aktif,&f.arrProdi[j].lulusan)
		nama = f.arrProdi[j].nama
		i = (cekProdi(nama,*f))
		fmt.Println(nama, cekProdi(nama,*f))
		if i != -1 {
			f.arrProdi[i].aktif = f.arrProdi[i].aktif + f.arrProdi[j].aktif
			f.arrProdi[i].lulusan = f.arrProdi[i].lulusan + f.arrProdi[j].lulusan
		}
	}
}

func cekProdi(s string, f fakultas) int {
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


func termuda(f fakultas) int {
	var tahun int
	tahun = f.arrProdi[0].tahun
	for i:= 0 ; i < f.N ; i++ {
		if tahun <= f.arrProdi[i].tahun {
			tahun = f.arrProdi[i].tahun 
		}
		i++
	}
	return tahun
}
func prestasi(f fakultas) string {
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

func cetakProdi(f fakultas, x string) {
	fmt.Print("Akreditasi Prodi terbanyak di Fakultas", f.nama, "adalah", x)
	for i := 0 ; i < f.N ; i++ {
		if x == f.arrProdi[i].akreditasi {
			fmt.Print(f.arrProdi[i].nama)
		}
	}
}
