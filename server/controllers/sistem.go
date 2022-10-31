package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"../config"
	"../database"
	"../models"
	"github.com/JamesStewy/go-mysqldump"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Sistem struct {
	Basic
}

func (a *Sistem) CheckAuth(c *gin.Context) {

	(*a).JsonSuccess(c, http.StatusOK, gin.H{"color": "#5eba00"})

}

func (a *Sistem) GetPassword(c *gin.Context) {
	if (*a).Role(c, "admin") {
		var sistem models.Sistem

		database.DB.First(&sistem)

		(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": sistem})
		//else role
	} else {
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
	//end role
}

func (a *Sistem) PostPasswordPiket(c *gin.Context) {
	if (*a).Role(c, "admin") {
		var request RequestPiket
		if err := (*c).ShouldBind(&request); err == nil {
			var sistem models.Sistem
			//cek data ada atau tidak
			if database.DB.First(&sistem, request.ID).Error != nil {
				(*a).JsonFail(c, http.StatusBadRequest, err.Error())
				return
			}
			//jika data ada maka update
			sistem.PasswordPiket = request.PasswordPiket

			if err := database.DB.Save(&sistem).Error; err != nil {
				(*a).JsonFail(c, http.StatusBadRequest, err.Error())
				return
			}

			(*a).JsonSuccess(c, http.StatusCreated, gin.H{})

			// ----------------------
		} else {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
		}
		//else role
	} else {
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
	//end role
}

func (a *Sistem) AuthPiket(c *gin.Context) {
	var request RequestPiket
	if err := (*c).ShouldBind(&request); err == nil {
		var sistem models.Sistem
		if database.DB.Where("id = ?", request.ID).Where("password_piket = ?", request.PasswordPiket).First(&sistem).Error != nil {
			(*a).JsonFail(c, http.StatusUnauthorized, "Wrong Password")
			return
		}
		(*a).JsonSuccess(c, http.StatusOK, gin.H{})
		// ----------------------
	} else {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
	}
}

func (a *Sistem) PostPasswordClock(c *gin.Context) {
	if (*a).Role(c, "admin") {
		var request RequestClock
		if err := (*c).ShouldBind(&request); err == nil {
			var sistem models.Sistem
			//cek data ada atau tidak
			if database.DB.First(&sistem, request.ID).Error != nil {
				(*a).JsonFail(c, http.StatusBadRequest, err.Error())
				return
			}
			//jika data ada maka update
			sistem.PasswordClock = strings.ToLower(request.PasswordClock)

			if err := database.DB.Save(&sistem).Error; err != nil {
				(*a).JsonFail(c, http.StatusBadRequest, err.Error())
				return
			}

			(*a).JsonSuccess(c, http.StatusCreated, gin.H{})

			// ----------------------
		} else {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
		}
		//else role
	} else {
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
	//end role
}

//do backup
func (a *Sistem) DoBackupDatabase(c *gin.Context) {
	if (*a).Role(c, "admin") {
		conf := config.Get()
		//create database
		if _, err := os.Stat("./backup"); os.IsNotExist(err) {
			os.Mkdir("./backup", 0700)
		}
		dumpDir := "backup"                                               // you should create this directory
		dumpFilenameFormat := fmt.Sprintf("%s-20060102T150405", "backup") // accepts time layout string and add .sql at the end of file
		db, err := sql.Open("mysql", conf.UsernameDB+":"+conf.PasswordDB+"@tcp("+conf.HostDB+")/"+conf.NamaDB+"?charset=utf8&parseTime=True&loc=Local")

		if err != nil {
			fmt.Println("Error opening database: ", err)
			return
		}

		// Register database with mysqldump
		dumper, err := mysqldump.Register(db, dumpDir, dumpFilenameFormat)
		if err != nil {
			fmt.Println("Error registering databse:", err)
			return
		}

		// Dump database to file
		resultFilename, err := dumper.Dump()

		if err != nil {
			fmt.Println("Error dumping:", err)
			return
		}

		filenya := strings.Split(resultFilename, "/")
		//fmt.Printf("File is saved to %s", resultFilename)
		var sistem models.Sistem
		//cek data ada atau tidak
		if database.DB.First(&sistem, 1).Error != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}
		//jika data ada maka update
		sistem.BackupDatabase = filenya[1]

		if err := database.DB.Save(&sistem).Error; err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}

		(*a).JsonSuccess(c, http.StatusOK, gin.H{"file": filenya[1]})
		// Close dumper and connected database
		dumper.Close()
		//else role
	} else {
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
}

type RequestPiket struct {
	ID            uint   `form:"id" json:"id" binding:"required"`
	PasswordPiket string `form:"password_piket" json:"password_piket" binding:"required"`
}

//request
type RequestClock struct {
	ID            uint   `form:"id" json:"id" binding:"required"`
	PasswordClock string `form:"password_clock" json:"password_clock" binding:"required"`
}