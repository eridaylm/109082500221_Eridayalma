package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	var berhasil bool = true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)

		if warna1 != "merah" {
			berhasil = false
		}
		if warna2 != "kuning" {
			berhasil = false
		}
		if warna3 != "hijau" {
			berhasil = false
		}
		if warna4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}