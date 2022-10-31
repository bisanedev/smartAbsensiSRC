package controllers

import (	
	"net/http"
	"path/filepath"	
	"github.com/gin-gonic/gin"
	"../models"
	"../database"
	"time"
	"strconv"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"log"
	"os"
	"strings"
)

type Uploader struct {
	Basic	
}

func (a *Uploader) StorageGuru (c *gin.Context) {		
		// update photo db		
		var user models.Users
		if database.DB.First(&user, (*c).PostForm("id")).Error != nil {
			(*a).JsonFail(c, http.StatusNotFound, "Data does not exist")
			return
		}
			
		// single upload		
		file, err := (*c).FormFile("files")
		nameFile := strings.Split(file.Filename,".")
		if err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}
		//remove if exist
	   	if (user.Foto != ""){
			(*a).fileRem("./storage/guru/"+user.Foto); 
		}
	    	    
		path := filepath.Join("./storage/guru", file.Filename )
		if err := (*c).SaveUploadedFile(file, path); err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())				
			return
		}
		//convert webp
		img, err := imaging.Open(path, imaging.AutoOrientation(true))
		if err != nil {
			log.Fatal(err)
		}
		unixname := strconv.FormatInt(time.Now().Unix(), 10)+"_"+nameFile[0]+".webp";
		konfersi := "./storage/guru/"+unixname;
		fileWebp, err := os.Create(konfersi)
		if err != nil {
			log.Fatal(err)
		}
		if err := webp.Encode(fileWebp, img, nil); err != nil {
		log.Fatal(err)
		}
		if err := fileWebp.Close(); err != nil {
			log.Fatal(err)
		}
		// update photo db
		user.Foto = unixname

		if err := database.DB.Save(&user).Error; err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}
		// hapus foto jpg
		(*a).fileRem(path);

		(*a).JsonSuccess(c, http.StatusOK, gin.H{"message": "Uploaded successfully"})			
}

func (a *Uploader) StorageSiswa (c *gin.Context) {		
		// update photo db		
		var siswa models.Siswa
		if database.DB.First(&siswa, (*c).PostForm("id")).Error != nil {
			(*a).JsonFail(c, http.StatusNotFound, "Data does not exist")
			return
		}
			
		// single upload		
		file, err := (*c).FormFile("files")
		nameFile := strings.Split(file.Filename,".")

		if err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}
		//remove if exist
	   	if (siswa.Foto != ""){
			(*a).fileRem("./storage/siswa/"+siswa.Foto); 
		}
	    	    
		path := filepath.Join("./storage/siswa", file.Filename )
		if err := (*c).SaveUploadedFile(file, path); err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())				
			return
		}
		//convert webp
		img, err := imaging.Open(path, imaging.AutoOrientation(true))
		if err != nil {
			log.Fatal(err)
		}
		unixname := strconv.FormatInt(time.Now().Unix(), 10)+"_"+nameFile[0]+".webp";
		konfersi := "./storage/siswa/"+unixname;
		fileWebp, err := os.Create(konfersi)
		if err != nil {
			log.Fatal(err)
		}
		if err := webp.Encode(fileWebp, img, nil); err != nil {
		log.Fatal(err)
		}
		if err := fileWebp.Close(); err != nil {
			log.Fatal(err)
		}
		// update photo di database absensi_child
		database.DB.Table("absen_children").Where("siswa_id = ?",(*c).PostForm("id")).Update("foto",unixname)
		// update photo db
		siswa.Foto = unixname

		if err := database.DB.Save(&siswa).Error; err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}
		// hapus foto jpg
		(*a).fileRem(path);

		(*a).JsonSuccess(c, http.StatusOK, gin.H{"message": "Uploaded successfully"})			
}