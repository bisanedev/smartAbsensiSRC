package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/biezhi/gorm-paginator/pagination"
	"../models"
	"../database"	
	"net/http"
	"strconv"	
)

type Siswa struct {
	Basic		
}

func (a *Siswa) Index (c *gin.Context) {

	page, _		:= strconv.Atoi((*c).DefaultQuery("page", "1"))
	limit, _	:= strconv.Atoi((*c).DefaultQuery("limit", "3"))	
	search		:= (*c).DefaultQuery("search", "%")
	order 		:= (*c).DefaultQuery("order", "id")
	by 			:= (*c).DefaultQuery("by", "desc")
	kelas		:= (*c).DefaultQuery("kelas", "null")

	var siswa []models.Siswa	

	if (kelas == "null"){
		paginator := pagination.Paging(&pagination.Param{
				DB:      database.DB.Preload("Kelas").Where("status = ?","a").Where("nama LIKE ?","%"+search+"%"),			
				Page:    page,
				Limit:   limit,			
				OrderBy: []string{order+" "+by+", urut"},			
				ShowSQL: false,
			}, &siswa)
		(*c).JSON(200, paginator)
	}else{
		paginator := pagination.Paging(&pagination.Param{
				DB:      database.DB.Preload("Kelas").Where("status = ?","a").Where("nama LIKE ?","%"+search+"%").Where("kelas_id = ?",kelas),			
				Page:    page,
				Limit:   limit,			
				OrderBy: []string{order+" "+by+", urut"},			
				ShowSQL: false,
			}, &siswa)
		(*c).JSON(200, paginator)
	}
}

func (a *Siswa) Arsip (c *gin.Context){

	page, _		:= strconv.Atoi((*c).DefaultQuery("page", "1"))
	limit, _	:= strconv.Atoi((*c).DefaultQuery("limit", "3"))	
	search		:= (*c).DefaultQuery("search", "%")
	order 		:= (*c).DefaultQuery("order", "id")
	by 			:= (*c).DefaultQuery("by", "desc")
	kelas		:= (*c).DefaultQuery("kelas", "null")

	var siswa []models.Siswa

	if (kelas == "null"){
		paginator := pagination.Paging(&pagination.Param{
				DB:      database.DB.Preload("Kelas").Not("status = ?","a").Where("nama LIKE ?","%"+search+"%"),			
				Page:    page,
				Limit:   limit,			
				OrderBy: []string{order+" "+by+", urut"},			
				ShowSQL: false,
			}, &siswa)
		(*c).JSON(200, paginator)
	}else{
		paginator := pagination.Paging(&pagination.Param{
				DB:      database.DB.Preload("Kelas").Not("status = ?","a").Where("nama LIKE ?","%"+search+"%").Where("kelas_id = ?",kelas),			
				Page:    page,
				Limit:   limit,			
				OrderBy: []string{order+" "+by+", urut"},			
				ShowSQL: false,
			}, &siswa)
		(*c).JSON(200, paginator)
	}

}

func (a *Siswa) Multi (c *gin.Context){
	if (*a).Role(c,"admin"){
	// roleakses	
	x := &MultiEdit{}	

	if err := (*c).ShouldBind(&x); err == nil {

		database.DB.Exec("UPDATE siswas SET status=? , kelas_id=? WHERE id IN (?)", x.Status, x.KelasID, x.Selected)
				
		(*a).JsonSuccess(c, http.StatusCreated, gin.H{})	

	}else{(*a).JsonFail(c, http.StatusBadRequest, "Silahkan Periksa Kembali Input Data !")}
	//else role
	}else{(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")}
}	

func (a *Siswa) MultiDelete (c *gin.Context){
	if (*a).Role(c,"admin"){
	// roleakses
	x := &MultiHapus{}

	type Result struct {
	  Foto string	  
	}	

	var foto []Result
	
	if err := (*c).ShouldBind(&x); err == nil {	

    database.DB.Raw("SELECT foto from siswas WHERE id IN (?)", x.Selected).Scan(&foto)	

    for _, s := range foto {
    	if(s.Foto != ""){
    		(*a).fileRem("./storage/siswa/"+s.Foto);
    	}
	}

	database.DB.Exec("DELETE from siswas WHERE id IN (?)", x.Selected)

	(*a).JsonSuccess(c, http.StatusCreated, gin.H{})

	}else{(*a).JsonFail(c, http.StatusBadRequest, "Silahkan Periksa Kembali Input Data !")}
	//else role
	}else{(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")}
}	

func (a *Siswa) Jumlah (c *gin.Context){

	kelas := (*c).DefaultQuery("kelas", "0")

	var count uint	

	if (*a).Role(c,"admin") {		
	database.DB.Model(&models.Siswa{}).Where("kelas_id = ?", kelas).Where("status = ?","a").Count(&count)
	(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": count})
	}else{(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")}
}

func (a *Siswa) Store (c *gin.Context) {		
	//role akses
	if (*a).Role(c,"admin") {
	// start role
	var request RequestSiswa

	if err := (*c).ShouldBind(&request); err == nil {

		siswa := &models.Siswa{			
			Nama:     	  request.Nama,
			Status:    	  request.Status,
			Jenis:	  	  request.Jenis,
			KelasID:	  request.KelasID,
			Urut:	  	  request.Urut,
		}

		if err := database.DB.Create(&siswa).Error; err != nil {
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

func (a *Siswa) Destroy (c *gin.Context) {
	//role akses
	if (*a).Role(c,"admin"){
	//----------	
	var siswa models.Siswa

	if database.DB.First(&siswa, (*c).Param("id")).Error != nil {
		(*a).JsonFail(c, http.StatusNotFound, "Data Not Found")
		return
	}
	// delete photo
	if (siswa.Foto != ""){
		(*a).fileRem("./storage/siswa/"+siswa.Foto); 
	}
	// delete siswa
	if err := database.DB.Unscoped().Delete(&siswa).Error; err != nil {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
		return
	}	
	(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
	//else role
	}else{
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
}

func (a *Siswa) All (c *gin.Context) {
if (*a).Role(c,"admin") {

	var siswa []models.Siswa	

	database.DB.Order("nama asc").Find(&siswa)

	(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": siswa})	
	}else{(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")}
}

func (a *Siswa) Update (c *gin.Context) {
	//role akses
	if (*a).Role(c,"admin") {
	//----------	
	var request RequestSiswa
	if err := (*c).ShouldBind(&request); err == nil {
		var siswa models.Siswa

		if database.DB.First(&siswa, (*c).Param("id")).Error != nil {
			(*a).JsonFail(c, http.StatusNotFound, "Data does not exist")
			return
		}
		
		siswa.Nama = request.Nama
		siswa.Jenis = request.Jenis		
		siswa.KelasID = request.KelasID
		siswa.Urut = request.Urut
		siswa.Status = request.Status

		if err := database.DB.Save(&siswa).Error; err != nil {
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

type MultiHapus struct {  
	Selected    []uint  `form:"selected[]" binding:"required"`
}

type MultiEdit struct {    
    Selected    []uint  `form:"selected[]" binding:"required"`
    Status		string  `form:"status" json:"status" binding:"required"`
    KelasID		uint 	`form:"kelas_id" json:"kelas_id" binding:"required"`
}

type RequestSiswa struct {	
	Nama		string  `form:"nama" json:"nama" binding:"required"`
	Status		string  `form:"status" json:"status" binding:"required"`
	Jenis		string  `form:"jenis" json:"jenis" binding:"required"`
	KelasID		uint 	`form:"kelas_id" json:"kelas_id" binding:"required"`
	Urut		uint  	`form:"urut" json:"urut" binding:"required"`
}