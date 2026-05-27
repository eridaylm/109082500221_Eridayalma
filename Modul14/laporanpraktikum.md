# <h1 align="center">Laporan Praktikum Modul 14 - Algoritma Pemrograman 2 </h1>
<p align="center">Eridayalma Zahra Yohar - 109082500221</p>

## Modul 14

### 1.Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. 



#### 1.go

```go
package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func selectionSort(T *arrInt, n int) {
	var t, i, j, idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for d := 0; d < n; d++ {
		var m int
		fmt.Scan(&m)

		var rumah arrInt
		for k := 0; k < m; k++ {
			fmt.Scan(&rumah[k])
		}

		selectionSort(&rumah, m)

		for k := 0; k < m; k++ {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[k])
		}
		fmt.Println()
	}
}



```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul14/Output/output-soal1.png)
Penjelasan : 
Program ini digunakan buat ngurutin nomor rumah kerabat Hercules di tiap daerah dari nilai terkecil ke data terbesar (ascending) pake algoritma Selection Sort. Di setiap daerah, program nerima input m buat nentuin jumlah rumah, terus baca semua nomor rumah ke array. Fungsi selectionSort itu kerjanya nyari elemen nilai terkecil dari posisi indeks saat ini sampe akhir array, terus nukerin sama elemen yang di posisi depan (i-1).

```go
func selectionSort(T *arrInt, n int) { // tanda bintang * di T *arrInt ini pake pointer (pass by reference). Ini fungsinya biar perubahan posisi angka pas di-sort di dalem fungsi ini bisa langsung ngerubah data array asli (rumah) yang ada di fungsi main, alo ngga ada bintangnya array di main ga bakal keganti apa-apa
    var t, i, j, idx_min int
    i = 1
    for i <= n-1 {
        idx_min = i - 1 // perulangan ini sengaja dibatesin sampai n-1 aja karna kalo sisa satu angka terakhir di ujung array, angka itu otomatis udah pasti ada di posisi yang bener. Terus di baris idx_min = i - 1, anggap kalo angka pertama dari sisa array yang belum urut itu nilai terkecil sementara
        j = i
        for j < n {
            if T[idx_min] > T[j] {
                idx_min = j //Ini Selection Sort, caranya nyari angka paling kecil ngecek satu-satu ke kanan. Kalo ketemu angka T[j] yang ternyata lebih kecil dari nilai di T[idx_min], langsung ganti posisi indeks nilai terkecilnya jadi ke indeks j itu
            }
            j = j + 1
        }
        t = T[idx_min]
        T[idx_min] = T[i-1]
        T[i-1] = t // ini namanya proses swapping atau tuker nilai. Disini pake variabel bantuan t buat wadah sementara, angka paling kecil yang udah ketemu tadi (T[idx_min]) bakal dituker posisinya ke depan sub-array yang lagi dicek
        i = i + 1
    }
}
```
####


### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil. 

#### 2.go

```go
package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func selectionSortAsc(T *arrInt, n int) {
	var t, i, j, idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func selectionSortDesc(T *arrInt, n int) {
	var t, i, j, idx_max int
	i = 1
	for i <= n-1 {
		idx_max = i - 1
		j = i
		for j < n {
			if T[idx_max] < T[j] {
				idx_max = j
			}
			j = j + 1
		}
		t = T[idx_max]
		T[idx_max] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for d := 0; d < n; d++ {
		var m int
		fmt.Scan(&m)

		var ganjil, genap arrInt
		nGanjil := 0
		nGenap := 0

		for k := 0; k < m; k++ {
			var x int
			fmt.Scan(&x)
			if x%2 != 0 {
				ganjil[nGanjil] = x
				nGanjil++
			} else {
				genap[nGenap] = x
				nGenap++
			}
		}

		selectionSortAsc(&ganjil, nGanjil)
		selectionSortDesc(&genap, nGenap)

		first := true
		for k := 0; k < nGanjil; k++ {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[k])
			first = false
		}
		for k := 0; k < nGenap; k++ {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(genap[k])
			first = false
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul14/Output/output-soal2.png)
Penjelasan : 
Program ini modifikasi dari soal pertama, yang di mana program bakal misahin nomor rumah ganjil dan genap ke dua array yang beda saat proses input data.  Kalo seluruh data udah terinput dan terpisah ke kelompoknya masing-masing, proses pengurutan nya gini:  
1. Nomor Rumah Ganjil: Diurutkan dari nilai terkecil ke terbesar (ascending) pake fungsi selectionSortAsc.  
2. Nomor Rumah Genap: Diurutkan dari nilai terbesar ke terkecil (descending) pake fungsi selectionSortDesc. 
Kalo kedua kelompok array udah selesai diurutin, program langsung print nomor rumah ganjil dulu, terus dilanjutin sama nomor rumah genap di satu baris yang sama

```go
for k := 0; k < m; k++ {
    var x int
    fmt.Scan(&x)
    if x%2 != 0 {
        ganjil[nGanjil] = x
        nGanjil++
    } else {
        genap[nGenap] = x
        nGenap++
    }
} //ini buat pilah data di fungsi main baca input nomor rumah, pake kondisi if x%2 != 0 buat mendeteksi bilangan ganjil. Kalo bener, angka dimasukin ke array ganjil. Kalo salah (genap), angka masuk ke array genap. Masing-masing array punya counter sendiri (nGanjil dan nGenap) buat mencatat jumlah elemen

// Sorting ganjil (Ascending)
if T[idx_min] > T[j] { idx_min = j } //kondisi pembanding adalah > (lebih besar dari). Artinya, program nyari nilai terkecil (idx_min) buat digeser ke depan biar hasilnya urut membesar

// Sorting genap (Descending)
if T[idx_max] < T[j] { idx_max = j } //kondisi pembanding dibalik menjadi < (lebih kecil dari). Program nyari nilai terbesar (idx_max) buat digeser ke depan biar hasilnya urut mengecil
```

### 3.  Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba buat menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba buat menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

#### 3.go

```go
package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func insertionSort(T *arrInt, n int) {
	var i, j, temp int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func main() {
	var data arrInt
	n := 0
	var x int

	for {
		fmt.Scan(&x)
		if x == -5313 {
			break
		}
		if x == 0 {
			insertionSort(&data, n)
			median := 0
			if n%2 != 0 {
				median = data[n/2]
			} else {
				median = (data[n/2-1] + data[n/2]) / 2
			}
			fmt.Println(median)
		} else {
			data[n] = x
			n++
		}
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul14/Output/output-soal3.png)
Penjelasan :
Program ini buat pantau dan hitung nilai tengah (median) dari data yang dimasukin. Program bakal terus terima input angka berulang kali. Kalo angkanya bukan 0 dan bukan -5313, langsung disimpen ke array. Pas yang dimasukin angka 0, fungsi insertionSort bakal dipanggil buat ngurutin semua data yang udah masuk dari yang paling kecil ke paling besar. Setelah datanya urut, baru deh hitung mediannya: kalo jumlah datanya ganjil, ambil langsung angka pas di tengahnya (data[n/2]); kalo genap, diambil dua angka yang ada di tengah, dijumlahin, terus dibagi dua pakai pembagian bilangan bulat. Prosesnya baru berhenti total kalo masuk angka -5313

```go
func insertionSort(T *arrInt, n int) {
    var i, j, temp int
    i = 1
    for i <= n-1 {
        j = i
        temp = T[j]
        for j > 0 && temp < T[j-1] {
            T[j] = T[j-1]
            j = j - 1
        }
        T[j] = temp
        i = i + 1
    }
} //tiap elemen ke-i disimpen sementara di variabel temp, terus nilai temp ini dibandingin sama elemen-elemen di sebelah kirinya (T[j-1]). Selama elemen di kiri lebih besar, elemen bakal digeser ke kanan (T[j] = T[j-1])

if n%2 != 0 {
    median = data[n/2]
} else {
    median = (data[n/2-1] + data[n/2]) / 2
} //ini buat nentuin median, kalo jumlah data (n) ganjil (n%2 != 0), median bisa di dapetin langsung dari indeks tengah, yaitu data[n/2].  Tapi kalo jumlah data genap, posisi tengah ada di antara dua elemen jadi nilai bakal diambil dari rata-rata data[n/2-1] dan data[n/2]. Karena pake tipe data int, hasil pembagian otomatis dibuletin ke bawah (pembagian integer)
```