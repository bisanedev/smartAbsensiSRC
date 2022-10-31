package models

type Siswa struct {
	Model	
	Status			string		`gorm:"type:varchar(2);not null;default:''" json:"status,omitempty"`
	Nama			string		`gorm:"type:varchar(100);not null;default:''" json:"nama,omitempty"`
	Jenis			string 		`gorm:"type:varchar(20);null;default:''" json:"jenis,omitempty"`
	Foto			string 		`gorm:"type:varchar(50);null;default:''" json:"foto,omitempty"`
	KelasID			uint 		`json:"kelas_id,omitempty"`
	Urut			uint 		`json:"urut,omitempty"`
	Kelas 			Kelas 		`gorm:"foreignkey:KelasID"`
}