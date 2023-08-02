package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

type persegi struct {
	sisi float64
}

// lingkaran
func (l lingkaran) jariJari() float64 { // l adalah receiver (penerima) yang merupakan objek dari struct lingkaran // jariJari adalah method yang mengembalikan float64
	return l.diameter / 2
}

func (l lingkaran) luas() float64 { // luas adalah method yang mengembalikan float64
	return math.Pi * math.Pow(l.jariJari(), 2) // math.Pow adalah pangkat
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// persegi
func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	// versi 1
	bangunDatar = lingkaran{15.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())

	// versi 2
	var bangunDatar2 hitung = lingkaran{15.0}
	var bangunLingkaran lingkaran = bangunDatar2.(lingkaran)
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar2.luas())
	fmt.Println("keliling  :", bangunDatar2.keliling())
	fmt.Println("jari-jari :", bangunLingkaran.jariJari())

	// perbedaan versi 1 dan 2 adalah pada versi 1, variabel bangunDatar langsung di-assign dengan struct lingkaran sehingga tidak perlu di-casting lagi pada saat memanggil method jariJari()
	// sedangkan pada versi 2, variabel bangunDatar di-assign dengan struct lingkaran melalui variabel bangunLingkaran sehingga perlu di-casting lagi pada saat memanggil method jariJari()
	// kesimpulan : casting dilakukan pada saat variabel di-assign dengan struct yang berbeda

}
