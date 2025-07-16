package models

type Siswa struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Nama  string `json:"nama"`
	Kelas string `json:"kelas"`
	Umur  int    `json:"umur"`
}