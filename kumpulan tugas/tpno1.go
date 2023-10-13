package main
import (
	"fmt"
	"math"
)

type point struct {
	x1, y1, x2, y2 float64
}

func jarak_1301223017(p1, p2 point) float64 {
        return math.Sqrt(math.Pow((p1.x1 - p2.x2), 2) + math.Pow((p1.y1 - p2.y2), 2))
}

func main() {
	var titik point
	var ab, cd float64
	var xx,yx, yy,xy, x3,y3,x4,y4 float64
	

	fmt.Scanln(&titik.x1, &titik.y1, &titik.x2, &titik.y2)
	ab = jarak_1301223017(titik, titik)
	xx = titik.x1
	xy = titik.y1
	yy = titik.x2
	yx = titik.y2
	
	fmt.Scanln(&titik.x1, &titik.y1, &titik.x2, &titik.y2)
	cd = jarak_1301223017(titik, titik)
	x3 = titik.x1
	y3 = titik.y1
	x4 = titik.x2
	y4 = titik.y2

	if ab > cd {
		fmt.Print("Titik terdekat adalah titik C(", x3,",", y3,") dengan D(", x4,",",y4, ") dengan jarak ");fmt.Printf("%.1f",cd);fmt.Print("." )
	} else {
		fmt.Print("Titik terdekat adalah titik A(", xx,",", xy,") dengan B(", yy,",",yx, ") dengan jarak ");fmt.Printf("%.1f",ab);fmt.Print("." )
	}
}
