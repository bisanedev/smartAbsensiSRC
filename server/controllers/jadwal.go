package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"../models"
	"../database"	
	"net/http"
	"strconv"
)

type Jadwal struct {
	Basic		
}

func (a *Jadwal) Index (c *gin.Context) {
	kelas, _	:= strconv.Atoi((*c).DefaultQuery("kelas", "0"))
	guru, _		:= strconv.Atoi((*c).DefaultQuery("guru", "0"))

	var jadwal []models.Jadwal
	db := database.DB

	if(kelas != 0){

		if err := db.Preload("Users", func(db *gorm.DB) *gorm.DB {return db.Where("status = ?","a").Select("id, status, nama, jenis, mapel, color, foto")}).Preload("Kelas").Where("kelas_id = ?",kelas).Order("jam asc").Find(&jadwal).Error;err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}

		(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": jadwal})

	}else if (guru != 0){
		if err := db.Preload("Users", func(db *gorm.DB) *gorm.DB {return db.Where("status = ?","a").Select("id, status, nama, jenis, mapel, color, foto")}).Preload("Kelas").Where("user_id = ?",guru).Order("jam asc").Find(&jadwal).Error;err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}

		(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": jadwal})
		
	}else{(*a).JsonFail(c, http.StatusBadRequest, "Data Tidak Ditemukan")}

}

func (a *Jadwal) JadwalAbsen (c *gin.Context) {
	kelas, _		:= strconv.Atoi((*c).DefaultQuery("kelas", "0"))
	hari 			:= (*c).DefaultQuery("hari", "null")
	var jadwal []models.Jadwal
	db := database.DB

	if (kelas == 0 || hari == "null"){
		(*a).JsonFail(c, http.StatusBadRequest, "Parameter Query Undefined")
	}else{

	db.Preload("Users", func(db *gorm.DB) *gorm.DB {return db.Where("status = ?","a").Select("id, status, nama, jenis, mapel, color, foto")}).Where("kelas_id = ?",kelas).Where("hari = ?",hari).Find(&jadwal)
	
	(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": jadwal})

	}
}

func (a *Jadwal) CekMapel (c *gin.Context) {
	kelas, _	:= strconv.Atoi((*c).DefaultQuery("kelas", "0"))
	hari		:= (*c).DefaultQuery("hari", "Senin")

	var jadwal []models.Jadwal
	var cekmapel []int64

	if err := database.DB.Joins("JOIN users ON users.id = user_id").Where("status = ?","a").Where("kelas_id = ?",kelas).Where("hari = ?",hari).Order("jam asc").Find(&jadwal).Pluck("jam", &cekmapel).Error;err != nil {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
		return
	}

	(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": cekmapel})
}	

func (a *Jadwal) Store (c *gin.Context) {
	// start store
	var request RequestJadwal

	if err := (*c).ShouldBind(&request); err == nil {
		
		jadwal := &models.Jadwal{			
			KelasID:  request.KelasID,
			UserID:   request.UserID,
			Hari:   request.Hari,
			Jam:   request.Jam,
			JamStart:   request.JamStart,
			JamEnd:   request.JamEnd,
		}

		if err := database.DB.Create(&jadwal).Error; err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}		
				
		(*a).JsonSuccess(c, http.StatusCreated, gin.H{"message": "Created successfully"})		
		
	} else {
		(*a).JsonFail(c, http.StatusBadRequest, "Silahkan Periksa Kembali Input Data !")
	}
	//end role
}

func (a *Jadwal) Destroy (c *gin.Context) {
	var jadwal models.Jadwal

	if database.DB.First(&jadwal, (*c).Param("id")).Error != nil {
		(*a).JsonFail(c, http.StatusNotFound, "Data Not Found")
		return
	}
	
	// delete jadwal
	if err := database.DB.Unscoped().Delete(&jadwal).Error; err != nil {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
		return
	}	
	(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
}

func (a *Jadwal) Update (c *gin.Context) {
	var request RequestJadwal
	if err := (*c).ShouldBind(&request); err == nil {
		var jadwal models.Jadwal
		if database.DB.First(&jadwal, (*c).Param("id")).Error != nil {
			(*a).JsonFail(c, http.StatusNotFound, "Data does not exist")
			return
		}
		
		jadwal.KelasID = request.KelasID
		jadwal.UserID = request.UserID
		jadwal.Hari = request.Hari
		jadwal.Jam = request.Jam
		jadwal.JamStart = request.JamStart
		jadwal.JamEnd = request.JamEnd

		if err := database.DB.Save(&jadwal).Error; err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}
				
		(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
	} else {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
	}
}

type RequestJadwal struct {
	KelasID		uint 	`form:"kelas_id" json:"kelas_id" binding:"required"`
	UserID		uint 	`form:"user_id" json:"user_id" binding:"required"`
	Hari		string  `form:"hari" json:"hari" binding:"required"`
	Jam			uint  	`form:"jam" json:"jam" binding:"required"`
	JamStart	string  `form:"jam_start" json:"jam_start" binding:"required"`
	JamEnd		string  `form:"jam_end" json:"jam_end" binding:"required"`
}