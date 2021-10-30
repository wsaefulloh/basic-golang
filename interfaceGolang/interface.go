package main

import "fmt"

//INTERFACE HITUNG 2D

type hitung2d interface {
	luas() float64
	keliling() float64
}

//INTERFACE HITUNG 3D

type hitung3d interface {
	volume() float64
}

//INTERFACE HITUNG

type hitung interface {
	hitung2d
	hitung3d
}

//BANGUN PERSEGI----------------------------------------------

type persegiPanjang struct {
	panjang float64
	lebar   float64
}

func (sisi persegiPanjang) luas() float64 {
	return sisi.lebar * sisi.panjang
}

func (sisi persegiPanjang) keliling() float64 {
	return (sisi.panjang * 2) + (sisi.lebar * 2)
}

//BANGUN BALOK------------------------------------------------

type balok struct {
	panjang float64
	lebar   float64
	tinggi  float64
}

func (sisi balok) volume() float64 {
	return sisi.lebar * sisi.panjang * sisi.tinggi
}

func (sisi balok) luas() float64 {
	pxl := sisi.panjang * sisi.lebar
	lxt := sisi.tinggi * sisi.lebar
	txp := sisi.panjang * sisi.tinggi
	return 2 * (pxl + lxt + txp)
}

func (sisi balok) keliling() float64 {
	return 4 * (sisi.lebar + sisi.panjang + sisi.tinggi)
}

//MAIN--------------------------------------------------

func main() {
	var bangunDatar hitung2d = persegiPanjang{2, 4}
	fmt.Println(bangunDatar.luas())
	fmt.Println(bangunDatar.keliling())

	var bangun3D hitung3d = balok{2, 4, 5}
	fmt.Println(bangun3D.volume())

	var hitungall hitung = balok{2, 4, 5}
	fmt.Println(hitungall.keliling())
	fmt.Println(hitungall.luas())
	fmt.Println(hitungall.volume())
}
