package main

import "fmt"

type deretBilang struct {
	limit int
}

func Append(slice []int, data int) []int {
	m := len(slice)
	newSlice := make([]int, (m + 1))
	copy(newSlice, slice)
	newSlice[m] = data
	return newSlice
}

func (deret deretBilang) prima() []int {
	n := deret.limit
	y := 0
	array := []int{}
	for i := 0; i <= n; i++ {
		if (i%2) != 0 && (i%3) != 0 {
			array = Append(array, i)
			y = y + 1
		}
	}
	return array
}

func (deret deretBilang) ganjil() []int {
	n := deret.limit
	y := 0
	array := []int{}
	for i := 0; i <= n; i++ {
		if (i % 2) != 0 {
			array = Append(array, i)
			y = y + 1
		}
	}
	return array
}

func (deret deretBilang) genap() []int {
	n := deret.limit
	y := 0
	array := []int{}
	for i := 0; i <= n; i++ {
		if (i % 2) == 0 {
			array = Append(array, i)
			y = y + 1
		}
	}
	return array[1:y]
}

func (deret deretBilang) fibonacci() []int {
	n := deret.limit
	i := 0
	j := 1
	array := []int{1}
	for i >= 0 {
		bilangan_fibonacci := i + j
		if bilangan_fibonacci <= n {
			array = Append(array, bilangan_fibonacci)
		} else {
			break
		}
		i = j
		j = bilangan_fibonacci
	}
	return array
}

func main() {
	deret := deretBilang{40}
	fmt.Println("limit deret angka", deret.limit)
	fmt.Println("deret bilangan prima mulai dari 0 -", deret.limit, "adalah", deret.prima())
	fmt.Println("deret bilangan ganjil mulai dari 0 -", deret.limit, "adalah", deret.ganjil())
	fmt.Println("deret bilangan genap mulai dari 0 -", deret.limit, "adalah", deret.genap())
	fmt.Println("deret bilangan fibonacci mulai dari 0 -", deret.limit, "adalah", deret.fibonacci())
}
