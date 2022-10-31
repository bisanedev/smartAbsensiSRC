package models

import "time"

type Absen struct {
	ID        		uint 			`gorm:"primary_key" json:"id"`
	CreatedAt 		time.Time 		`json:"created_at,omitempty"`
	UpdatedAt 		time.Time		`json:"updated_at,omitempty"`	
	SemesterID 		uint 			`json:"semester_id,omitempty"`
	KelasID			uint			`json:"kelas_id,omitempty"`
	UserID			uint			`json:"user_id,omitempty"`
	Tanggal			time.Time 		`json:"tanggal,omitempty"`	
	Keluar 			time.Time 		`json:"keluar,omitempty"`
	Jam 			uint			`json:"jam,omitempty"`
	Semester		string 			`gorm:"type:varchar(100);not null;default:''" json:"semester,omitempty"`
	Kelas			string 			`gorm:"type:varchar(100);not null;default:''" json:"kelas,omitempty"`
	Mapel			string 			`gorm:"type:varchar(30);null;default:''" json:"mapel,omitempty"`
	MapelGuru		string			`gorm:"type:varchar(100);not null;default:''" json:"mapel_guru,omitempty"`
	MapelStart		time.Time		`json:"mapel_start,omitempty"`
	MapelEnd		time.Time		`json:"mapel_end,omitempty"`
	PiketID			uint			`json:"piket_id,omitempty"`
	PiketGuru		string 			`gorm:"type:varchar(300);null;default:''" json:"piket_guru,omitempty"`
	Siswa			[]AbsenChild 	`gorm:"foreignkey:AbsenID"`	
}