package main

import (
	mhs "StudyJamGolang/bayupamuji_oop/data"
	"fmt"
)

func main() {
	// mhsITTP := &mhs.Mahasiswa{}

	// mhsITTP.SetNim(12435454)
	// mhsITTP.SetName("Afifah")
	// mhsITTP.SetProdi("Sistem Informasi")
	// mhsITTP.SetSemester(5)
	// mhsITTP.SetIpk(3.20)

	// maxSks := mhsITTP.MaxSks()

	// println("===== Data Mahasiswa =====")
	// println("Nim:", mhsITTP.GetNim())
	// println("Name:", mhsITTP.GetName())
	// println("Prodi:", mhsITTP.GetProdi())
	// println("Semester:", mhsITTP.GetSemester())
	// println("Maks sks pada semester ini: ", maxSks)

	mhsVeteran := &mhs.MahasiswaYudisium{}

	mhsVeteran.SetNim(12435455)
	mhsVeteran.SetName("Muhammad Hanif")
	mhsVeteran.SetProdi("Sistem Informasi")
	mhsVeteran.SetSemester(8)
	mhsVeteran.SetIpk(3.20)
	mhsVeteran.SetIsLulusSidang(true)

	name := mhsVeteran.GetName()
	nim := mhsVeteran.GetNim()
	prodi := mhsVeteran.GetProdi()
	semester := mhsVeteran.GetSemester()
	ipk := mhsVeteran.GetIpk()
	status := mhsVeteran.CekStatus()

	println("===== Mahasiswa Yudisium =====\n")
	
	println("Nim:", nim)
	println("Name:", name)
	println("Prodi:", prodi)
	println("Semester:", semester)
	fmt.Printf("Ipk: %.2f\n", ipk)
	println("status yudisium:", status)
	
	println("\n===== ITTP 2023 =====")

}