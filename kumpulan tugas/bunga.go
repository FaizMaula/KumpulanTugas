package main
import "fmt"

type tabBunga[4] string

func main (){
	var t tabBunga
	var n = 4
	var bunga string
	var hapus,ganti string
	for i:= 0 ;i<=3;i++{ 
		fmt.Scan(&t[i])
	}
	fmt.Print("Ingin mengganti bunga?"); fmt.Scan(&ganti)
	if ganti == "ya" || ganti == "YA" {
		fmt.Print("Bunga yang ingin diganti:");fmt.Scan(&bunga)
		rename_1301223017(&t,n,bunga)
		fmt.Println(t)
	}
	fmt.Print("Ingin menghapus bunga?"); fmt.Scan(&hapus)
	if hapus == "ya" || hapus == "YA" {
		fmt.Print("Bunga yang ingin dihapus:");fmt.Scan(&bunga)
		delete_1301223017(&t,n,bunga)
		fmt.Println(t[0:3])
	} else {
		fmt.Print(t)
	}
	
	
	
}

func cari_1301223017(t tabBunga, n int, bunga string) int {
	var i, found int
	
	found = -1
	for i < n && found == -1 {
		if t[i] == bunga {
			found = i
		}
		i = i + 1
	}
	return found
}

func rename_1301223017(t *tabBunga, n int, X string) {
	var found int
	found = cari_1301223017(*t,n,X)
	if found == -1 {
		fmt.Println("Bunga tidak ditemukan")
	} else {
		fmt.Print("Diganti dengan:");fmt.Scan(&t[found])
	}
}

func delete_1301223017(t *tabBunga, n int, X string){
	var found,i int
	
	found = cari_1301223017(*t,n,X)
	if found == -1 {
		fmt.Println("Bunga tidak ditemukan")
	} else {
		i = found
		for i < n-1 {
			t[i] = t[i+1]
			i++
		}
		n--
	}
}