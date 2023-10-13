package main
import (
	"fmt"
	"math"
)

type point struct {
	x1, y1, x2, y2 float64
}

func jarak(p1, p2 point) float64 {
        return math.Sqrt(math.Pow((p1.x1 - p2.x2), 2) + math.Pow((p1.y1 - p2.y2), 2))
}

func main() {
	var titik point
	var ab, cd, ax, ay, bx, by, cx, cy, dx, dy float64
	var kalimat1, kalimat2, dot, koma, tkurung, dengan, djarak, b, d string

	fmt.Scanln(&titik.x1, &titik.y1, &titik.x2, &titik.y2)
	ab = jarak(titik, titik)
	ax = titik.x1
	ay = titik.y1
	bx = titik.x2
	by = titik.y2

	fmt.Scanln(&titik.x1, &titik.y2, &titik.x2, &titik.y2)
	cd = jarak(titik, titik)
	cx = titik.x1
	cy = titik.y1
	dx = titik.x2
	dy = titik.y2
	
	kalimat1 = "Titik terdekat adalah titik A("
	kalimat2 = "Titik terdekat adalah titik C("
	tkurung = ")"
	dot = "."
	koma = ","
	dengan = "dengan"
	djarak = "dengan jarak"
	b = "B("
	d = "D("

	if ab < cd {
		fmt.Printf("%s%.0f%s%.0f%s %s %s%.0f%s%.0f%s %s %.1f%s", kalimat1, ax, koma, ay, tkurung, dengan, b, bx, koma, by, tkurung, djarak, ab, dot)
	} else {
		fmt.Printf("%s%.0f%s%.0f%s %s %s%.0f%s%.0f%s %s %.1f%s", kalimat2, cx, koma, cy, tkurung, dengan, d, dx, koma, dy, tkurung, djarak, cd, dot)
	}
}
