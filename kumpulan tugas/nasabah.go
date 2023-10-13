package main

import "fmt"

type nasabah struct {
    idNasabah string
    nama      string
    bank      string
    rek       int
}

type tabNasabah [2022]nasabah

func addNasabah(T *tabNasabah, N int) {
    for i := 0; i < N; i++ {
        fmt.Scan(&T[i].idNasabah, &T[i].nama, &T[i].bank, &T[i].rek)
        if i == 2021 {
            fmt.Print("data penuh")
        }
    }
}

func cetak(T tabNasabah, N int, X string) {
    fmt.Scan(&X)
    for i := 0; i < N; i++ {
        if T[i].bank == X {
            fmt.Printf("%s %s %s %d\n", T[i].idNasabah, T[i].nama, T[i].bank, T[i].rek)
        }
    }
}

func main() {
    var T tabNasabah
    var N int = 10
    var X string

    addNasabah(&T, N)
	fmt.Scan(X)
    cetak(T,N,X)
}