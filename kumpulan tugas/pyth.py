print("hello world")
#p = 10

p = int(input("Enter a number"))
if p == 20 :
    print("p is 20")
elif p != 20:
    print("p is not 20") 
#q = 20
q = int(input("Enter a number"))
lu = p * q
ke = 2 * (p + q)
print(lu,ke)


'''
package main

import (
	"fmt"
)

func main() {
	var p, l, lu, ke, di int

	fmt.Scan(&p, &l)

	lu = p * l
	ke = 2 * (p + l)
	di = math.sqrt((p * p) + (l * l))

	fmt.Println(lu, ke, di)
}
'''
