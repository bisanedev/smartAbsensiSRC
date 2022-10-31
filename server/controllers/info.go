package controllers

import (
	"net/http"

	"../database"
	"../models"
	"github.com/gin-gonic/gin"
)

type Info struct {
	Basic
}

//diagram pie kehadiran all siswa harian
func (a *Info) RekapHarian(c *gin.Context) {

	var jumlah JumlahInsight

	database.DB.Table("absen_children").Select("SUM(absen_children.kode='h') AS hadir,SUM(absen_children.kode='i') AS izin,SUM(absen_children.kode='s') AS sakit,SUM(absen_children.kode='t') AS terlambat,SUM(absen_children.kode='a') AS alpa").Joins("left join absens on absens.id = absen_children.absen_id").Where("DATE(absens.tanggal) = CURDATE()").Scan(&jumlah)

	(*a).JsonSuccess(c, http.StatusOK, gin.H{"jumlah": jumlah})
}

//cek scra realtime all siswa alpa,izin , terlmabat , sakit
func (a *Info) AbsenHarian(c *gin.Context) {
	var jumlah uint
	var absen []JumlahAbsenHarian
	database.DB.Model(&models.Siswa{}).Where("status = ?", "a").Count(&jumlah)
	database.DB.Table("absen_children").Select("absen_children.nama,absens.mapel_guru,absens.mapel,absens.mapel_start,absens.mapel_end,absens.kelas,absen_children.foto,absen_children.jenis,absen_children.kode,absen_children.kode_note").Joins("left join absens on absens.id = absen_children.absen_id").Where("absens.keluar= '0000-00-00 00:00:00'").Where("CURTIME() BETWEEN absens.mapel_start AND absens.mapel_end").Group("absen_children.siswa_id").Order("absens.jam desc").Scan(&absen)
	(*a).JsonSuccess(c, http.StatusOK, gin.H{"jumlah": jumlah, "absen": absen})
}

//diagram pie kehadiran All siswa per semester
func (a *Info) RekapSemester(c *gin.Context) {

	tanggal := (*c).DefaultQuery("tanggal", "null")

	var jumlah JumlahInsight

	database.DB.Table("absen_children").Select("SUM(absen_children.kode='h') AS hadir,SUM(absen_children.kode='i') AS izin,SUM(absen_children.kode='s') AS sakit,SUM(absen_children.kode='t') AS terlambat,SUM(absen_children.kode='a') AS alpa").Joins("left join absens on absens.id = absen_children.absen_id").Where("absens.semester_id = ?", CekSemesterID(tanggal)).Scan(&jumlah)

	(*a).JsonSuccess(c, http.StatusOK, gin.H{"jumlah": jumlah})
}

//diagram pie kehadiran All siswa Berdasarkan guru per semester
func (a *Info) InsightGuru(c *gin.Context) {

	tanggal := (*c).DefaultQuery("tanggal", "null")
	guru := (*c).DefaultQuery("guru", "0")

	var jumlah JumlahInsight

	database.DB.Table("absen_children").Select("SUM(absen_children.kode='h') AS hadir,SUM(absen_children.kode='i') AS izin,SUM(absen_children.kode='s') AS sakit,SUM(absen_children.kode='t') AS terlambat,SUM(absen_children.kode='a') AS alpa").Joins("left join absens on absens.id = absen_children.absen_id").Where("absens.user_id = ?", guru).Where("absens.semester_id = ?", CekSemesterID(tanggal)).Group("absens.user_id").Scan(&jumlah)
	semester, tahun := CekSemesterMulti(tanggal)
	(*a).JsonSuccess(c, http.StatusOK, gin.H{"jumlah": jumlah, "semester": semester, "tahun": tahun})

}

type JumlahAbsenHarian struct {
	Nama       string `json:"nama"`
	Jenis      string `json:"jenis"`
	Foto       string `json:"foto"`
	Mapel      string `json:"mapel"`
	MapelGuru  string `json:"mapel_guru"`
	MapelStart string `json:"mapel_start"`
	MapelEnd   string `json:"mapel_end"`
	Kelas      string `json:"kelas"`
	Kode       string `json:"kode"`
	KodeNote   string `json:"kode_note,omitempty"`
}

type JumlahInsight struct {
	Hadir     uint `json:"hadir"`
	Izin      uint `json:"izin"`
	Sakit     uint `json:"sakit"`
	Terlambat uint `json:"terlambat"`
	Alpa      uint `json:"alpa"`
}
