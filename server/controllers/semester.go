package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/biezhi/gorm-paginator/pagination"
	"../models"
	"../database"	
	"net/http"
	"strconv"
	"time"
)

type Semester struct {
	Basic		
}

func (a *Semester) Index (c *gin.Context) {

	page, _		:= strconv.Atoi((*c).DefaultQuery("page", "1"))
	limit, _	:= strconv.Atoi((*c).DefaultQuery("limit", "3"))
	search		:= (*c).DefaultQuery("search", "%")
	order 		:= (*c).DefaultQuery("order", "id")
	by 			:= (*c).DefaultQuery("by", "desc")	

	var semester []models.Semester	

	paginator := pagination.Paging(&pagination.Param{
			DB:      database.DB.Where("nama LIKE ?","%"+search+"%"),			
			Page:    page,
			Limit:   limit,			
			OrderBy: []string{order+" "+by},			
			ShowSQL: false,
		}, &semester)
	(*c).JSON(200, paginator)
}

func (a *Semester) Store (c *gin.Context) {		
	//role akses
	if (*a).Role(c,"admin") {
	// start role
	var request RequestSemester

	if err := (*c).ShouldBind(&request); err == nil {

		waktustart, _:= time.Parse("2006-01-02", request.Waktustart)
		waktuend, _:= time.Parse("2006-01-02", request.Waktuend)

		semester := &models.Semester{			
			Nama:  request.Nama,
			Semester:  request.Semester,
			Waktustart:  waktustart,
			Waktuend: waktuend,
		}

		if err := database.DB.Create(&semester).Error; err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}		
				
		(*a).JsonSuccess(c, http.StatusCreated, gin.H{"message": "Created successfully"})		
		
	} else {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
	}
	//else role
	}else{
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
	//end role
}

func (a *Semester) Destroy (c *gin.Context) {
	//role akses
	if (*a).Role(c,"admin"){
	//----------	
	var semester models.Semester

	if database.DB.First(&semester, (*c).Param("id")).Error != nil {
		(*a).JsonFail(c, http.StatusNotFound, "Data Not Found")
		return
	}
	
	// delete kelas
	if err := database.DB.Unscoped().Delete(&semester).Error; err != nil {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
		return
	}	
	(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
	//else role
	}else{
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
}

func (a *Semester) Update (c *gin.Context) {
	//role akses
	if (*a).Role(c,"admin") {
	//----------	
	var request RequestSemester
	if err := (*c).ShouldBind(&request); err == nil {
		var semester models.Semester

		if database.DB.First(&semester, (*c).Param("id")).Error != nil {
			(*a).JsonFail(c, http.StatusNotFound, "Data does not exist")
			return
		}
		waktustart, _:= time.Parse("2006-01-02", request.Waktustart)
		waktuend, _:= time.Parse("2006-01-02", request.Waktuend)

		semester.Nama = request.Nama
		semester.Semester = request.Semester
		semester.Waktustart = waktustart
		semester.Waktuend = waktuend

		if err := database.DB.Save(&semester).Error; err != nil {
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

func (a *Semester) ListAll (c *gin.Context) {

	var semester []ListSemester	
	database.DB.Table("Semesters").Select("CONCAT(semester, ' TA ',nama ) AS nama,waktustart").Order("nama desc").Scan(&semester)
	(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": semester})

}

type ListSemester struct {		
	Nama			string 			`json:"nama"`	
	Waktustart		string 			`json:"waktustart"`	
}

type RequestSemester struct {		
	Nama			string 			`form:"nama" json:"nama" binding:"required"`
	Semester		string 			`form:"semester" json:"semester" binding:"required"`
	Waktustart		string 			`form:"waktustart" json:"waktustart" binding:"required"`
	Waktuend		string 			`form:"waktuend" json:"waktuend" binding:"required"`
}