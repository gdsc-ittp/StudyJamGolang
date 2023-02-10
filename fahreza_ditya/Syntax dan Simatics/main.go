package main

import "fmt"

var x int // variable global

func main() {
	fmt.Println("\nfmt.Println :")
	fmt.Println("Hello World! I'am studiying GoLang")

	fmt.Println("\nGlobal variable :")
	x = -5
	fmt.Println("Nilai varible x adalah", x)

	fmt.Println("\nLocal variable :")
	y := 10 // variable local
	fmt.Println("Nilai variable y adalah", y)

	fmt.Println("\nConstant variable :")
	const z = 15
	fmt.Println("Nilai constant z adalah", z)

	fmt.Println("\nBranching if else if :")
	if x > 0 {
		fmt.Println("Variable x bernilai positif")
	} else if x < 0 {
		fmt.Println("Variable x bernilai negatif")
	} else {
		fmt.Println("Variable x adalah zero")
	}

	fmt.Println("\nBranching switch :")
	switch y {
	case 1:
		fmt.Println("Nilai y adalah 1")
	case 2:
		fmt.Println("Nilai y adalah 2")
	case 3:
		fmt.Println("Nilai y adalah 3")
	case 4:
		fmt.Println("Nilai y adalah 4")
	case 5:
		fmt.Println("Nilai y adalah 5")
	case 6:
		fmt.Println("Nilai y adalah 6")
	case 7:
		fmt.Println("Nilai y adalah 7")
	default:
		fmt.Println("Nilai y lebih dari 7")
	}

	fmt.Println("\nLooping for with argument :")
	for i := 1; i <= y; i++ {
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("\nLooping for without argument :")
	a := 1
	for {
		fmt.Println("Perulangan ke", a)
		a++
		if a > 10 {
			break
		}
	}

	fmt.Println("\nLooping for range :")
	numbers := []int{1, 2, 3, 4, 5}
	for i, n := range numbers {
		fmt.Println("index", i, "bernilai", n)
	}
}
