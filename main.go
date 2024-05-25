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
2. cara masuk ke menu panitia adalah dengan mengetikkan "panitia" pada menu utama
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
}

type PemilihTetap struct {
	nama    string
	pilihan string
}

type DaftarCalonAnggotaParlemen [NMAX]AnggotaParlemen
type DaftarPemilih [NMAX]PemilihTetap

func main() {
	var parlemen DaftarCalonAnggotaParlemen
	var pemilih DaftarPemilih
	var nSize_parlemen, nSize_pemilih int
	var input_1, input_2 int
	var jalan bool = true
	var waktu_pemilihan bool = false
	var jalan2 bool = true
	for jalan {
		cetak_main_menu()
		input_1 = process_input_main_menu()
		if input_1 == 1 {
			memilih_anggota_parlemen(&parlemen, nSize_parlemen, &pemilih, &nSize_pemilih, waktu_pemilihan)
		} else if input_1 == 2 {
			menu_menampilkan_data_parlemen()
			input_2 = process_input_menampilkan_data_parlemen()
			if input_2 == 1 {
				//fmt.Println(parlemen, nSize_parlemen)
				menampilkan_data_terurut_berdasarkan_nama(&parlemen, nSize_parlemen)
			} else if input_2 == 2 {
				menampilkan_data_terurut_berdasarkan_partai(&parlemen, nSize_parlemen)
			} else if input_2 == 3 {
				menampilkan_data_terurut_berdasarkan_suara(&parlemen, nSize_parlemen)
			}
		} else if input_1 == 3 {
			menu_pencarian_data_parlemen()
			input_2 = process_input_pencarian_data_parlemen()
			if input_2 == 1 {
				pencarian_berdasarkan_nama(&parlemen, nSize_parlemen)
			} else if input_2 == 2 {
				pencarian_berdasarkan_partai(&parlemen, nSize_parlemen)
			} else if input_2 == 3 {
				pencarian_berdasarkan_pemilih(&parlemen, pemilih, nSize_parlemen, nSize_pemilih)
			}

		} else if input_1 == 5 {
			jalan = false
		} else if input_1 == 6 {
			jalan2 = true
			for jalan2 {
				cetak_menu_panitia()
				input_2 = process_input_panitia()
				if input_2 == 1 {
					menambah_anggota_parlemen(&parlemen, &nSize_parlemen)
					//fmt.Println(parlemen, nSize_parlemen)
				} else if input_2 == 2 {
					menghapus_anggota_parlemen(&parlemen, &nSize_parlemen)
				} else if input_2 == 3 {
					mengedit_anggota_parlemen(&parlemen, nSize_parlemen)
				} else if input_2 == 4 {
					menghapus_data_suara_parlemen(&parlemen, nSize_parlemen, &pemilih, &nSize_pemilih)
				} else if input_2 == 6 {
					jalan2 = false
				} else if input_2 == 5 {
					waktu_pemilihan = cek_rentang_waktu()
				}
			}
		} else if input_1 == 4 {
			kalkulasi_threshold_kandidat(&parlemen, nSize_parlemen, nSize_pemilih)
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

func cetak_main_menu() {
	fmt.Println("Selamat datang di aplikasi pemilihan umum")
	print_waktu()
	fmt.Println("Silahkan masukkan opsi yang ingin anda pilih")
	fmt.Println("1. Memilih anggota parlemen")
	fmt.Println("2. List calon anggota")
	fmt.Println("3. Pencarian data calon anggota")
	fmt.Println("4. menampilkan pemenang")
	fmt.Println("5. Exit")
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
	fmt.Println("1. Pencarian berdasarkan nama")
	fmt.Println("2. Pencarian berdasarkan partai")
	fmt.Println("3. Pencarian berdasarkan pemilih")
}

func menu_menampilkan_data_parlemen() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Silahkan masukkan pencarian yang ingin anda gunakan")
	fmt.Println("1. menampilkan nama berdasarkan nama")
	fmt.Println("2. menampilkan nama berdasarkan partai")
	fmt.Println("3. menampilkan nama berdasarkan suara")
}

func process_input_menampilkan_data_parlemen() int {
	var x int
	var jalan bool = true
	for jalan {
		fmt.Scan(&x)
		if x >= 1 && x <= 3 {
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
		if x >= 1 && x <= 3 {
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
	var x string
	var jalan bool = true
	for jalan {
		fmt.Scan(&x)
		if x == "panitia" {
			return 6
		} else if x == "1" {
			return 1
		} else if x == "2" {
			return 2
		} else if x == "3" {
			return 3
		} else if x == "4" {
			return 4
		} else if x == "5" {
			return 5
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
	var nama1, nama2 string
	var jalan, ada bool
	ada = false
	jalan = waktu
	if !jalan {
		fmt.Println("Anda tidak sedang berada di masa pemilihan, Silahkan kembali saat masa pemilihan sedang berlangsung.")
	}
	for jalan {
		ada = false
		for !ada {
			fmt.Println("Silahkan masukkan pilihan anda dengan format 'nama_pemilih kandidat_yang_dipilih'. jika sudah selesai memilih, silahkan ketik -1 -1")
			fmt.Scan(&nama1, &nama2)
			if nama2 == "-1" || nama1 == "-1" {
				ada = true
				jalan = false
			}
			for i := 0; i < size; i++ {
				if Data_Parlemen[i].nama == nama2 {
					ada = true
					Data_Parlemen[i].suara++
					menambah_data_pemilih(data_pemilih, nama1, nama2, size_pemilih)
				}
			}
			if !ada {
				fmt.Println("Kandidat yang anda pilih tidak ada, silahkan pilih kembali")
			}
		}
	}
}

func menambah_data_pemilih(data_pemilih *DaftarPemilih, nama, pilihan string, size_pemilih *int) {
	var size = len(*data_pemilih)
	for i := 0; i < size; i++ {
		if data_pemilih[i].nama == "" {
			data_pemilih[i].nama = nama
			data_pemilih[i].pilihan = pilihan
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

	fmt.Println("Silahkan masukkan jumlah anggota parlemen yang ingin ditambahkan terlebih dahulu")
	fmt.Scan(&n)
	fmt.Println("Silahkan masukkan data anggota parlemen dengan format 'nama nama_partai'")

	for i := *size; i < *size+n; i++ {
		fmt.Scan(&Data_Parlemen[i].nama, &Data_Parlemen[i].partai)
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
		fmt.Println("Silahkan masukkan perubahan yang ingin anda ubah denan format 'nama nama_partai")
		fmt.Scan(&Data_Parlemen[idx].nama, &Data_Parlemen[idx].partai)
	} else {
		fmt.Println("Nama anggota yang anda cari tidak ditemukan!")
	}
}

func menghapus_anggota_parlemen(Data_Parlemen *DaftarCalonAnggotaParlemen, size *int) {
	var nama string
	var ada bool = false
	var idx int = -1
	fmt.Println("Silahkan masukkan nama anggota parlemen yang ingin anda hapus")
	fmt.Scan(&nama)
	for i := 0; i < *size; i++ {
		if Data_Parlemen[i].nama == nama {
			ada = true
			idx = i
		}
	}

	if ada {
		for i := idx; i < *size-1; i++ {
			Data_Parlemen[i] = Data_Parlemen[i+1]
		}
		*size--
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
	fmt.Println("NAMA |  PARTAI  |  SUARA")

	for k := 0; k < size; k++ {
		fmt.Println(Data_Parlemen[k].nama, Data_Parlemen[k].partai, Data_Parlemen[k].suara)
	}
}

func menampilkan_data_terurut_berdasarkan_nama(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	sort_nama(Data_Parlemen, size)
	//fmt.Println(*Data_Parlemen, size)
	fmt.Println("NAMA |  PARTAI  |  SUARA")
	for k := 0; k < size; k++ {
		fmt.Println(Data_Parlemen[k].nama, Data_Parlemen[k].partai, Data_Parlemen[k].suara)
	}
}

func menampilkan_data_terurut_berdasarkan_partai(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	sort_partai(Data_Parlemen, size)
	fmt.Println("NAMA |  PARTAI  |  SUARA")
	for k := 0; k < size; k++ {
		fmt.Println(Data_Parlemen[k].nama, Data_Parlemen[k].partai, Data_Parlemen[k].suara)
	}
}

func pencarian_berdasarkan_nama(Data_Parlemen *DaftarCalonAnggotaParlemen, size int) {
	var nama_pencarian string
	var ada bool = false
	var idx int = -1
	fmt.Println("Silahkan masukkan nama yang ingin anda cari")
	fmt.Scan(&nama_pencarian)
	for i := 0; i < size; i++ {
		if Data_Parlemen[i].nama == nama_pencarian {
			ada = true
			idx = i
		}
	}
	if ada {
		fmt.Println("Berikut data yang anda cari:")
		fmt.Println(Data_Parlemen[idx].nama, Data_Parlemen[idx].partai, Data_Parlemen[idx].suara)
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
			fmt.Println(Data_Parlemen[i].nama, Data_Parlemen[i].partai, Data_Parlemen[i].suara)
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
		}
	}
}

func kalkulasi_threshold_kandidat(Data_Parlemen *DaftarCalonAnggotaParlemen, size_parlemen, size_pemilih int) {
	var rerata float64
	sort_nama(Data_Parlemen, size_parlemen)
	rerata = float64(size_pemilih) / float64(size_parlemen)
	fmt.Println("Pemilih yang lolos untuk sebagai anggota legislatif adalah:")
	for i := 0; i < size_parlemen; i++ {
		if rerata < float64(Data_Parlemen[i].suara) {
			fmt.Printf("%s dari partai %s dengan perolehan suara %d\n", Data_Parlemen[i].nama, Data_Parlemen[i].partai, Data_Parlemen[i].suara)
		}
	}
}

func menghapus_data_suara_parlemen(Data_Parlemen *DaftarCalonAnggotaParlemen, size_parlemen int, data_pemilih *DaftarPemilih, size_pemilih *int) {
	var nama string
	var ada bool = false
	var idx int = -1
	var nama_yang_dipilih string
	fmt.Println("Silahkan masukkan nama pemilih yang ingin anda hapus")
	fmt.Scan(&nama)

	for i := 0; i < *size_pemilih; i++ {
		if data_pemilih[i].nama == nama {
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

	}
}

//udah selesai, tinggal benerin beberapa bug di func main
