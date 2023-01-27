package data

type Mahasiswa struct {
	nim      int
	name     string
	prodi    string
	semester int
	ipk      float64
}

func (mhs *Mahasiswa) MaxSks() int {
	maxSks := 0
	switch {
	case mhs.ipk > 3.00:
		maxSks = 24
	case mhs.ipk < 3.00 && mhs.ipk > 2.5:
		maxSks = 20
	default:
		maxSks = 18
	}
	return maxSks
}

// function setter

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

func (mhs *Mahasiswa) SetIpk(ipk float64) {
	mhs.ipk = ipk
}

// function getter

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

func (mhs *Mahasiswa) GetIpk() float64 {
	return mhs.ipk
}