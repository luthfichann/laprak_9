package main

import ("fmt")

func main() {
	var jumlahOrang, jumlahMotor int

	fmt.Scan(&jumlahOrang)
	jumlahMotor = jumlahOrang / 2
	jumlahOrang = jumlahOrang % 2
	if jumlahOrang != 0 {
		jumlahMotor = jumlahMotor + 1
	}
	fmt.Println(jumlahMotor)
}