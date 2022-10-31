package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/biezhi/gorm-paginator/pagination"
	"../models"
	"../database"	
	"net/http"	
	"crypto/md5"
	"encoding/hex"	
	"strings"	
	"strconv"	
)

type Users struct {
	Basic	
}

func (a *Users) Index (c *gin.Context) {

	page, _		:= strconv.Atoi((*c).DefaultQuery("page", "1"))
	limit, _	:= strconv.Atoi((*c).DefaultQuery("limit", "3"))
	search		:= (*c).DefaultQuery("search", "%")
	order 		:= (*c).DefaultQuery("order", "id")
	by 			:= (*c).DefaultQuery("by", "desc")

	var users []models.Users	

	paginator := pagination.Paging(&pagination.Param{
			DB:      database.DB.Where("status = ?","a").Select("id, status, nama, jenis, username, mapel, color, foto, role").Where("nama LIKE ?","%"+search+"%"),			
			Page:    page,
			Limit:   limit,			
			OrderBy: []string{order+" "+by},			
			ShowSQL: false,
		}, &users)
	(*c).JSON(200, paginator)
}

func (a *Users) Arsip (c *gin.Context) {

	page, _		:= strconv.Atoi((*c).DefaultQuery("page", "1"))
	limit, _	:= strconv.Atoi((*c).DefaultQuery("limit", "3"))
	search		:= (*c).DefaultQuery("search", "%")
	order 		:= (*c).DefaultQuery("order", "id")
	by 			:= (*c).DefaultQuery("by", "desc")
	status		:= (*c).DefaultQuery("status", "null")

	var users []models.Users	

	if (status == "null"){
		paginator := pagination.Paging(&pagination.Param{
				DB:      database.DB.Not("status = ?","a").Select("id, status, nama, jenis, username, mapel, color, foto, role").Where("nama LIKE ?","%"+search+"%"),			
				Page:    page,
				Limit:   limit,			
				OrderBy: []string{order+" "+by},			
				ShowSQL: false,
			}, &users)
		(*c).JSON(200, paginator)		
	}else{
		paginator := pagination.Paging(&pagination.Param{
				DB:      database.DB.Where("status = ?",status).Select("id, status, nama, jenis, username, mapel, color, foto, role").Where("nama LIKE ?","%"+search+"%"),			
				Page:    page,
				Limit:   limit,			
				OrderBy: []string{order+" "+by},			
				ShowSQL: false,
			}, &users)
		(*c).JSON(200, paginator)
	}

}

func (a *Users) Store (c *gin.Context) {		
	//role akses
	if (*a).Role(c,"admin") {
	// start role
	var request CreateRequest
	if err := (*c).ShouldBind(&request); err == nil {
		var count int
		database.DB.Model(&models.Users{}).Where("username = ?", request.Username).Count(&count)

		if (count > 0) {
			(*a).JsonFail(c, http.StatusBadRequest, "Username already exists")
			return
		}

		password := []byte(request.Password)
		md5Ctx := md5.New()
		md5Ctx.Write(password)
		cipherStr := md5Ctx.Sum(nil)
		disablespace := strings.Replace(request.Username, " ", "", -1)
		user := &models.Users{
			Username: strings.ToLower(disablespace),
			Status:    request.Status,
			Nama:     request.Name,			
			Jenis:	  request.Jenis,
			Mapel:	  request.Mapel,
			Color:	  request.Color,
			Role:	  strings.ToLower(request.Role),
			Password: hex.EncodeToString(cipherStr),
		}

		if err := database.DB.Create(&user).Error; err != nil {
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

func (a *Users) Update (c *gin.Context) {
	//role akses
	if (*a).Role(c,"admin") {
	//----------	
	var request UpdateRequest
	if err := (*c).ShouldBind(&request); err == nil {
		var user models.Users
		if database.DB.First(&user, (*c).Param("id")).Error != nil {
			(*a).JsonFail(c, http.StatusNotFound, "Data does not exist")
			return
		}		
		disablespace := strings.Replace(request.Username, " ", "", -1)
		user.Status = request.Status
		user.Nama = request.Name
		user.Jenis = request.Jenis
		user.Username = strings.ToLower(disablespace)
		user.Mapel = request.Mapel
		user.Color = request.Color
		user.Role = request.Role

		if err := database.DB.Save(&user).Error; err != nil {
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

func (a *Users) All (c *gin.Context) {

	var user []models.Users 	
	database.DB.Select("id, nama, jenis, mapel, color, foto").Where("mapel !=''").Where("status = ?","a").Order("nama asc").Find(&user)
	(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": user})	
	
}

func (a *Users) MultiDelete (c *gin.Context) {
	if (*a).Role(c,"admin"){
	// roleakses
	x := &MultiHapus{}

	type Result struct {
	  Foto string	  
	}	

	var foto []Result
	
	if err := (*c).ShouldBind(&x); err == nil {

    database.DB.Raw("SELECT foto from users WHERE id IN (?)", x.Selected).Scan(&foto)	

    for _, s := range foto {
    	if(s.Foto != ""){
    		(*a).fileRem("./storage/guru/"+s.Foto);
    	}
	}

	database.DB.Exec("DELETE from users WHERE id IN (?)", x.Selected)

	(*a).JsonSuccess(c, http.StatusCreated, gin.H{})

	}else{(*a).JsonFail(c, http.StatusBadRequest, "Silahkan Periksa Kembali Input Data !")}
	//else role
	}else{(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")}
}

func (a *Users) Show (c *gin.Context) {
	var user models.Users	
	
	if database.DB.Select("id, status, nama, jenis, username, mapel, color, foto, role").First(&user, (*c).Param("id")).Error != nil {
		(*a).JsonFail(c, http.StatusNotFound, "Data Not Found")
		return
	}

	(*a).JsonSuccess(c, http.StatusCreated, gin.H{"data": user})
	
}

func (a *Users) Destroy (c *gin.Context) {
	//role akses
	if ((*a).Role(c,"admin") && (*c).Param("id") != "1"){
	//----------	
	var user models.Users

	if database.DB.First(&user, (*c).Param("id")).Error != nil {
		(*a).JsonFail(c, http.StatusNotFound, "Data Not Found")
		return
	}
	// delete photo
	if (user.Foto != ""){
		(*a).fileRem("./storage/guru/"+user.Foto); 
	}
	// delete user
	if err := database.DB.Unscoped().Delete(&user).Error; err != nil {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
		return
	}	
	(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
	//else role
	}else{
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
}

func (a *Users) Password (c *gin.Context) {	
	if (*a).Role(c,"admin") {
		//role akses
		var request UpdatePassword
		if err := (*c).ShouldBind(&request); err == nil {
			var user models.Users
			if database.DB.First(&user, (*c).Param("id")).Error != nil {
				(*a).JsonFail(c, http.StatusNotFound, "Data does not exist")
				return
			}
			if (request.Password != request.Repassword){
				(*a).JsonFail(c, http.StatusNotFound, "Konfirmasi Kata Sandi Tidak Sama!")
				return
			}
			password := []byte(request.Password)
			md5Ctx := md5.New()
			md5Ctx.Write(password)
			cipherStr := md5Ctx.Sum(nil)
			user.Password = hex.EncodeToString(cipherStr)			

			if err := database.DB.Save(&user).Error; err != nil {
				(*a).JsonFail(c, http.StatusBadRequest, err.Error())
				return
			}
					
			(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
		} else {
			(*a).JsonFail(c, http.StatusBadRequest, "Silahkan Periksa Kembali Input Data !")
		}
	//else role
	}else{
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
	//end role
}

type UpdatePassword struct {
	Password	string `form:"password" json:"password" binding:"required"`
	Repassword	string `form:"repassword" json:"repassword" binding:"required"`
}

type UpdateRequest struct {
	Status		string `form:"status" json:"status" binding:"required"`
	Username	string `form:"username" json:"username" binding:"required"`
	Name		string `form:"name" json:"name" binding:"required"`
	Jenis		string `form:"jenis" json:"jenis" binding:"required"`
	Mapel		string `form:"mapel" json:"mapel"`
	Color		string `form:"color" json:"color"`
	Role		string `form:"role" json:"role" binding:"required"`
}

type CreateRequest struct {
	Status		string `form:"status" json:"status" binding:"required"`
	Username	string `form:"username" json:"username" binding:"required"`
	Name		string `form:"name" json:"name" binding:"required"`
	Jenis		string `form:"jenis" json:"jenis" binding:"required"`
	Mapel		string `form:"mapel" json:"mapel"`
	Color		string `form:"color" json:"color"`
	Role		string `form:"role" json:"role" binding:"required"`
	Password	string `form:"password" json:"password" binding:"required"`
}