package main

import "fmt"

func main() {
	var aksi string

	fmt.Println("1. mobil")
	fmt.Println("2. motor")
	fmt.Scan(&aksi)

	if aksi == "1" {
		fmt.Println(" akan telat masuk kuliah.")
		fmt.Println(" akan mendapatkan pengurangan nilai.")
	}
	if aksi == "2" {
		fmt.Println("Baju Farezki akan basah.")
		fmt.Println(" Farezki harus membawa baju salin.")
	}
}

// program Konsekuensi
// kamus :
// 	aksi : string
// algoritma :
// 	input(aksi)
// 	if aksi == 1 then
// 		output("akan telat masuk kuliah.")
// 		output("akan mendapatkan pengurangan nilai.")
// 	endif
// 	if aksi == 2 then
// 		output("Baju Farezki akan basah.")
// 		output("Farezki harus membawa baju salin.")
// 	endif
// endprogram