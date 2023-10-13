package main

import (
	"fmt"
	"math"
)

func main() {
	var jarib, jarik, tinggib, tinggik int
	var luasp, luask, luasb float64
	fmt.Scan(&jarib, &jarik, &tinggib, &tinggik)
	hitungluasselimut_1301223226(jarib, tinggib, &luasb)
	hitungluasselimut_1301223226(jarik, tinggik, &luask)
	luasp = luasAlas_1301223226(jarib) + luasAlas_1301223226(jarik) + luasb - luask
	fmt.Printf("%.3f\n", luasb)
	fmt.Printf("%.3f\n", luask)
	fmt.Printf("%.3f\n", luasp)
}
func luasAlas_1301223226(r int) float64 {
	return 3.14 * float64(r*r)
}
func garisPelukis_1301223226(r, t int) float64 {
	return math.Sqrt(float64(r*r + t*t))
}
func hitungluasselimut_1301223226(r int, t int, luas *float64) {
	var s float64 = garisPelukis_1301223226(r, t)
	*luas = 3.14 * float64(r) * s
}