package models

type Users struct {
	Model
	Status			string		`gorm:"type:varchar(2);not null;default:''" json:"status,omitempty"`
	Username 		string		`gorm:"type:varchar(30);unique_index" json:"username,omitempty"`
	Nama			string		`gorm:"type:varchar(100);not null;default:''" json:"nama,omitempty"`
	Jenis			string 		`gorm:"type:varchar(20);null;default:''" json:"jenis,omitempty"`
	Mapel			string 		`gorm:"type:varchar(30);null;default:''" json:"mapel,omitempty"`
	Color			string 		`gorm:"type:varchar(25);null;default:''" json:"color,omitempty"`
	Foto			string 		`gorm:"type:varchar(50);null;default:''" json:"foto,omitempty"`
	Role			string 		`gorm:"type:varchar(30);not null;default:''" json:"role,omitempty"`
	Password		string		`gorm:"type:varchar(128);not null;default:''" json:"password,empty"`		
	Token			int64		`json:"token,omitempty"`
}
