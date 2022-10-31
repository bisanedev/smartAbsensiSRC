package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/biezhi/gorm-paginator/pagination"
	"../models"
	"../database"	
	"net/http"
	"strconv"
)

type Kelas struct {
	Basic		
}

func (a *Kelas) Index (c *gin.Context) {

	page, _		:= strconv.Atoi((*c).DefaultQuery("page", "1"))
	limit, _	:= strconv.Atoi((*c).DefaultQuery("limit", "3"))
	search		:= (*c).DefaultQuery("search", "%")
	order 		:= (*c).DefaultQuery("order", "id")
	by 			:= (*c).DefaultQuery("by", "desc")
	status		:= (*c).DefaultQuery("status", "a")

	var kelas []models.Kelas	

	paginator := pagination.Paging(&pagination.Param{
			DB:      database.DB.Where("status = ?",status).Where("nama LIKE ?","%"+search+"%"),			
			Page:    page,
			Limit:   limit,			
			OrderBy: []string{order+" "+by},			
			ShowSQL: false,
		}, &kelas)
	(*c).JSON(200, paginator)
}

func (a *Kelas) Store (c *gin.Context) {		
	//role akses
	if (*a).Role(c,"admin") {
	// start role
	var request RequestKelas

	if err := (*c).ShouldBind(&request); err == nil {

		kelas := &models.Kelas{			
			Status:  request.Status,
			Nama:  request.Kelas,
			Maxpel: request.Maxpel,
		}

		if err := database.DB.Create(&kelas).Error; err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}		
				
		(*a).JsonSuccess(c, http.StatusCreated, gin.H{"message": "Created successfully"})		
		
	} else {
		(*a).JsonFail(c, http.StatusBadRequest, "Silahkan Periksa Kembali Input Data !")
	}
	//else role
	}else{
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
	//end role
}

func (a *Kelas) Destroy (c *gin.Context) {
	//role akses
	if (*a).Role(c,"admin"){
	//----------	
	var kelas models.Kelas

	if database.DB.First(&kelas, (*c).Param("id")).Error != nil {
		(*a).JsonFail(c, http.StatusNotFound, "Data Not Found")
		return
	}
	
	// delete kelas
	if err := database.DB.Unscoped().Delete(&kelas).Error; err != nil {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
		return
	}	
	(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
	//else role
	}else{
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
}

func (a *Kelas) Update (c *gin.Context) {
	//role akses
	if (*a).Role(c,"admin") {
	//----------	
	var request RequestKelas
	if err := (*c).ShouldBind(&request); err == nil {
		var kelas models.Kelas
		if database.DB.First(&kelas, (*c).Param("id")).Error != nil {
			(*a).JsonFail(c, http.StatusNotFound, "Data does not exist")
			return
		}
		
		kelas.Nama = request.Kelas
		kelas.Status = request.Status
		kelas.Maxpel = request.Maxpel

		if err := database.DB.Save(&kelas).Error; err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}
				
		(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
	} else {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
	}
	//else role
	}else{
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
}

func (a *Kelas) All (c *gin.Context) {

	var kelas []models.Kelas	

	database.DB.Where("status = ?","a").Order("LENGTH(nama),nama asc").Find(&kelas)

	(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": kelas})		
}

type RequestKelas struct {	
	Kelas		string  `form:"kelas" json:"kelas" binding:"required"`
	Status		string  `form:"status" json:"status" binding:"required"`
	Maxpel		uint  	`form:"maxpel" json:"maxpel" binding:"required"`	
}