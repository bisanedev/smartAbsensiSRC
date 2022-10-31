package models

type Kelas struct {
	Model
	Status	string		`gorm:"type:varchar(2);not null;default:''" json:"status,omitempty"`		
	Nama 	string		`gorm:"type:varchar(100);not null;default:''" json:"nama,omitempty"`	
	Maxpel 	uint		`json:"maxpel,omitempty"`
	Siswa	[]Siswa		`gorm:"foreignkey:KelasID"`
}