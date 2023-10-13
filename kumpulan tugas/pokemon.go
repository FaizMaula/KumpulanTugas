package main
import "fmt"
type poke struct{
	nama string
	cp,hp,iv_atk,iv_def,iv_hp,total int
	berat, tinggi float64
}
type arrpoke [100]poke

func inputArr(t *arrpoke, n *int){
	t[0].nama,t[0].cp,t[0].hp,t[0].iv_atk,t[0].iv_def,t[0].iv_hp,t[0].total,t[0].berat,t[0].tinggi = "Drifblim",1708,222,15,7,(82-15-7),82,8.04,1.05
	*n++
	t[1].nama,t[1].cp,t[1].hp,t[1].iv_atk,t[1].iv_def,t[1].iv_hp,t[1].total,t[1].berat,t[1].tinggi = "Chansey", 695,297,13,15,(86-15-13),86,45.9,1.29
	*n++
	t[2].nama,t[2].cp,t[2].hp,t[2].iv_atk,t[2].iv_def,t[2].iv_hp,t[2].total,t[2].berat,t[2].tinggi = "Blissey", 1132,263,11,10,(77-10-11),77,49.35,1.65
	*n++
	t[3].nama,t[3].cp,t[3].hp,t[3].iv_atk,t[3].iv_def,t[3].iv_hp,t[3].total,t[3].berat,t[3].tinggi = "Hariyama", 2360,231,9,15, (86-15-9),86, 156.99, 1.89
	*n++
	t[4].nama,t[4].cp,t[4].hp,t[4].iv_atk,t[4].iv_def,t[4].iv_hp,t[4].total,t[4].berat,t[4].tinggi = "Slowking", 2410, 178,13,15, (95-15-13),95, 81.23,2.04
	*n++
	t[5].nama,t[5].cp,t[5].hp,t[5].iv_atk,t[5].iv_def,t[5].iv_hp,t[5].total,t[5].berat,t[5].tinggi = "Wailmer", 1007,195,9,15,(86-15-9),86,209.49,2.5
	*n++
}
func max(t arrpoke,n int){
	max := t[0].nama
	min := t[0].nama
	for i:=0;i<=n;i++{
		if max < t[i].nama {
			max = t[i].nama
		} else if min > t[i].nama {
			min = t[i].nama
		}
	}
	fmt.Println("Min :",min)
	fmt.Println("Max :",max)
}
func selSort(t *arrpoke,n int) {
	var temp poke
	pass := 1
	for pass <= n {
		idx := pass - 1
		i := idx
		for i <= n {
			if t[idx].nama > t[i].nama {
				idx = i
			}
			i++
		}
		temp = t[pass-1]
		t[pass-1] = t[idx]
		t[idx] = temp
		pass++
	}
}
func binarySearch(t arrpoke, n int, x string) {
	var kr,kn,mid int
	var found bool = false
	kr = 0
	kn = n-1
	for kr<= kn && !found {
		mid = (kr+kn)/2
		if t[mid].nama > x {
			kn = mid - 1
		} else if t[mid].nama < x {
			kr = mid + 1
		} else {
			found = true
		}
	}
	if found {
		fmt.Println("ini y", t[mid])
	}
}

func main(){
	var n int
	var t arrpoke
	var x string
	inputArr(&t,&n)
	selSort(&t,n)
	fmt.Scan(&x)
	binarySearch(t,n,x)
	for i:=0 ;i<=n;i++{
		fmt.Println(t[i])
	}
}