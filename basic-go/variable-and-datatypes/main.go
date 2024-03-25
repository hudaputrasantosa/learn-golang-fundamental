package main

import "fmt"

func main() {
	// ====================VARIABEL====================================================================

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

	fmt.Println("\nHalo, namaku ", firstname, lastname, ", umurku", umur, " lahir tahun", tahunLahir)

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

	// ====================TIPE DATA====================================================================

	var positifNumber uint8 = 255
	var negatifNumber int8 = -100
	decimal := 2.66
	var exist bool = true
	boolCondition := 5 <= 2
	const bigNumber uint = 9999
	// cant replace value
	// bigNumber = 10000

	fmt.Printf("bilangan positif : %d\n", positifNumber)
	fmt.Printf("bilangan negatif : %d\n", negatifNumber)
	fmt.Printf("bilangan decimal : %f\n", decimal)
	fmt.Printf("bilangan decimal bulat : %.3f\n", decimal)
	fmt.Printf("boolean exist : %t\n", exist)
	fmt.Printf("boolean condition 5 <= 2? %t\n", boolCondition)
	fmt.Printf("Big Number : %d\n", bigNumber)

	//multi constanta
	const (
		years = 2002
		today  = "senin"
		myNumber uint = 10
		isRain bool = false
	)
}
