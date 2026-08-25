package main

import "fmt"

// SayHalo mencetak salam ke terminal.
// Ini contoh komentar dokumentasi untuk fungsi exported (huruf besar pertama).
func SayHello() {
	fmt.Println("Halo dari Go!")
}

/*
Ini contoh multi-line comment.
Biasa dipakai buat penjelasan panjang,
atau buat "menonaktifkan" beberapa baris kode sementara waktu.
*/

func main() {
	// Single-line comment - penjelasan singkat di kode
	SayHello()

	// x adalah contoh variabel biasa (unexported, tidak butuh doc comment formal)
	x := 5
	fmt.Println("Nilai x:", x)
}