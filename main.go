package main

import (
	"fmt"
	"time"
)


const NMAX int = 512
type AnggotaParlemen struct {
	nama string
	partai string
	suara int
}

type PemilihTetap struct {
	nama string
	pilihan string
}

type DaftarCalonAnggotaParlemen [NMAX]AnggotaParlemen
type DaftarPemilih [NMAX]PemilihTetap


func main() {
	fmt.Print("\033[H\033[2J")
	//var pemilih DaftarPemilih
	//fmt.Println(pemilih[0].nama)
	//print_waktu()
	//fmt.Println(cek_rentang_waktu())
	//cetak_main_menu()
	//fmt.Println(process_input_main_menu())
}


func print_waktu(){
	// Mendapatkan tanggal dan waktu saat ini
	now := time.Now()

	// Mendapatkan tanggal
	tahun, bulan, hari := now.Date()
	fmt.Printf("Tanggal hari ini: %02d/%02d/%d\n", hari, bulan, tahun)
}

func input_rentang_waktu() string{
	var waktu_string string
	//menginput waktu dalam bentuk string
	fmt.Scan(&waktu_string)
	return waktu_string
}

func cek_rentang_waktu() bool{
	var rentang1, rentang2 string

	//mendapatkan input waktu
	fmt.Println("masukkan rentang waktu awal dengan format dd/mm/yyyy")
	rentang1 = input_rentang_waktu()
	//menggunakan module time untuk mengubah format penanggalan menjadi bertipe time.time
	cek_start, err := time.Parse("02/01/2006", rentang1)
	//mengecek error dalam input
	if err != nil {
        fmt.Println("Format tanggal awal tidak valid:", err)
    }


	fmt.Println("masukkan rentang waktu akhir dd/mm/yyyy")
	rentang2 = input_rentang_waktu()
	cek_akhir, err := time.Parse("02/01/2006", rentang2)
	if err != nil {
        fmt.Println("Format tanggal akhir tidak valid:", err)
    }


	//mengecek tanggal hari ini
	hari_ini := time.Now()
	//me return nilai true atau false sesuai dengan rentang waktu
	return hari_ini.After(cek_start) && hari_ini.Before(cek_akhir)
}

func cetak_main_menu(){
	fmt.Println("Selamat datang di aplikasi pemilihan umum")
	fmt.Println("Silahkan masukkan opsi yang ingin anda pilih")
	fmt.Println("1. Memilih anggota parlemen")
	fmt.Println("2. List calon anggota")
	fmt.Println("3. Pencarian data calon anggota")
	fmt.Println("4. Exit")
}



func process_input_main_menu() int{
	var x string
	var jalan bool = true
	for jalan{
		fmt.Scan(&x)
		if x == "panitia"{
			return 5
		} else if x == "1"{
			return 1
		} else if x == "2"{
			return 2
		} else if x == "3"{
			return 3
		} else if x == "4"{
			return 4
		} else{
			fmt.Print("\033[H\033[2J")
			cetak_main_menu()
			fmt.Println("Input tidak benar, silahkan masukkan kembali input yang benar")
		}
	}
	return 0
}

func memilih_anggota_parlemen(){
	
}