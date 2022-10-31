package models

type AbsenChild struct {	
	ID        		uint 		`gorm:"primary_key" json:"id"`
	AbsenID			uint 		`json:"absen_id,omitempty"`
	SiswaID			uint 		`json:"siswa_id,omitempty"`
	Urut			uint 		`json:"urut,omitempty"`
	Nama			string		`gorm:"type:varchar(100);not null;default:''" json:"nama,omitempty"`
	Jenis			string 		`gorm:"type:varchar(20);null;default:''" json:"jenis,omitempty"`
	Foto			string 		`gorm:"type:varchar(50);null;default:''" json:"foto,omitempty"`
	Kode 			string 		`gorm:"type:varchar(5);null;default:''" json:"kode,omitempty"`
	KodeNote 		string 		`gorm:"type:varchar(300);null;default:''" json:"kode_note,omitempty"`
}