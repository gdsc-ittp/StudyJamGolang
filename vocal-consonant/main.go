// Nur Muhammad Himawan
// Study Jam Golang #3 - GDSC ITTP

package main

import (
	"bufio"
	"fmt"
	"os"
)

func num_vokal(words string) int {
	jumlah_vokal := 0

	for _, word := range words {
		if word == 'a' || word == 'i' || word == 'u' || word == 'e' || word == 'o' ||
			word == 'A' || word == 'I' || word == 'U' || word == 'E' || word == 'O' {
			jumlah_vokal++
		}
	}

	return jumlah_vokal
}

func num_konsonan(words string) int {
	jumlah_konsonan := 0

	for _, word := range words {
		if word != 'a' && word != 'i' && word != 'u' && word != 'e' && word != 'o' && word != ' ' &&
			word != 'A' && word != 'I' && word != 'U' && word != 'E' && word != 'O' {
			jumlah_konsonan++
		}
	}

	return jumlah_konsonan
}

func main() {
	word := ""
	vokal := 0
	konsonan := 0

	fmt.Print("Masukan kalimat : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	word = scanner.Text()

	//memanggil prosedur yang mengembalikan nilai int
	vokal = num_vokal(word)
	konsonan = num_konsonan(word)

	fmt.Println("Jumlah huruf vokal    : ", vokal)
	fmt.Println("Jumlah huruf konsonan : ", konsonan)
}
