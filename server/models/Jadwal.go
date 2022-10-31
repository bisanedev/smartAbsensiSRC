package models

type Jadwal struct {
	Model						
	KelasID		uint		`json:"kelas_id,omitempty"`	
	UserID		uint		`json:"user_id,omitempty"`
	Hari 		string		`gorm:"type:varchar(10);not null;default:''" json:"hari,omitempty"`
	Jam 		uint		`json:"jam,omitempty"`
	JamStart	string		`gorm:"type:varchar(10);not null;default:''" json:"jam_start,omitempty"`
	JamEnd 		string		`gorm:"type:varchar(10);not null;default:''" json:"jam_end,omitempty"`
	Kelas 		Kelas		`gorm:"foreignkey:KelasID"`
	Users 		Users		`gorm:"foreignkey:UserID"`
}