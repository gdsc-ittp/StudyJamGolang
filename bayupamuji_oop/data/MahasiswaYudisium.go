package data

type MahasiswaYudisium struct {
	isLulusSidang bool
	Mahasiswa
}

func (mhs *MahasiswaYudisium) SetSemester(semester int) {
	if semester < 7 {
		mhs.semester = 7
	} else {
		mhs.semester = semester
	}
}

func (mhs *MahasiswaYudisium) SetIsLulusSidang(isLulusSidang bool) {
	mhs.isLulusSidang = isLulusSidang
}

func (mhs MahasiswaYudisium) GetIsLulusSidang() bool {
	return mhs.isLulusSidang
}

func (mhsYds MahasiswaYudisium) CekStatus() string {

	var status string

	if mhsYds.isLulusSidang {
		status = "Bisa Yudisium"
	} else {
		status = "Tidak bisa Yudisium"
	}

	return status
}

func (mhsYds MahasiswaYudisium) MaxSks() int {
	var maxSks int
	if mhsYds.isLulusSidang {
		maxSks = 0
	} else {
		maxSks = 6
	}
	return maxSks
}