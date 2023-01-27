package main

import "fmt"

type Bank struct {
	nama  string
	saldo int
}

func (bank *Bank) SetSaldo(saldo int) {
	if saldo < 0 {
		fmt.Println("Saldo tidak boleh negatif")
		return
	}
	bank.saldo = saldo
}
func (bank Bank) Saldo() int {
	return bank.saldo
}

func main() {
	bank := &Bank{
		nama:  "Bank ABC",
		saldo: 50000,
	}
	fmt.Println("Saldo awal", bank.saldo)
	bank.SetSaldo(10000)
	fmt.Println("Saldo setelah diubah:", bank.Saldo())
}
