package main

import "fmt"

func main() {
	jam := 0
	var totaljam, total int

	for jam != -1 {
		fmt.Print("Masukkan jam: ")
		fmt.Scan(&jam)
		totaljam += jam

	}
	total = 5000 + (totaljam-1)*3000
	fmt.Println("Jumlah jam parkir: ", totaljam)
	fmt.Println("Total bayar: ", total)

}
