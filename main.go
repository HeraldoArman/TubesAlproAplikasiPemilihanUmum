/*
TubesAlproAplikasiPemilihanUmum
Program ini ditujukan untuk menyelesaikan tugas besar mata kuliah algrotima pemrograman semester 2
Program yang dibuat adalah Aplikasi Pemilihan umum dengan spesifikasi seperti berikut

Deskripsi
Aplikasi digunakan untuk melakukan pemilihan umum calon legislatif dan partai tertentu. Pengguna aplikasi ini adalah pemilih dan juga petugas kpu.
Spesifikasi:
- Pengguna bisa menambahkan, mengubah (edit), dan juga menghapus data calon dan pemilih.
- Pemilih bisa melakukan pemilihan pada durasi waktu yang ditentukan saja, di luar itu, hanya bisa melihat daftar calon saja.
- Pengguna bisa menampilkan data terurut berdasarkan hasil perolehan suara. Berdasarkan partai, berdasarkan nama calon dan partai.
- Tentukan nilai threshold atau ambang batas suatu calon untuk bisa terpilih.
- Pengguna bisa melakukan pencarian data calon yang berasal dari partai tertentu, pencarian berdasarkan nama calon, dan juga nama pemilih dari calon tertentu.

Authors
- Heraldo Arman
- Riyan Permana Purba

CATATAN PENGGUNAAN APLIKASI
1. Sebelum memilih, pastikan panitia mengatur rentang waktu dan kandidat terlebih dahulu
2. Password untuk mengakses halaman panitia adalah "password"
3. Jika terdapat bug, tolong segera melapor ke author
*/
package main

import (
	"fmt"
	"time"
)

const NMAX int = 512

type AnggotaParlemen struct {
	nama   string
	partai string
	suara  int
	id string
}

type PemilihTetap struct {
	nama    string
	pilihan string
	id 		string
}

type DaftarCalonAnggotaParlemen [NMAX]AnggotaParlemen
type DaftarPemilih [NMAX]PemilihTetap

/*
	TODO Testing lebih lanjut
*/

func main() {
	var parlemen DaftarCalonAnggotaParlemen
	var pemilih DaftarPemilih
	var nSize_parlemen, nSize_pemilih int
	var input_1, input_2, input_3 int
	var jalan bool = true
	var waktu_pemilihan bool = false
	var jalan2 bool = true

	var password string = "password"
	var input_password string
	var password_benar bool = false

	fmt.Print("\033[H\033[2J")
	for jalan{
		cetak_menu_login()
		if password_benar{
			fmt.Print("\033[H\033[2J")
		}
		password_benar = false
		input_1 = process_input_login()
		if input_1 == 1{
			jalan2 = true
			fmt.Print("\033[H\033[2J")
			for jalan2{
				cetak_main_menu()
				input_2 = process_input_main_menu()
				if input_2 == 1 {
					fmt.Print("\033[H\033[2J")
					memilih_anggota_parlemen(&parlemen, nSize_parlemen, &pemilih, &nSize_pemilih, waktu_pemilihan)
				} else if input_2 == 2 {
					fmt.Print("\033[H\033[2J")
					menu_menampilkan_data_parlemen()
					input_3 = process_input_menampilkan_data_parlemen()
					if input_3 == 1 {
						fmt.Print("\033[H\033[2J")
						menampilkan_data_terurut_berdasarkan_nama(&parlemen, nSize_parlemen)
					} else if input_3 == 2 {
						fmt.Print("\033[H\033[2J")
						menampilkan_data_terurut_berdasarkan_partai(&parlemen, nSize_parlemen)
					} else if input_3 == 3 {
						fmt.Print("\033[H\033[2J")
						menampilkan_data_terurut_berdasarkan_suara(&parlemen, nSize_parlemen)
					} else if input_3 == 4{
						fmt.Print("\033[H\033[2J")
						menampilkan_data_terurut_berdasarkan_id(&parlemen, nSize_parlemen)
					}
				} else if input_2 == 3 {
					fmt.Print("\033[H\033[2J")
					menu_pencarian_data_parlemen()
					input_3 = process_input_pencarian_data_parlemen()
					if input_3 == 1 {
						fmt.Print("\033[H\033[2J")
						pencarian_berdasarkan_nama(&parlemen, nSize_parlemen)
					} else if input_3 == 2 {
						fmt.Print("\033[H\033[2J")
						pencarian_berdasarkan_partai(&parlemen, nSize_parlemen)
					} else if input_3 == 3 {
						fmt.Print("\033[H\033[2J")
						pencarian_berdasarkan_pemilih(&parlemen, pemilih, nSize_parlemen, nSize_pemilih)
					} else if input_3 == 4{
						fmt.Print("\033[H\033[2J")
						pencarian_berdasarkan_id(&parlemen, nSize_parlemen)
					} else if input_3 == 5{
						fmt.Print("\033[H\033[2J")
						pencarian_berdasarkan_IDpemilih(&parlemen, pemilih, nSize_parlemen, nSize_pemilih)
					}
				} else if input_2 == 4 {
					fmt.Print("\033[H\033[2J")
					kalkulasi_threshold_kandidat(&parlemen, nSize_parlemen, nSize_pemilih)
				} else if input_2 == 5{
					jalan2 = false
					fmt.Print("\033[H\033[2J")
				}
			}
		} else if input_1 == 2{
			fmt.Println("Masukkan password panitia:")
			fmt.Scan(&input_password)
			if input_password != password{
				password_benar = false
				fmt.Print("\033[H\033[2J")
				fmt.Println("Password yang anda masukkan salah!")
			} else{
				password_benar = true
				jalan2 = true
				for jalan2 {
					fmt.Print("\033[H\033[2J")
					cetak_menu_panitia()
					input_2 = process_input_panitia()
					if input_2 == 1 {
						fmt.Print("\033[H\033[2J")
						menambah_anggota_parlemen(&parlemen, &nSize_parlemen)
						//fmt.Println(parlemen, nSize_parlemen)
					} else if input_2 == 2 {
						fmt.Print("\033[H\033[2J")
						menghapus_anggota_parlemen(&parlemen, &nSize_parlemen)
					} else if input_2 == 3 {
						fmt.Print("\033[H\033[2J")
						mengedit_anggota_parlemen(&parlemen, nSize_parlemen)
					} else if input_2 == 4 {
						fmt.Print("\033[H\033[2J")
						menghapus_data_suara_parlemen(&parlemen, nSize_parlemen, &pemilih, &nSize_pemilih)
					} else if input_2 == 6 {
						password_benar = false
						jalan2 = false
						fmt.Print("\033[H\033[2J")
					} else if input_2 == 5 {
						fmt.Print("\033[H\033[2J")
						waktu_pemilihan = cek_rentang_waktu()
					}
				}
			}
		} else if input_1 == 3{
			fmt.Print("\033[H\033[2J")
			jalan = false
		}
	}
}

func print_waktu() {
	// Mendapatkan tanggal dan waktu saat ini
	now := time.Now()

	// Mendapatkan tanggal
	tahun, bulan, hari := now.Date()
	fmt.Printf("Tanggal hari ini: %02d/%02d/%d\n", hari, bulan, tahun)
}

func input_rentang_waktu() string {
	var waktu_string string
	//menginput waktu dalam bentuk string
	fmt.Scan(&waktu_string)
	return waktu_string
}

func cek_rentang_waktu() bool {
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

func cetak_menu_login(){
	fmt.Println("Selamat datang di aplikasi pemilihan umum")
	print_waktu()
	fmt.Println("Silahkan masukkan opsi yang ingin anda pilih")
	fmt.Println("1. login sebagai pemilih")
	fmt.Println("2. login sebagai panitia")
	fmt.Println("3. exit")
}

func cetak_main_menu() {
	fmt.Println("Selamat datang di aplikasi pemilihan umum")
	print_waktu()
	fmt.Println("Silahkan masukkan opsi yang ingin anda pilih")
	fmt.Println("1. Memilih anggota parlemen")
	fmt.Println("2. List calon anggota")
	fmt.Println("3. Pencarian data calon anggota")
	fmt.Println("4. menampilkan pemenang")
	fmt.Println("5. Kembali ke menu awal")
}

func cetak_menu_panitia() {
	fmt.Println("Selamat datang di menu panitia")
	print_waktu()
	fmt.Println("Silahkan masukkan opsi yang ingin anda pilih")
	fmt.Println("1. Menginput anggota parlemen baru")
	fmt.Println("2. Menghapus anggota parlemen")
	fmt.Println("3. Mengedit anggota parlemen")
	fmt.Println("4. Menghapus data suara parlemen")
	fmt.Println("5. Mengubah rentang waktu pemilihan")
	fmt.Println("6. Kembali ke menu awal")
}

func menu_pencarian_data_parlemen() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Silahkan masukkan pencarian yang ingin anda gunakan")
	fmt.Println("1. Pencarian berdasarkan nama kandidat")
	fmt.Println("2. Pencarian berdasarkan partai")
	fmt.Println("3. Pencarian berdasarkan nama pemilih")
	fmt.Println("4. Pencarian berdasarkan id kandidat")
	fmt.Println("5. Pencarian berdasarkan id pemilih")


}

func menu_menampilkan_data_parlemen() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Silahkan masukkan pencarian yang ingin anda gunakan")
	fmt.Println("1. menampilkan nama berdasarkan nama")
	fmt.Println("2. menampilkan nama berdasarkan partai")
	fmt.Println("3. menampilkan nama berdasarkan suara")
	fmt.Println("4. menampilkan nama berdasarkan ID")

}

func process_input_login() int{
	var x int
	var jalan bool = true
	for jalan{
		fmt.Scan(&x)
		if x >= 1 && x <= 3{
			return x
		} else {
			fmt.Print("\033[H\033[2J")
			cetak_menu_login()
			fmt.Println("Input tidak benar, silahkan masukkan kembali input yang benar")
		}
	}
	return -1
}

func process_input_menampilkan_data_parlemen() int {
	var x int
	var jalan bool = true
	for jalan {
		fmt.Scan(&x)
		if x >= 1 && x <= 4 {
			return x
		} else {
			fmt.Print("\033[H\033[2J")
			menu_menampilkan_data_parlemen()
			fmt.Println("Input tidak benar, silahkan masukkan kembali input yang benar")
		}
	}
	return -1
}

func process_input_pencarian_data_parlemen() int {
	var x int
	var jalan bool = true
	for jalan {
		fmt.Scan(&x)
		if x >= 1 && x <= 5 {
			return x
		} else {
			fmt.Print("\033[H\033[2J")
			menu_pencarian_data_parlemen()
			fmt.Println("Input tidak benar, silahkan masukkan kembali input yang benar")
		}
	}
	return -1
}

func process_input_main_menu() int {
	var x int
	var jalan bool = true
	for jalan {
		fmt.Scan(&x)

		if x >= 1 && x <= 5 {
			return x
		} else {
			fmt.Print("\033[H\033[2J")
			cetak_main_menu()
			fmt.Println("Input tidak benar, silahkan masukkan kembali input yang benar")
		}
	}
	return -1
}

func process_input_panitia() int {
	var x int
	var jalan bool = true
	for jalan {
		fmt.Scan(&x)
		if x <= 6 && x >= 1 {
			return x
		} else {
			fmt.Print("\033[H\033[2J")
			cetak_menu_panitia()
			fmt.Println("Input tidak benar, silahkan masukkan kembali input yang benar")
		}
	}
	return -1
}

func memilih_anggota_parlemen(Data_Parlemen *DaftarCalonAnggotaParlemen, size int, data_pemilih *DaftarPemilih, size_pemilih *int, waktu bool) {
	/* I.S. terdefinisi:
		        - daftar calon anggota parlemen (Data_Parlemen) yang berisi size calon anggota dengan nama dan jumlah suara masing-masing calon.
						- daftar pemilih (data_pemilih) yang berisi data pemilih yang sudah memilih, serta ukuran daftar pemilih(size_pemilih).
						- waktu yang menunjukkan apakah pemilihan sedang berlangsung atau tidak (true untuk sedang berlangsung, false untuk tidak sedang
	            berlangsung).
	     F.S. Data_Parlemen diperbarui dengan penambahan suara untuk kandidat yang dipilih oleh pemilih.
						data_pemilih diperbarui dengan menambahkan pemilih yang sudah memilih beserta kandidat yang dipilih.
						size_pemilih diperbarui sesuai dengan jumlah pemilih yang valid setelah pemilihan. */
	var nama1, nama2, id_pemilih string
	var jalan, ada bool
	ada = false
	jalan = waktu
	if !jalan {
		fmt.Println("Anda tidak sedang berada di masa pemilihan, Silahkan kembali saat masa pemilihan sedang berlangsung.")
	}
	for jalan {
		ada = false
		for !ada {
			fmt.Println("Silahkan masukkan pilihan anda dengan format 'nama_pemilih nama_kandidat_yang_dipilih id_pemilih' atau 'nama_pemilih id_kandidat_yang_dipilih id_pemilih'. jika sudah selesai memilih, silahkan ketik -1 -1")
			fmt.Scan(&nama1, &nama2, &id_pemilih)
			if nama2 == "-1" || nama1 == "-1" {
				ada = true
				jalan = false
			}
			for i := 0; i < size; i++ {
				if Data_Parlemen[i].nama == nama2 || Data_Parlemen[i].id == nama2{
					ada = true
					Data_Parlemen[i].suara++
					menambah_data_pemilih(data_pemilih, nama1, Data_Parlemen[i].nama, size_pemilih, id_pemilih)
				}
			}
			if !ada {
				fmt.Println("Kandidat yang anda pilih tidak ada, silahkan pilih kembali")
			}
		}
	}
}

func menambah_data_pemilih(data_pemilih *DaftarPemilih, nama, pilihan string, size_pemilih *int, id_pemilih string) {
	var size = len(*data_pemilih)
	for i := 0; i < size; i++ {
		if data_pemilih[i].nama == "" {
			data_pemilih[i].nama = nama
			data_pemilih[i].pilihan = pilihan
			data_pemilih[i].id = id_pemilih
			*size_pemilih++
			i = size + 10
		}
	}
}

func menambah_anggota_parlemen(Data_Parlemen *DaftarCalonAnggotaParlemen, size *int) {
	/* I.S.
	F.S.
	*/
	var n int
	var data_kosong DaftarCalonAnggotaParlemen
	fmt.Println("Silahkan masukkan jumlah anggota parlemen yang ingin ditambahkan terlebih dahulu")
	fmt.Scan(&n)
	fmt.Println("Silahkan masukkan data anggota parlemen dengan format 'nama nama_partai id'")
	fmt.Println("Apabila terdapat ID yang sama, maka masukan dengan ID yang terakhir akan otomatis tidak dimasukkan kedalam database")

	for i := *size; i < *size+n; i++ {
		fmt.Scan(&Data_Parlemen[i].nama, &Data_Parlemen[i].partai, &Data_Parlemen[i].id)
		for j := 0; j < i; j++{
			if Data_Parlemen[i].id == Data_Parlemen[j].id{
				Data_Parlemen[i] = data_kosong[1]
				*size--
			}
		}
	}
	*size += n
}

func mengedit_anggota_parlemen(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	var ada bool = false
	var idx int = -1
	var nama string
	fmt.Println("Silahkan ketikkan nama anggota parlemen yang ingin anda edit")
	fmt.Scan(&nama)
	for i := 0; i < size; i++ {
		if Data_Parlemen[i].nama == nama {
			ada = true
			idx = i
		}
	}
	if ada {
		fmt.Println("Silahkan masukkan perubahan yang ingin anda ubah denan format 'nama nama_partai id")
		fmt.Scan(&Data_Parlemen[idx].nama, &Data_Parlemen[idx].partai, &Data_Parlemen[idx].id)
	} else {
		fmt.Println("Nama anggota yang anda cari tidak ditemukan!")
	}
}

func menghapus_anggota_parlemen(Data_Parlemen *DaftarCalonAnggotaParlemen, size *int) {
	var nama, id string
	var ada bool = false
	var idx int = -1
	fmt.Println("Silahkan masukkan nama anggota parlemen yang ingin anda hapus dengan format 'nama id'")
	fmt.Scan(&nama, &id)
	for i := 0; i < *size; i++ {
		if Data_Parlemen[i].nama == nama && Data_Parlemen[i].id == id {
			ada = true
			idx = i
		}
	}

	if ada {
		for i := idx; i < *size-1; i++ {
			Data_Parlemen[i] = Data_Parlemen[i+1]
		}
		*size--
	} else{
		fmt.Println("Nama atau id yang anda masukkan tidak dapat ditemukan")
	}
}

func swap_string(s1, s2 *string) {
	temp1 := *s1
	*s1 = *s2
	*s2 = temp1
}

func swap_int(s1, s2 *int) {
	temp2 := *s1
	*s1 = *s2
	*s2 = temp2
}

func menampilkan_data_terurut_berdasarkan_suara(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	sort_suara(Data_Parlemen, size)
	fmt.Printf("%-10s %-18s %-10s %-10s\n", "NAMA", "PARTAI", "SUARA", "ID")

	for k := 0; k < size; k++ {
		fmt.Printf("%-10s %-18s %-10d %-10s\n", Data_Parlemen[k].nama, Data_Parlemen[k].partai, Data_Parlemen[k].suara, Data_Parlemen[k].id)
	}
}

func menampilkan_data_terurut_berdasarkan_nama(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	sort_nama(Data_Parlemen, size)
	//fmt.Println(*Data_Parlemen, size)
	fmt.Printf("%-10s %-18s %-10s %-10s\n", "NAMA", "PARTAI", "SUARA", "ID")
	for k := 0; k < size; k++ {
		fmt.Printf("%-10s %-18s %-10d %-10s\n", Data_Parlemen[k].nama, Data_Parlemen[k].partai, Data_Parlemen[k].suara, Data_Parlemen[k].id)
	}
}

func menampilkan_data_terurut_berdasarkan_partai(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	sort_partai(Data_Parlemen, size)
	fmt.Printf("%-10s %-18s %-10s %-10s\n", "NAMA", "PARTAI", "SUARA", "ID")
	for k := 0; k < size; k++ {
		fmt.Printf("%-10s %-18s %-10d %-10s\n", Data_Parlemen[k].nama, Data_Parlemen[k].partai, Data_Parlemen[k].suara, Data_Parlemen[k].id)
	}
}

func menampilkan_data_terurut_berdasarkan_id(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	sort_id(Data_Parlemen, size)
	fmt.Printf("%-10s %-18s %-10s %-10s\n", "NAMA", "PARTAI", "SUARA", "ID")
	for k := 0; k < size; k++ {
		fmt.Printf("%-10s %-18s %-10d %-10s\n", Data_Parlemen[k].nama, Data_Parlemen[k].partai, Data_Parlemen[k].suara, Data_Parlemen[k].id)
	}
}

func pencarian_berdasarkan_nama(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	var nama_pencarian string
	//var idx int = -1
	var count int
	var idx_arr [NMAX]int
	fmt.Println("Silahkan masukkan nama yang ingin anda cari")
	fmt.Scan(&nama_pencarian)
	sort_nama(Data_Parlemen, size)
	for i := 0; i < size; i++ {
		if Data_Parlemen[i].nama == nama_pencarian {
			count++
			for j := 0; j < count; j++{
				idx_arr[j] = i
			}
		}
	}

	if count > 0 {
		fmt.Println("Berikut data yang anda cari:")
		for i := 0; i < count; i++{
			fmt.Println(Data_Parlemen[idx_arr[i]].nama, Data_Parlemen[idx_arr[i]].partai, Data_Parlemen[idx_arr[i]].suara, Data_Parlemen[idx_arr[i]].id)
		}
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak dapat ditemukan")
	}
}

func pencarian_berdasarkan_id(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	var id_pencarian string
	var ada bool = false
	var idx int = -1
	var kiri, kanan, tengah int
	fmt.Println("Silahkan masukkan id yang ingin anda cari")
	fmt.Scan(&id_pencarian)

	sort_id(Data_Parlemen, size)
	kiri = 0
	kanan = size - 1
	for kanan >= kiri{
		tengah = (kanan + kiri) / 2
		if Data_Parlemen[tengah].id == id_pencarian{
			idx = tengah
			ada = true
		} else if Data_Parlemen[tengah].id < id_pencarian{
			kiri = tengah + 1
		} else{
			kanan = tengah - 1
		}
	}
	if ada {
		fmt.Println("Berikut data yang anda cari:")
		fmt.Println(Data_Parlemen[idx].nama, Data_Parlemen[idx].partai, Data_Parlemen[idx].suara, Data_Parlemen[idx].id)
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak dapat ditemukan")
	}
}


func pencarian_berdasarkan_partai(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	var nama_pencarian string
	var ada bool = false
	sort_partai(Data_Parlemen, size)
	fmt.Println("Silahkan masukkan partai yang ingin anda cari")
	fmt.Scan(&nama_pencarian)
	for i := 0; i < size; i++ {
		if nama_pencarian == Data_Parlemen[i].partai {
			ada = true
			fmt.Printf("%-10s %-18s %-10d %-10s\n", Data_Parlemen[i].nama, Data_Parlemen[i].partai, Data_Parlemen[i].suara, Data_Parlemen[i].id)
		}
	}
	if !ada {
		fmt.Println("Maaf, partai yang anda cari tidak ada")
	}

}

func pencarian_berdasarkan_pemilih(Data_Parlemen *DaftarCalonAnggotaParlemen, Data_pemilih DaftarPemilih, size_parlemen, size_pemilih int) {
	var nama_pencarian string
	var ada bool = false
	fmt.Scan("Silahkan masukkan nama pemilih yang ingin anda cari")
	fmt.Scan(&nama_pencarian)
	for i := 0; i < size_pemilih; i++ {
		if nama_pencarian == Data_pemilih[i].nama {
			ada = true
			for j := 0; j < size_parlemen; j++ {
				if Data_pemilih[i].pilihan == Data_Parlemen[j].nama {
					fmt.Println("Nama anggota yang dipilih:", Data_pemilih[i].pilihan)
					fmt.Println(Data_Parlemen[j].nama, Data_Parlemen[j].partai, Data_Parlemen[j].suara)
				}
			}
		}
	}
	if !ada {
		fmt.Println("Nama yang anda pilih tidak dapat ditemukan")
	}
}

func pencarian_berdasarkan_IDpemilih(Data_Parlemen *DaftarCalonAnggotaParlemen, Data_pemilih DaftarPemilih, size_parlemen, size_pemilih int) {
	var nama_pencarian string
	var ada bool = false
	fmt.Println("Silahkan masukkan nama pemilih yang ingin anda cari")
	fmt.Scan(&nama_pencarian)
	for i := 0; i < size_pemilih; i++ {
		if nama_pencarian == Data_pemilih[i].id {
			ada = true
			for j := 0; j < size_parlemen; j++ {
				if Data_pemilih[i].pilihan == Data_Parlemen[j].nama {
					fmt.Println("Nama anggota yang dipilih:", Data_pemilih[i].pilihan)
					fmt.Println(Data_Parlemen[j].nama, Data_Parlemen[j].partai, Data_Parlemen[j].suara)
				}
			}
		}
	}
	if !ada {
		fmt.Println("Nama yang anda pilih tidak dapat ditemukan")
	}
}

func sort_nama(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	var max_idx int
	for i := 0; i < size-1; i++ {
		max_idx = i
		for j := i + 1; j < size; j++ {
			if Data_Parlemen[j].nama < Data_Parlemen[max_idx].nama {
				max_idx = j
			}
		}
		if max_idx != i {
			swap_string(&Data_Parlemen[i].nama, &Data_Parlemen[max_idx].nama)
			swap_string(&Data_Parlemen[i].partai, &Data_Parlemen[max_idx].partai)
			swap_int(&Data_Parlemen[i].suara, &Data_Parlemen[max_idx].suara)
			swap_string(&Data_Parlemen[i].id, &Data_Parlemen[max_idx].id)
		}
	}
}

func sort_partai(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	var max_idx int
	for i := 0; i < size-1; i++ {
		max_idx = i
		for j := i + 1; j < size; j++ {
			if Data_Parlemen[j].partai < Data_Parlemen[max_idx].partai {
				max_idx = j
			}
		}
		if max_idx != i {
			swap_string(&Data_Parlemen[i].nama, &Data_Parlemen[max_idx].nama)
			swap_string(&Data_Parlemen[i].partai, &Data_Parlemen[max_idx].partai)
			swap_int(&Data_Parlemen[i].suara, &Data_Parlemen[max_idx].suara)
			swap_string(&Data_Parlemen[i].id, &Data_Parlemen[max_idx].id)

		}
	}
}

func sort_suara(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	var max_idx int
	for i := 0; i < size-1; i++ {
		max_idx = i
		for j := i + 1; j < size; j++ {
			if Data_Parlemen[j].suara > Data_Parlemen[max_idx].suara {
				max_idx = j
			}
		}
		if max_idx != i {
			swap_string(&Data_Parlemen[i].nama, &Data_Parlemen[max_idx].nama)
			swap_string(&Data_Parlemen[i].partai, &Data_Parlemen[max_idx].partai)
			swap_int(&Data_Parlemen[i].suara, &Data_Parlemen[max_idx].suara)
			swap_string(&Data_Parlemen[i].id, &Data_Parlemen[max_idx].id)

		}
	}
}

func sort_id(Data_Parlemen *DaftarCalonAnggotaParlemen, size int){
	var j int
	for i := 1; i < size; i++{
		j = i
		for j > 0 && Data_Parlemen[j-1].id > Data_Parlemen[j].id{
			swap_string(&Data_Parlemen[j-1].nama, &Data_Parlemen[j].nama)
			swap_string(&Data_Parlemen[j-1].partai, &Data_Parlemen[j].partai)
			swap_int(&Data_Parlemen[j-1].suara, &Data_Parlemen[j].suara)
			swap_string(&Data_Parlemen[j-1].id, &Data_Parlemen[j].id)
			j--
		}
	}
}

func kalkulasi_threshold_kandidat(Data_Parlemen *DaftarCalonAnggotaParlemen, size_parlemen, size_pemilih int) {
	var rerata float64
	sort_nama(Data_Parlemen, size_parlemen)
	rerata = float64(size_pemilih) / float64(size_parlemen)
	fmt.Println("Pemilih yang lolos untuk sebagai anggota legislatif adalah:")
	fmt.Printf("%-10s %-18s %-10s %-10s\n", "NAMA", "PARTAI", "SUARA", "ID")
	for i := 0; i < size_parlemen; i++ {
		if rerata < float64(Data_Parlemen[i].suara) {
			fmt.Printf("%-10s %-18s %-10d %-10s\n", Data_Parlemen[i].nama, Data_Parlemen[i].partai, Data_Parlemen[i].suara, Data_Parlemen[i].id)
		}
	}
}

func menghapus_data_suara_parlemen(Data_Parlemen *DaftarCalonAnggotaParlemen, size_parlemen int, data_pemilih *DaftarPemilih, size_pemilih *int) {
	var nama, id string
	var ada bool = false
	var idx int = -1
	var nama_yang_dipilih string
	fmt.Println("Silahkan masukkan nama dan id pemilih yang ingin anda hapus dengan format 'nama id'")
	fmt.Scan(&nama, &id)

	for i := 0; i < *size_pemilih; i++ {
		if data_pemilih[i].nama == nama && data_pemilih[i].id == id {
			ada = true
			idx = i
			nama_yang_dipilih = data_pemilih[i].pilihan
		}
	}
	if ada {
		for i := idx; i < *size_pemilih-1; i++ {
			data_pemilih[i] = data_pemilih[i+1]
		}
		*size_pemilih--
		for i := 0; i < size_parlemen; i++ {
			if Data_Parlemen[i].nama == nama_yang_dipilih {
				Data_Parlemen[i].suara--
			}
		}
	} else{
		fmt.Println("Data yang anda masukkan tidak dapat ditemukan")
	}
}