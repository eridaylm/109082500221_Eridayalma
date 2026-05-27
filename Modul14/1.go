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
