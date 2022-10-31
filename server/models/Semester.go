package models

import "time"

type Semester struct {
	Model			
	Nama			string 			`gorm:"type:varchar(300);not null;default:''" json:"nama,omitempty"`
	Semester		string 			`gorm:"type:varchar(100);not null;default:''" json:"semester,omitempty"`
	Waktustart		time.Time 		`json:"waktustart,omitempty"`
	Waktuend		time.Time 		`json:"waktuend,omitempty"`
	Absen			[]Absen 		`gorm:"foreignkey:SemesterID"`
}