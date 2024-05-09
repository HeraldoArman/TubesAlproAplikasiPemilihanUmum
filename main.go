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
	var pemilih DaftarPemilih
	fmt.Println(pemilih[0].nama)
	print_waktu()
	fmt.Println(cek_rentang_waktu())
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
	fmt.Scan(&waktu_string)
	return waktu_string
}

func cek_rentang_waktu() bool{
	//var rentang string
	//var d1, m1, y1, d2, m2, y2, total_waktu_sekarang, total_waktu_1, total_waktu_2 int
	//rentang = input_rentang_waktu()
//
	//now := time.Now()
	//tahun_sekarang, bulan_sekarang, hari_sekarang := now.Date()
	//total_waktu_sekarang = 
//
//
	//d1 = int(rentang[0]+rentang[1])
	//m1 = int(rentang[3]+rentang[4])
	//y1 = int(rentang[6]+rentang[7]+rentang[8]+rentang[9])
	//d2 = int(rentang[13]+rentang[14])
	//m2 = int(rentang[16]+rentang[17])
	//y2 = int(rentang[19]+rentang[20]+rentang[21]+rentang[22])
	var rentang1, rentang2 string

	fmt.Println("masukkan rentang waktu awal dengan format dd/mm/yyyy")
	rentang1 = input_rentang_waktu()
	cek_start, err := time.Parse("02/01/2006", rentang1)
	if err != nil {
        fmt.Println("Format tanggal awal tidak valid:", err)
    }


	fmt.Println("masukkan rentang waktu akhir dd/mm/yyyy")
	rentang2 = input_rentang_waktu()
	cek_akhir, err := time.Parse("02/01/2006", rentang2)
	if err != nil {
        fmt.Println("Format tanggal akhir tidak valid:", err)
    }

	hari_ini := time.Now()

	return hari_ini.After(cek_start) && hari_ini.Before(cek_akhir)
	
}