package main
import "fmt"
func main (){
	var hariAju, hariAmbil string 
	var tglAju, tglAmbil int 
	var blnAju, blnAmbil int

	fmt.Println("Pengajuan paspor: ")
	fmt.Println("Hari: "); fmt.Scan(&hariAju) 
	fmt.Println("Tanggal: "); fmt.Scan(&tglAju) 
	fmt.Println("Bulan: "); fmt.Scan(&blnAju)

	
	tglAmbil = tglSelesai(tglAju,hariAju)
	hariAmbil = hariSelesai(hariAju)
	updateBulan(tglAju,tglAmbil,blnAju,&blnAmbil)
	fmt.Println("Pengambilan paspor: ", hariAmbil, tglAmbil, "/", blnAmbil)

}
func hariSelesai(d string) string {
    if d == "Senin" || d == "Monday" {
       return "Rabu"
    } else if d == "Selasa" || d == "Tuesday"{
       return "Kamis"
    } else if d == "Rabu" || d == "Wednesday"{
       return "Jumat"
    } else if d == "Kamis" || d == "Thursday"{
       return "Senin"
    } else { 
       return "Selasa"
	}
}
func tglSelesai(d1 int, d2 string) int {
    var selesai int  
	if d2 == "Kamis"|| d2 == "Jumat"{
		selesai = d1 + (2*2)
	} else { 
		selesai = d1 + 2
	}
	if selesai > 30 {
		selesai = selesai % 10
	}
	return selesai
}
func updateBulan(d1, d2 int,b1 int, b2 *int) {
    if d1 > d2 {
		*b2 = b1 + 1
		if *b2 > 12 {
          *b2 = 1
		}
	} else {
		*b2 = b1
	}
}