package main

import "fmt"

// Contoh aturan sintaks dasar Go.
func main() {
	// 1. Tanda titik koma ditambahkan otomatis oleh compiler,
	//    sehingga tidak perlu ditulis secara manual.
	fmt.Println("Baris ini tidak menggunakan titik koma secara manual")

	// 2. Kurung kurawal buka wajib berada di baris yang sama dengan pernyataannya.
	//    Jika { dipindahkan ke baris baru, kode berikut akan menyebabkan error:
	//    func main()
	//    {
	//    }

	// 3. Go membedakan huruf besar dan huruf kecil (case-sensitive).
	nama := "Aang Solihin"
	Nama := "AANG SOLIHIN"
	fmt.Println(nama, Nama)

	// 4. Exported dan unexported (contoh dalam satu file).
	hitungLokal := hitungSesuatu()
	fmt.Println("Hasil fungsi unexported:", hitungLokal)

	fmt.Println("Fungsi exported:", HitungSesuatu())
}

// hitungSesuatu: huruf awal kecil berarti unexported
// (hanya dapat diakses dalam package ini).
func hitungSesuatu() int {
	return 10
}

// HitungSesuatu: huruf awal kapital berarti exported
// (dapat diakses dari package lain).
func HitungSesuatu() int {
	return 100
}