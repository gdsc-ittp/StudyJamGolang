// Example

package main

import "fmt"

// type Car struct {
// 	make  string
// 	model string
// 	year  int
// }

// func (c *Car) start() {
// 	fmt.Println("Engine started")
// }

// func (c *Car) stop()  {
// 	fmt.Println(("Engine stoped"))
// }

type Bank struct {
	nama string
	saldo int
}

func (b *Bank) setSaldo(saldo int)  {
	if saldo < 0 {
		fmt.Println("Saldo tidak boleh zero atau negatif")
		return
	}
	b.saldo = saldo
}

func (b Bank) Saldo() int {
	return b.saldo
}

func main() {
	// car := Car{
	// 	make: "Toyota",
	// 	model: "Carry",
	// 	year: 2019,
	// }

	// car.start()
	// car.stop()

	bank := &Bank{
		nama: "ABC",
		saldo: 0,
	}

	bank.setSaldo(15000)
	fmt.Println("Saldo :",bank.Saldo())
}