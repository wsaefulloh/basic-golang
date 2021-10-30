package main

import (
	"fmt"
	"strconv"
)

func main() {
	pembulatan(4.37)
	pembulatan(4.32)
	pembulatan(4.324)
}

func pembulatan(value float64) {
	var i = 0
	var nilai_perseratusan = value
	for i >= 0 {
		nilai_perseratusan = nilai_perseratusan - 0.1
		if nilai_perseratusan < 0.1 {
			nilai_perbandingan := nilai_perseratusan - 0.05
			if nilai_perbandingan >= 0.0 {
				var nilai_tambah = 0.1 - nilai_perseratusan
				var nilai_pembulatan = value + nilai_tambah
				var str = strconv.FormatFloat(nilai_pembulatan, 'f', 2, 64)
				fmt.Println(str)
				break
			} else {
				var nilai_pembulatan = value - nilai_perseratusan
				var str = strconv.FormatFloat(nilai_pembulatan, 'f', 2, 64)
				fmt.Println(str)
				break
			}
		}
		i++
	}
}
