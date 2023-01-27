package data

type Mahasiswa struct {
	nim      int
	name     string
	prodi    string
	semester int
}

func (mhs *Mahasiswa) SetNim(nim int) {
	mhs.nim = nim
}

func (mhs *Mahasiswa) SetName(name string) {
	mhs.name = name
}

func (mhs *Mahasiswa) SetProdi(prodi string) {
	mhs.prodi = prodi
}

func (mhs *Mahasiswa) SetSemester(semester int) {
	mhs.semester = semester
}

func (mhs *Mahasiswa) GetNim() int {
	return mhs.nim
}

func (mhs *Mahasiswa) GetName() string {
	return mhs.name
}

func (mhs *Mahasiswa) GetProdi() string {
	return mhs.prodi
}

func (mhs *Mahasiswa) GetSemester() int {
	return mhs.semester
}