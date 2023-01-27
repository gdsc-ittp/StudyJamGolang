package main

import (
	mhs "StudyJamGolang/bayupamuji_oop/data"
)

func main() {
	mhsITTP := &mhs.Mahasiswa{}

	mhsITTP.SetNim(12435454)
	mhsITTP.SetName("Afifah")
	mhsITTP.SetProdi("Sistem Informasi")
	mhsITTP.SetSemester(5)

	println("===== Data Mahasiswa =====")
	println("Nim:", mhsITTP.GetNim())
	println("Name:", mhsITTP.GetName())
	println("Prodi:", mhsITTP.GetProdi())
	println("Semester:", mhsITTP.GetSemester())
}