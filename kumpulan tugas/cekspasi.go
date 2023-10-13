package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan input dengan spasi: ")
	input, _ := reader.ReadString('\n')

	fmt.Println("Input yang dimasukkan:", input)
}