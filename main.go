package main

import (
	"fmt"
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
}
