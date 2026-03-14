# <h1 align="center">Laporan Praktikum Modul 2 - Algoritma Pemrograman 2 </h1>
<p align="center">Eridayalma Zahra Yohar - 109082500221</p>

## Modul 2 

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### 2a.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul2/Output/output-soal1.jpeg)
Penjelasan : 
Program ini buat baca tiga input string dari user lalu disimpan di variabel satu, dua, dan tiga. Setelah itu program nampilin dulu urutan awalnya. Selanjutnya program nuker posisi datanya pakai variabel temp sebagai penyimpanan sementara. Nilai pertama dipindah ke belakang, nilai kedua jadi pertama, dan nilai ketiga jadi kedua. Terus di akhir, program nampilin urutan baru yang udah berubah. Jadi intinya program ini cuma buat geser urutan tiga data string.

### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunanwarna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.
#### 2b.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul2/Output/output-soal2.jpeg)
Penjelasan : 
Program ini membaca 4 input warna dari user sebanyak 5 kali percobaan. Setiap input dicek apakah urutannya merah, kuning, hijau, ungu. Kalau ada yang ngga sesuai, maka berhasil menjadi false. Di akhir program nanti ditampilkan apakah input warna tersebut sesuai atau ngga.

### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
#### 2c.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul2/Output/output-soal3.jpeg)
Penjelasan :
Program ini buat ngitung biaya kirim parsel dari beratnya. User masukin berat dalam gram, lalu diubah jadi kg dan sisa gram.
Biaya 1 kg = 10.000. Kalau ada sisa gram dan berat ≤ 10 kg, sisa gram ikut dihitung. Kalau > 10 kg, sisa gramnya gratis. Terakhir program nampilin detail berat dan total biayanya.