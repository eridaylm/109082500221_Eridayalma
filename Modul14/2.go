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
