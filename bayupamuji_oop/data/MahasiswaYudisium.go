package data

type MahasiswaYudisium struct {
	isLulusSidang bool
	Mahasiswa
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