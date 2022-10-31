package models

type Sistem struct {
	Model
	PasswordPiket	string		`gorm:"type:varchar(40);not null;default:''" json:"password_piket,omitempty"`
	PasswordClock	string		`gorm:"type:varchar(40);not null;default:''" json:"password_clock,omitempty"`
	BackupDatabase	string		`gorm:"type:varchar(40);not null;default:''" json:"backup_database,omitempty"`	
}