package main

import "fmt"

type InfoProduk interface {
	cetak()
}

type Produk struct {
	judul    string
	penulis  string
	penerbit string
	harga    int
	diskon   int
}

func (produk *Produk) setJudul(judul string) {
	if judul != "" {
		produk.judul = judul
	}
}

func (produk *Produk) setPenulisl(penulis string) {
	if penulis != "" {
		produk.penulis = penulis
	}
}

func (produk *Produk) setPenerbit(penerbit string) {
	if penerbit != "" {
		produk.penerbit = penerbit
	}
}

func (produk *Produk) setHarga(harga int) {
	if harga > 0 {
		produk.harga = harga
	}
}

func (produk *Produk) setDiskon(diskon int) {
	if diskon > 0 && diskon < 100 {
		produk.diskon = diskon
	} else {
		produk.diskon = 0
	}
}

func (produk Produk) getJudul() string {
	if produk.judul != "" {
		return produk.judul
	}
	return "Judul belum ditentukan"
}

func (produk Produk) getPenulis() string {
	if produk.penulis != "" {
		return produk.penulis
	}
	return "Penulis belum ditentukan"
}

func (produk Produk) getPenerbit() string {
	if produk.penerbit != "" {
		return produk.penerbit
	}
	return "Penerbit belum ditentukan"
}

func (produk Produk) getHarga() int {
	if produk.harga > 0 {
		return produk.harga - (produk.harga * produk.diskon / 100)
	}
	return 0
}

func (produk Produk) getDiskon() int {
	if produk.diskon >= 0 {
		return produk.diskon
	}
	return 0
}

func (produk Produk) cetak() {
	fmt.Println("\nJudul\t :", produk.getJudul())
	fmt.Println("Penulis\t :", produk.getPenulis())
	fmt.Println("Penerbit :", produk.getPenerbit())
	fmt.Println("Harga\t : Rp.", produk.getHarga())
	fmt.Println("Diskon\t :", produk.getDiskon(), "%")
}

func detailProduk(i InfoProduk) {
	i.cetak()
}

func main() {
	produk1 := Produk{
		judul:    "",
		penulis:  "",
		penerbit: "",
		harga:    0,
		diskon:   0,
	}

	produk1.setJudul("Mariposa")
	produk1.setPenulisl("Hidayatul Fajriyah")
	produk1.setPenerbit("Coconut Books")
	produk1.setHarga(55000)
	produk1.setDiskon(10)

	detailProduk(produk1)
}
