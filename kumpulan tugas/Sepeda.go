package main
import "fmt"
func main(){
	var j,m int
	var mem bool
	fmt.Scan(&j,&m,&mem)
	if durasi_1301223226(j,m) > 3 {
		if mem {
			fmt.Print(((durasi_1301223226(j,m)) * 3500)-((durasi_1301223226(j,m) * 3500)*10)/100)
		} else {
			fmt.Print(((durasi_1301223226(j,m)) * 5000)-((durasi_1301223226(j,m) * 5000)*10)/100)
		}
	} else {
		if mem {
			fmt.Print((durasi_1301223226(j,m)) * 3500)
		} else {
			fmt.Print((durasi_1301223226(j,m)) * 5000)
		}
	}
	
	
}

func durasi_1301223226(jam, menit int)int{
		if menit >0 && jam == 0 {
			return 1
		} else if menit < 10 && jam != 0 {
			return jam
		} else if  menit == 0 && jam == 0{
			return 0
		} else {
			return jam + 1
		}
		
}