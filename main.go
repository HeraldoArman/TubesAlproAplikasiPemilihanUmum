package main

import (
	"fmt"
	"strings"
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

//masih belum selesai
func main() {
	var masukan_main_menu int

	fmt.Print("\033[H\033[2J")
	cetak_main_menu()
	masukan_main_menu = process_input_main_menu()
	

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
	print_waktu()
	fmt.Println("Silahkan masukkan opsi yang ingin anda pilih")
	fmt.Println("1. Memilih anggota parlemen")
	fmt.Println("2. List calon anggota")
	fmt.Println("3. Pencarian data calon anggota")
	fmt.Println("4. Exit")
}

func cetak_menu_panitia(){
	fmt.Println("Selamat datang di menu panitia")
	print_waktu()
	fmt.Println("Silahkan masukkan opsi yang ingin anda pilih")
	fmt.Println("1. Menginput anggota parlemen baru")
	fmt.Println("2. Menghapus anggota parlemen")
	fmt.Println("3. Mengedit anggota parlemen")
	fmt.Println("4. Menghapus data parlemen")
	fmt.Println("5. Mengubah rentang waktu pemilihan")
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



//fungsi ini masih belom selesai
func memilih_anggota_parlemen(Data_Parlemen *DaftarCalonAnggotaParlemen, size int, data_pemilih *DaftarPemilih){
	var input string
	var ada bool
	var jalan = true
	if cek_rentang_waktu(){
		fmt.Println("Anda tidak sedang berada di masa pemilihan, Silahkan kembali saat masa pemilihan sedang berlangsung.")
	} else{
		for jalan{
			ada = false
			for !ada{
				fmt.Println("Silahkan masukkan pilihan anda dengan format 'nama_pemilih (spasi) kandidat_yang_dipilih'")
				fmt.Scan(&input)
				if input == "-1"{
					jalan = false
					break
				}
				kata := strings.Split(input, " ")
				//nama_pemilih = kata[0]
				//nama_pilihan = kata[1]
				for i := 0; i < size; i++{
					if Data_Parlemen[i].nama == kata[1]{
						Data_Parlemen[i].suara++
						ada = true
						for j := 0; j < NMAX; j++{
							if data_pemilih[j].nama == ""{
								data_pemilih[j].nama = kata[0]
								data_pemilih[j].pilihan = kata[1]
								break
							}
						}
					}
				}
				if !ada{
					fmt.Println("Kandidat yang anda pilih tidak ada, silahkan pilih kembali")
				}
			}
		}
	}
}