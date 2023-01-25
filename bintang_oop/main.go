// Study Jam 3 Golang - Mini Project

package main

import "fmt"

type User struct {
	id       string
	username string
	email    string
	noTelp   string
}
type Order struct {
	id              string
	movieTitle      string
	bioskopLocation string
	watchDate       string
	seatType        string
	ticketPrice     int
	seatAmount      int
	discount        int
}
type Payment struct {
	payment string
}

func (u User) User() string {
	return "Pesan tiket film untuk user " + u.username + " dengan ID " + u.id
}
func (u User) NoTelp() string {
	return "(" + u.noTelp + ")"
}
func (o *Order) SeatAmount() int {
	if o.seatAmount <= 0 {
		o.discount = 0
	}
	return o.seatAmount
}
func (o Order) TotalPrice() int {
	return (o.ticketPrice * o.seatAmount) - (o.discount * o.seatAmount)
}
func (p Payment) Payment() string {
	return "Jenis Pembayaran: " + p.payment
}

func main() {
	user := &User{
		id:       "MX4D7TR26",
		username: "Bintang",
		email:    "example@gmail.com",
		noTelp:   "0812-3456-7890",
	}
	order := &Order{
		id:              "72164371881318",
		movieTitle:      "Jhon Wick: Chapter 4",
		bioskopLocation: "The Park Solobaru",
		watchDate:       "2023/04/28",
		seatType:        "Reguler",
		ticketPrice:     40000,
		seatAmount:      1,
		discount:        5000,
	}
	payment := Payment{payment: "E-Wallet: DANA"}

	fmt.Println("================================================================")
	fmt.Println(user.User())
	fmt.Println("================================================================")
	fmt.Println("ID Transaksi	: " + order.id)
	fmt.Println("Judul Film	: " + order.movieTitle)
	fmt.Println("Lokasi Bioskop	: " + order.bioskopLocation)
	fmt.Println("Tanggal Nonton	: " + order.watchDate)
	fmt.Println("Jenis Kursi	: " + order.seatType)
	fmt.Println("Harga Tiket	: Rp", order.ticketPrice, ",-")

	if order.seatAmount <= 0 {
		fmt.Println("Jumlah Kursi	:", order.SeatAmount(), "(Silahkan pilih minimal 1 kursi!)")
	} else {
		fmt.Println("Jumlah Kursi	:", order.SeatAmount())
	}

	fmt.Println("================================================================")
	fmt.Println("Diskon		: Rp", order.discount, ",-")
	fmt.Println("Total Biaya	: Rp", order.TotalPrice(), ",-")
	fmt.Println("================================================================")
	fmt.Println(payment.Payment(), user.NoTelp())
	fmt.Println("================================================================")
}
