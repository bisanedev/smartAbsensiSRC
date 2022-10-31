package database

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"

	"../config"
	"../models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"		
	"database/sql"		
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	conf := config.Get()
	// create database if not exits
	dbCreate, err2 := sql.Open("mysql", conf.UsernameDB+":"+conf.PasswordDB+"@tcp("+conf.HostDB+")/?charset=utf8&parseTime=True&loc=Local")
	if err2 != nil {		
		fmt.Println("Koneksi Database Tak Tersambung :", err2)
	}
	defer dbCreate.Close()
	_,err3 := dbCreate.Exec("CREATE DATABASE IF NOT EXISTS "+conf.NamaDB)
	if err3 != nil {		
		fmt.Println("Koneksi Database Tak Tersambung ! :", err3)
	}
	// end create database if not exits
	db, err := gorm.Open("mysql", conf.UsernameDB+":"+conf.PasswordDB+"@tcp("+conf.HostDB+")/"+conf.NamaDB+"?charset=utf8&parseTime=True&loc=Local")
	// debug
	//db.LogMode(true)
	// ramdom word
	// ramdon Password piket
	rand.Seed(time.Now().Unix())
	randomlah := []string{
		"MulaiBelajar",
		"MakinTau",
		"DimanaSaja",
		"KuraKuraNinja",
		"PuraPuraTau",
		"SayaKurangTahu",
		"AyamJantan",
		"LadaHitam",
		"LaparKenyang",
		"PakKucing",
		"SuaraMerdu",
		"SendiriSaja",
		"PakMisnaden",
		"TangguhkanDulu",
	}
	n := rand.Int() % len(randomlah)
	// Error
	if err == nil {
		db.DB().SetMaxIdleConns(500)
		DB = db
		// Creating the table
		db.AutoMigrate(&models.Users{}, &models.Sistem{}, &models.Kelas{}, &models.Siswa{}, &models.Jadwal{}, &models.Semester{}, &models.Absen{}, &models.AbsenChild{})
		db.Model(models.Siswa{}).AddForeignKey("kelas_id", "kelas(id)", "SET NULL", "RESTRICT")
		db.Model(models.Jadwal{}).AddForeignKey("kelas_id", "kelas(id)", "SET NULL", "RESTRICT")
		db.Model(models.Jadwal{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
		db.Model(models.Absen{}).AddForeignKey("kelas_id", "kelas(id)", "SET NULL", "RESTRICT")
		db.Model(models.Absen{}).AddForeignKey("user_id", "users(id)", "SET NULL", "RESTRICT")
		db.Model(models.Absen{}).AddForeignKey("semester_id", "semesters(id)", "SET NULL", "RESTRICT")
		db.Model(models.AbsenChild{}).AddForeignKey("absen_id", "absens(id)", "CASCADE", "RESTRICT")
		db.Model(models.AbsenChild{}).AddForeignKey("siswa_id", "siswas(id)", "SET NULL", "RESTRICT")
		//--------
		count := 0
		db.Model(models.Users{}).Count(&count)
		if count == 0 {
			// create user admin
			password := []byte("password")
			md5Ctx := md5.New()
			md5Ctx.Write(password)
			cipherStr := md5Ctx.Sum(nil)
			user := &models.Users{
				Status:   "a",
				Username: strings.ToLower("admin"),
				Nama:     "Administrator",
				Jenis:    "Laki-Laki",
				Role:     strings.ToLower("superuser"),
				Password: hex.EncodeToString(cipherStr),
			}
			db.Create(&user)
		}
		// create password piket randomly
		piket := 0
		db.Model(models.Sistem{}).Count(&piket)
		if piket == 0 {
			sistem := &models.Sistem{
				PasswordPiket: randomlah[n],
				PasswordClock: strings.ToLower(randomlah[n]),
			}
			db.Create(&sistem)
		}
		// enddmdmada
		return db, err

	}

	return nil, err
}
