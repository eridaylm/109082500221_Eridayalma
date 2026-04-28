# <h1 align="center">Laporan Praktikum Modul 10 - Algoritma Pemrograman 2 </h1>
<p align="center">Eridayalma Zahra Yohar - 109082500221</p>

## Modul 10 

### 1.Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.



#### 1.go

```go
package main

import "fmt"

func main() {
	var n int
	var berat [1000]float64

	fmt.Print("Masukan banyak data kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f\n", min)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", max)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul10/Output/output-soal1.png)
Penjelasan : 
Program ini buat nyari berat kelinci paling kecil dan paling besar dari data yang dimasukkin, pertama program minta user masukkin jumlah kelinci yang mau ditimbang, terus minta masukkin berat tiap kelinci satu per satu abis itu semua data berat ini disimpen ke array.
Abis itu program mulai nyari nilai terkecil dan terbesar. Caranya, program anggap dulu berat kelinci pertama sebagai patokan untuk yang terkecil dan terbesar. Terus program loop dari data kedua sampai terakhir. Tiap kali nemu berat yang lebih kecil dari patokan terkecil, patokan itu diganti, begitu juga kalo nemu yang lebih besar dari patokan terbesar, patokan terbesarnya diganti
Setelah semua data dicek, program nampilin hasil berat kelinci paling kecil dan paling besar.

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.

#### 2.go

```go
package main

import "fmt"

func main() {
	var jmlikan int
	var ikanperwadah int
	var beratIkan [1000]float64

	fmt.Print("Masukan jumlah ikan (x) dan ikan per wadah (y): ")
	fmt.Scan(&jmlikan, &ikanperwadah)

	for i := 0; i < jmlikan; i++ {
		fmt.Printf("Masukan berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	jmlwadah := jmlikan / ikanperwadah
	if jmlikan%ikanperwadah != 0 {
		jmlwadah++
	}

	beratperwadah := make([]float64, jmlwadah)

	for i := 0; i < jmlikan; i++ {
		nomorWadah := i / ikanperwadah
		beratperwadah[nomorWadah] += beratIkan[i]
	}

	fmt.Print("Total berat per wadah: ")
	ttlsmuawadah := 0.0
	for i, total := range beratperwadah {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", total)
		ttlsmuawadah += total
	}
	fmt.Println()

	rataRataWadah := ttlsmuawadah / float64(jmlwadah)
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", rataRataWadah)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul10/Output/output-soal2.png)
Penjelasan : 
Program ini buat ngebagi berat ikan ke beberapa wadah dengan kapasitas tetap, terus ngitung total tiap wadah dan rata-ratanya.
Pertama, user masukkin total jumlah ikan dan berapa kapasitas ikan per wadah, terus user masukkin berat tiap ikan satu per satu ke array, program ngitung butuh berapa wadah caranya bagi jumlah ikan sama kapasitas wadah. Kalau ada sisa, wadahnya ditambah satu. Terus program bagi ikan ke wadah-wadah secara urut, rumusnya ikan ke-0 sampai ke-(kapasitas-1) masuk wadah pertama, seterusnya. Sambil ngisi, program juga jumlahin total berat per wadah. Terakhir program nampilin total berat setiap wadah, terus ngitung rata-ratanya dari semua wadah.


### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.


#### 3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(dataBerat arrBalita, jumlahData int, beratMin, beratMax *float64) {
	*beratMin = dataBerat[0]
	*beratMax = dataBerat[0]

	for i := 1; i < jumlahData; i++ {
		if dataBerat[i] < *beratMin {
			*beratMin = dataBerat[i]
		}
		if dataBerat[i] > *beratMax {
			*beratMax = dataBerat[i]
		}
	}
}

func rerata(dataBerat arrBalita, jumlahData int) float64 {
	totalBerat := 0.0
	for i := 0; i < jumlahData; i++ {
		totalBerat += dataBerat[i]
	}
	return totalBerat / float64(jumlahData)
}

func main() {
	var databalita arrBalita
	var jmlbalita int

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&jmlbalita)

	for i := 0; i < jmlbalita; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&databalita[i])
	}

	var beratMin, beratMax float64
	hitungMinMax(databalita, jmlbalita, &beratMin, &beratMax)

	rataRata := rerata(databalita, jmlbalita)

	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul10/Output/output-soal3.png)
Penjelasan :
Program ini buat nyari berat balita paling kecil, paling besar, dan rata-rata beratnya. Program minta user masukkin jumlah balita, terus masukkin berat badan setiap balita satu per satu ke array, ada dua fungsi utama yang dipake. Fungsi pertama buat nyari nilai minimum dan maksimum, caranya sama kayak program kelinci tadi yaitu bandingin satu per satu data. Fungsi kedua buat ngitung rata-rata, caranya jumlahin semua data terus dibagi jumlah elemen. Kalo udah dapet semua hasil, program nampilin tiga output: berat terkecil, berat terbesar, dan rata-rata.