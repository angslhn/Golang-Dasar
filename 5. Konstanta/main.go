package main

import "fmt"

// Konstanta tingkat package tanpa tipe (untyped constant).
const namaAplikasi = "Latihan Go"

// Konstanta dengan tipe eksplisit (typed constant).
const batasPercobaan int = 3

// Blok const digunakan untuk mengelompokkan beberapa konstanta sekaligus.
const (
	pi        = 3.14159
	mataUang  = "Rupiah"
	modeDebug = false
)

// Konstanta dapat dibentuk dari ekspresi konstanta lain,
// selama nilainya dapat dihitung saat kompilasi.
const (
	jumlahJamSehari  = 24
	jumlahMenitSejam = 60
	jumlahMenitSehari = jumlahJamSehari * jumlahMenitSejam
)

// iota adalah penghitung otomatis yang dimulai dari 0
// dan bertambah 1 pada setiap baris di dalam blok const.
const (
	statusMenunggu = iota // bernilai 0
	statusDiproses        // bernilai 1
	statusSelesai         // bernilai 2
)

// iota dapat dikombinasikan dengan ekspresi, misalnya agar dimulai dari 1.
const (
	senin = iota + 1 // bernilai 1
	selasa           // bernilai 2
	rabu             // bernilai 3
)

// Konstanta tanpa tipe dapat dipakai pada beberapa tipe yang kompatibel.
const faktorPengali = 2

// Konstanta bertipe eksplisit hanya dapat dipakai bersama tipe yang sama.
const faktorBulat int = 2

func main() {
	// 1. Konstanta tanpa tipe dan dengan tipe eksplisit.
	fmt.Println("Aplikasi:", namaAplikasi)
	fmt.Println("Batas percobaan masuk:", batasPercobaan)

	// 2. Konstanta dari blok const.
	fmt.Println("Mata uang:", mataUang, "| Mode debug:", modeDebug)

	// 3. Konstanta pi dipakai dalam perhitungan dengan variabel bertipe float64.
	jariJari := 7.0
	luasLingkaran := pi * jariJari * jariJari
	fmt.Printf("Luas lingkaran berjari-jari %.1f adalah %.2f\n", jariJari, luasLingkaran)

	// 4. Konstanta hasil ekspresi dari konstanta lain.
	fmt.Println("Jumlah menit dalam sehari:", jumlahMenitSehari)

	// 5. Penggunaan iota.
	fmt.Println("Status ->", statusMenunggu, statusDiproses, statusSelesai)
	fmt.Println("Hari ->", senin, selasa, rabu)

	// 6. Fleksibilitas konstanta tanpa tipe: faktorPengali dapat dipakai
	//    pada variabel int maupun float64 tanpa konversi.
	var jumlahBuku int = 10
	var beratBuku float64 = 2.5
	fmt.Println("Jumlah buku dikali pengali:", jumlahBuku*faktorPengali)
	fmt.Println("Berat buku dikali pengali:", beratBuku*faktorPengali)

	// 7. Konstanta lokal di dalam fungsi juga diperbolehkan.
	const salam = "Selamat belajar Go"
	fmt.Println(salam)

	// 8. Konstanta bertipe eksplisit tidak dapat dicampur dengan tipe lain.
	//    Baris berikut menyebabkan error karena faktorBulat bertipe int,
	//    sedangkan beratBuku bertipe float64:
	//    fmt.Println(beratBuku * faktorBulat)

	// 9. Nilai konstanta tidak dapat diubah.
	//    Baris berikut menyebabkan error "cannot assign to batasPercobaan":
	//    batasPercobaan = 5

	// 10. Konstanta tidak dapat dideklarasikan dengan :=.
	//     Baris berikut menyebabkan error sintaks:
	//     const tinggi := 10
	
	_ = faktorBulat
}