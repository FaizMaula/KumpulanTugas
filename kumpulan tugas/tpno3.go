package main
import "fmt"

type AsDos_T struct {
	kode , nama ,MK string
	nim , jumlah int
}

type TabAsDos_T [2500]AsDos_T

func main (){
	
	var tabel TabAsDos_T
    var n int
    var mk string

    BacaAsDos_1301223017(&tabel, &n)
	fmt.Scan(mk)
    cetakAsDos_1301223017(tabel,n,mk)
	
}

func BacaAsDos_1301223017(tabel *TabAsDos_T, n *int){	
	var i int
	 for i<=2500 && tabel[i].kode != "XXX" {
		i++
        fmt.Scan(&tabel[i].kode, &tabel[i].nama, &tabel[i].MK, &tabel[i].nim, &tabel[i].jumlah)
		*n++
    }
}

func cetakAsDos_1301223017(tabel TabAsDos_T, n int, mk string) {
    fmt.Scan(&mk)
    for i := 0; i < n; i++ {
        if tabel[i].MK == mk {
            fmt.Printf("%s %s\n", tabel[i].nama, tabel[i].kode)
        }
    }
}