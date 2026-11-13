package main

import ("fmt")

func main() {
	var aksi string

	fmt.Println("1. mobil")
	fmt.Println("2. motor")
	fmt.Scan(&aksi)

	if aksi == "1" {
		fmt.Println("Farezki akan telat masuk kuliah.")
		fmt.Println("Farezki akan mendapatkan pengurangan nilai.")
	}
	if aksi == "2" {
		fmt.Println("Baju Farezki akan basah.")
		fmt.Println("Farezki harus membawa baju salin.")
	}
}