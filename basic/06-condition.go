package main

import "fmt"

func main() {

	/*
		Go tidak mendukung seleksi kondisi menggunakan ternary.
		Statement seperti: var data = (isExist ? "ada" : "tidak ada")
		adalah invalid dan menghasilkan error.
	*/


	// seleksi kondisi menggunakan keyword if, else if, dan else
	var point = 8 //lulus

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}
		/*
		pemrograman lain, ketika ada seleksi kondisi yang isi blok-nya hanya 1 baris saja, kurung kurawal boleh tidak dituliskan. Berbeda dengan aturan di Go, kurung kurawal harus tetap dituliskan meski isinya hanya 1 blok satement.
		*/


	// variable temporary pada if else
	var nilai = 8840.0

	if percent := nilai / 100; percent >= 100 {
    fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
    fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
    fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
		/*
		deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang menggunakan tanda :=. Penggunaan keyword var di situ tidak diperbolehkan karena akan menyebabkan error.
		*/


	// seleksi kondisi menggunakan keyword switch dan case
	var point2 = 6

	switch point2 { // switch dengan kondisi
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// switch dengan short statement
	absen := 90
	score := 90

	switch { // switch tanpa kondisi
	case absen >= 90 && score >= 75:
		fmt.Println("Lulus")
	default:
		fmt.Println("Tidak Lulus")
	}

	// kurung kurawal pada switch case
	var point3 = 6

	switch point3 {
	case 8:
			fmt.Println("perfect")
	case 7, 6, 5, 4:
			fmt.Println("awesome")
	default:
			{
					fmt.Println("not bad")
					fmt.Println("you can be better!")
			}
	}

	// switch dengen gaya if else
	var point4 = 6

	switch {
	case point4 == 8:
			fmt.Println("perfect")
	case (point4 < 8) && (point4 > 3):
			fmt.Println("awesome")
	default:
			{
					fmt.Println("not bad")
					fmt.Println("you need to learn more")
			}
	}

	// switch dengan fallthrough
	var point5 = 6

	switch {
	case point5 == 8:
			fmt.Println("perfect")
	case (point5 < 8) && (point5 > 3):
			fmt.Println("awesome")
			fallthrough //akan tetap memaksa menjalankan case selanjutnya tanpa menghiraukan nilai kondisi, meskipun nilai kondisi bernilai false.
	case point5 < 5:
			fmt.Println("you need to learn more")
	default:
			{
					fmt.Println("not bad")
					fmt.Println("you need to learn more")
			}
	} // hasilnya akan mencetak awesome, you need to learn more, karena ada keyword fallthrough


	// seleksi	kondisi bersarang
	var point6 = 10

	if point6 > 7 {
			switch point6 {
			case 10:
					fmt.Println("perfect!")
			default:
					fmt.Println("nice!")
			}
	} else {
			if point6 == 5 {
					fmt.Println("not bad")
			} else if point6 == 3 {
					fmt.Println("keep trying")
			} else {
					fmt.Println("you can do it")
					if point6 == 0 {
							fmt.Println("try harder!")
					}
			}
	}


		/*
		switch pada pemrograman Go memiliki perbedaan dibanding bahasa lain. Di Go, ketika sebuah case terpenuhi, tidak akan dilanjutkan ke pengecekan case selanjutnya, meskipun tidak ada keyword break di situ. Konsep ini berkebalikan dengan switch pada umumnya, yang ketika sebuah case terpenuhi, maka akan tetap dilanjut mengecek case selanjutnya kecuali ada keyword break.
		*/

}