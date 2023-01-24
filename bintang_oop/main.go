// Study Jam 3 Golang - Mini Project

package main

import "fmt"

type BookingTicketMovie interface {
	processBooking() string
}
type Username struct {
	username string
}
type Title struct {
	movieTitle string
}
type Location struct {
	bioskopLocation string
}
type Date struct {
	watchDate string
}
type SeatType struct {
	seatType string
}
type TicketPrice struct {
	ticketPrice string
}
type SeatAmount struct {
	seatAmount string
}

func (u Username) processBooking() string {
	return "Pesan tiket film untuk user " + u.username
}
func (t Title) processBooking() string {
	return "Judul Film: " + t.movieTitle
}
func (l Location) processBooking() string {
	return "Lokasi Bioskop: " + l.bioskopLocation
}
func (d Date) processBooking() string {
	return "Tanggal Nonton: " + d.watchDate
}
func (s SeatType) processBooking() string {
	return "Jenis Kursi: " + s.seatType
}
func (t TicketPrice) processBooking() string {
	return "Harga Tiket: " + t.ticketPrice
}
func (s SeatAmount) processBooking() string {
	return "Jumlah Kursi: " + s.seatAmount
}

func makeBookingRequest(b BookingTicketMovie) {
	fmt.Println(b.processBooking())
}
func main() {
	username := Username{username: "Bintang"}
	title := Title{movieTitle: "Jhon Wick: Chapter 4"}
	location := Location{bioskopLocation: "The Park Solobaru"}
	date := Date{watchDate: "2023/04/28"}
	seatType := SeatType{seatType: "Regular"}
	ticketPrice := TicketPrice{ticketPrice: "IDR 40K"}
	seatAmount := SeatAmount{seatAmount: "4"}

	fmt.Println("================================================================")
	makeBookingRequest(username)
	fmt.Println("================================================================")
	makeBookingRequest(title)
	makeBookingRequest(location)
	makeBookingRequest(date)
	makeBookingRequest(seatType)
	makeBookingRequest(ticketPrice)
	makeBookingRequest(seatAmount)
	fmt.Println("================================================================")
}