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
