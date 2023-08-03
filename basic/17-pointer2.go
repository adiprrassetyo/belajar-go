package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
} // *Address artinya kita mengambil alamat memori dari Address dan kita akan mengubah semua value dari Address

func main() {

	// address
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// var address4 *Address = &Address{"Subang", "Jawa Barat", "Indonesia"} // ini juga bisa
	address2 := &address1             //versi 1 //untuk merubah value dari address1 kita membutuhkan tanda & untuk mengambil alamat memori dari address1
	var address3 *Address = &address1 //versi 2

	address2.City = "Bandung" //jika hanya ingin mengubah city nya saja maka kita tidak perlu menggunakan tanda *
	// tetapi jika kita ingin mengubah semua value nya maka kita harus menggunakan *
	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"} //versi 1 //ini hanya akan mengubah address2 saja atau address1 tidak akan berubah
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} //versi 2 //ini akan mengubah semua value dari address1 dan address2 juga akan berubah karena kita menggunakan tanda *

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	var alamatPointer *Address = &alamat    // maka kita akan mengubah semua value dari alamat
	ChangeCountryToIndonesia(alamatPointer) // dan fungsi dapat mengubah semua value dari alamat
	fmt.Println(alamat)

}
