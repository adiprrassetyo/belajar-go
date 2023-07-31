package main

import "fmt"

func main() {
	/*
	pada go perulangan hanya for saja, tetapi meski demikian, kemampuannya merupakan gabungan for, foreach, dan while ibarat bahasa pemrograman lain.
	*/

	// perulangan dengan keyword for
	for i := 0; i < 5; i++ {
    fmt.Println("Angka", i)
	} //perulangan di atas hanya akan berjalan ketika variabel i bernilai di bawah 5, dengan ketentuan setiap kali perulangan, nilai variabel i akan di-iterasi atau ditambahkan 1 (i++ artinya ditambah satu, sama seperti i = i + 1). Karena i pada awalnya bernilai 0, maka perulangan akan berlangsung 5 kali, yaitu ketika i bernilai 0, 1, 2, 3, dan 4.

	// perulangan dengan keyword for dengan hanya kondisi
	var i3 = 0

	for i3 < 5 {
    fmt.Println("Angka", i3)
    i3++
	}

	// perulangan dengan keyword for tanpa argumen
	var i2 = 0

	for {
			fmt.Println("Angka", i2)

			i2++
			if i2 == 5 {
					break
			}
	}

	// perulangan dengan keyword for - range
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// perulangan dengan keyword for - range (lanjutan)
	var fruits2 = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits2 {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// perulangan dengan keyword break dan continue
	for i := 1; i <= 10; i++ {
    if i % 2 == 1 {
        continue
    }

    if i > 8 {
        break
    }

    fmt.Println("Angka", i)
	} // pada perulangan di atas, jika nilai i adalah bilangan ganjil, maka perulangan akan dilanjutkan ke iterasi selanjutnya (continue). Jika nilai i adalah 9, maka perulangan akan dihentikan (break). Jika nilai i adalah bilangan genap, maka perintah fmt.Println("Angka", i) akan dijalankan.

	// perulangan bersarang
	for i := 0; i < 5; i++ {
    for j := i; j < 5; j++ {
        fmt.Print(j, " ")
    }

    fmt.Println()
	} // pada perulangan di atas, variabel i akan diiterasi dari 0 hingga 4. Kemudian, variabel j akan diiterasi dari nilai i hingga 4. Jadi, jika i adalah 0, maka j akan diiterasi dari 0 hingga 4. Jika i adalah 1, maka j akan diiterasi dari 1 hingga 4. Dan seterusnya.

	// pemanfaatan label pada perulangan
	outerLoop:
	for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
					if i == 3 {
							break outerLoop
					}
					fmt.Print("matriks [", i, "][", j, "]", "\n")
			}
	}


}