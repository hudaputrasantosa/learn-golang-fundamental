package main

import "fmt"

func main(){
	var nama string = "Huda"
	var umur uint = 22

	// Alternatif
	namaAlt := "Huda"
	umurAlt := 22

	fmt.Printf("Halo nama saya %s!\numur saya %d", namaAlt, umurAlt)

	//multi varaible
	var firstname, lastname string = "Huda Putra", "Santosa"
	var umur, tahunLahir int
	umur, tahunLahir = 22, 2002

	fmt.Println("\nHalo, namaku ", firstname, lastname,", umurku", umur ," lahir tahun", tahunLahir)

	// underscore variable sebagai variabel yang tidak perlu dipanggil (keranjang sampah)
	_ = "Saya sedang malas"
	namaku, _ := "Huda", "Putras"

	fmt.Printf("Hai, %s", namaku)

	//membuat variabel pointer dengan tipe data tertentu
	myHobbie := new(string)
	*myHobbie = "Game"
	
	fmt.Printf("\nHobiku %s", myHobbie)
	fmt.Printf("\nHobiku %s", *myHobbie)
	
	*myHobbie = "Lari"
	fmt.Printf("\nHobiku %s", *myHobbie)

}
