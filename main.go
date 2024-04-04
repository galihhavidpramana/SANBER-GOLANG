package main

import (
	"fmt"
	"math"
	"strings"
)

// soal formativee 8
// soal no1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * (float64(b.panjang*b.lebar) + float64(b.panjang*b.tinggi) + float64(b.lebar*b.tinggi))
}

// soal nomor 2
type phone struct {
	name   string
	brand  string
	year   int
	colors []string
}

// Definisikan interface untuk menampilkan data
type displayData interface {
	display() string
}

// Implementasikan method display untuk struct phone
func (p phone) display() string {
	return fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\nColors: %s\n", p.name, p.brand, p.year, strings.Join(p.colors, ", "))
}

//	jawaban no 3
//
// Fungsi luasPersegi dengan tipe data return interface kosong
func luasPersegi(sisi int, withText bool) interface{} {
	if sisi == 0 {
		if withText {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi
	if withText {
		return fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}

// soal nomor 4
var prefix interface{} = "hasil penjumlahan dari "
var kumpulanAngkaPertama interface{} = []int{6, 8}
var kumpulanAngkaKedua interface{} = []int{12, 14}

// gunakan seluruh variabel tersebut untuk menghasilkan output "hasil penjumlahan dari 6 + 8 + 12 + 14 = 40"
// gunakan type assertion untuk mengerjakan soal ini

func main() {

	// soal nomor 1
	segitiga := segitigaSamaSisi{alas: 5, tinggi: 6}
	fmt.Println("Segitiga Sama Sisi:")
	fmt.Println("Luas:", segitiga.luas())
	fmt.Println("Keliling:", segitiga.keliling())

	persegi := persegiPanjang{panjang: 4, lebar: 4}
	fmt.Println("\nPersegi Panjang:")
	fmt.Println("Luas:", persegi.luas())
	fmt.Println("Keliling:", persegi.keliling())

	tabung := tabung{jariJari: 3, tinggi: 5}
	fmt.Println("\nTabung:")
	fmt.Println("Volume:", tabung.volume())
	fmt.Println("Luas Permukaan:", tabung.luasPermukaan())

	balok := balok{panjang: 3, lebar: 4, tinggi: 5}
	fmt.Println("\nBalok:")
	fmt.Println("Volume:", balok.volume())
	fmt.Println("Luas Permukaan:", balok.luasPermukaan())

	// jawaban nomor 2
	// Membuat objek phone
	myPhone := phone{
		name:   "iPhone 13",
		brand:  "Apple",
		year:   2021,
		colors: []string{"Black", "White", "Blue"},
	}
	// Memanggil method display dari interface displayData
	fmt.Println("Phone Information:")
	fmt.Println(myPhone.display())

	// soal nomor 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
	// 	buatlah function yang tipe data return-nya adalah interface kosong, dengan kondisi:
	// jika parameter kedua bernilai true maka tampilkan kalimat (asumsi sisinya 2)"luas persegi dengan sisi 2 cm adalah 4 cm"
	// jika parameter kedua bernilai false maka tampilkan hanya hasil angka saja (misal 4)
	// jika parameter pertama 0 dan parameter kedua bernilai true tampilkan "Maaf anda belum menginput sisi dari persegi"
	// jika parameter pertama 0 dan parameter kedua bernilai false maka tampilkan nil

	// jawaban nomor 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// Menggunakan type assertion untuk mengakses nilai dari variabel interface{}
	prefixStr := prefix.(string)
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	// Menjumlahkan angka-angka dari kedua slice
	total := 0
	for _, angka := range angkaPertama {
		total += angka
	}
	for _, angka := range angkaKedua {
		total += angka
	}
	// Menampilkan output sesuai format yang diminta
	fmt.Printf("%s%d + %d + %d + %d = %d\n", prefixStr, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)

}
