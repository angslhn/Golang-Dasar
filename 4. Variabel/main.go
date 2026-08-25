package main

import "fmt"

// Variabel tingkat package: wajib memakai kata kunci var,
// deklarasi singkat := tidak diperbolehkan di luar fungsi.
var namaAplikasi = "Latihan Go"

// Blok var digunakan untuk mengelompokkan beberapa variabel sekaligus.
var (
	versi  = 1
	aktif  = true
	pemilik string // tanpa nilai awal, otomatis berisi string kosong
)

func main() {
	// 1. Deklarasi dengan var dan tipe data eksplisit.
	var namaMahasiswa string = "Aang"
	var jumlahMataKuliah int = 6
	fmt.Println("Nama:", namaMahasiswa, "| Mata kuliah:", jumlahMataKuliah)

	// 2. Deklarasi dengan var tanpa tipe: tipe ditentukan compiler (type inference).
	var kota = "Bandung"
	var suhu = 27.5
	fmt.Printf("Kota: %s (tipe %T), suhu: %.1f (tipe %T)\n", kota, kota, suhu, suhu)

	// 3. Deklarasi singkat dengan := (hanya boleh di dalam fungsi).
	jurusan := "Informatika"
	semester := 5
	fmt.Println("Jurusan:", jurusan, "| Semester:", semester)

	// 4. Zero value: nilai bawaan jika variabel tidak diberi nilai awal.
	var angkaBulat int
	var angkaDesimal float64
	var status bool
	var catatan string
	fmt.Printf("Zero value -> int: %d, float64: %.1f, bool: %t, string: %q\n",
		angkaBulat, angkaDesimal, status, catatan)

	// 5. Deklarasi banyak variabel sekaligus dengan tipe yang sama.
	var panjang, lebar int = 10, 5
	fmt.Println("Panjang:", panjang, "| Lebar:", lebar)

	// 6. Deklarasi banyak variabel dengan tipe berbeda memakai :=.
	produk, harga, tersedia := "Buku Tulis", 15000, true
	fmt.Println("Produk:", produk, "| Harga:", harga, "| Tersedia:", tersedia)

	// 7. Mengubah nilai variabel: gunakan operator = (tanpa :=).
	jumlahMataKuliah = 7
	fmt.Println("Jumlah mata kuliah setelah diperbarui:", jumlahMataKuliah)

	// 8. Aturan := : minimal satu variabel di sisi kiri harus baru.
	//    Pada baris berikut, tinggi adalah variabel baru, sedangkan panjang
	//    hanya diperbarui nilainya.
	panjang, tinggi := 20, 8
	fmt.Println("Panjang:", panjang, "| Tinggi:", tinggi)

	// 9. Variabel tingkat package dapat diakses dari dalam fungsi.
	pemilik = "Aang Solihin"
	fmt.Printf("Aplikasi %s versi %d, aktif: %t, pemilik: %s\n",
		namaAplikasi, versi, aktif, pemilik)
}