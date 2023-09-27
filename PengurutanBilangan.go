package main

import "fmt"

func bubblesort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Println(arr)
			fmt.Println("Nilai yang ditukar : ", arr[j], "dg", arr[j+1])
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				fmt.Println("tidak ada penukaran")
			}
		}
		fmt.Println("Jumlah iterasi ke ", i+1)
	}
}

func main() {
	var names [5]string
	names[0] = "pertama"
	names[1] = "kedua"
	names[2] = "ketiga"
	names[3] = "keempat"
	names[4] = "kelima"

	var nomor1 = []int{3, 5, 9, 8, 6}

	fmt.Println(names)
	fmt.Println(nomor1)
	fmt.Println("Panjang array: ", len(nomor1))

	var nomor2 = []int{3, 5, 9, 8, 6}
	fmt.Println("data sebelum ditambah: ", nomor2)
	nomor2 = append(nomor2, 4, 5, 6)
	fmt.Println("data setelah ditambah: ", nomor2)

	var angka = []int{9, 2, 1, 8, 4, 1}
	fmt.Println("data sebelum diurutkan: ", angka)
	bubblesort(angka)

	fmt.Println("data setelah diurutkan :", angka)

}
