package main

import "fmt"

func main() {
	var beratGr, kg, sisaGr, biayaKg, biayaSisa, total int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratGr)

	kg = beratGr / 1000
	sisaGr = beratGr % 1000

	biayaKg = kg * 10000

	if sisaGr > 0 {
		if kg > 10 {
			biayaSisa = 0
		} else {
			if sisaGr >= 500 {
				biayaSisa = sisaGr * 5
			} else {
				biayaSisa = sisaGr * 15
			}
		}
	} else {
		biayaSisa = 0
	}

	total = biayaKg + biayaSisa

	fmt.Println("Detail berat:", kg, "kg +", sisaGr, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", total)
}