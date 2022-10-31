package controllers

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"

	//"net/http"
	"strconv"
	"time"

	"../database"
	//"../models"
)

type Rekap struct {
	Basic
}

func (a *Rekap) Siswa(c *gin.Context) {

	page, _ := strconv.Atoi((*c).DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi((*c).DefaultQuery("limit", "3"))
	search := (*c).DefaultQuery("search", "%")
	order := (*c).DefaultQuery("order", "percent")
	by := (*c).DefaultQuery("by", "asc")
	tanggal := (*c).DefaultQuery("tanggal", "null")
	loc, _ := time.LoadLocation("Asia/Jakarta")
	tgl, _ := time.ParseInLocation("2006-01", tanggal, loc)

	var rekap []JumlahRekapSiswa

	paginator := pagination.Paging(&pagination.Param{
		DB:      database.DB.Table("absen_children").Select("SUM(absen_children.kode='h') / COUNT(absen_children.kode) * 100.0 AS percent,absen_children.nama,absen_children.jenis,absen_children.foto,absens.kelas,SUM(absen_children.kode='h') AS hadir,SUM(absen_children.kode='i') AS izin,SUM(absen_children.kode='s') AS sakit,SUM(absen_children.kode='t') AS terlambat,SUM(absen_children.kode='a') AS alpa").Joins("left join absens on absens.id = absen_children.absen_id").Where("YEAR(absens.tanggal) = ?", tgl.Year()).Where("MONTH(absens.tanggal) = ?", tgl.Month()).Where("absen_children.nama LIKE ?", "%"+search+"%").Group("absen_children.siswa_id,MONTH(absens.tanggal)"),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{order + " " + by},
		ShowSQL: false,
	}, &rekap)

	(*c).JSON(200, paginator)

}

func (a *Rekap) Kelas(c *gin.Context) {

	page, _ := strconv.Atoi((*c).DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi((*c).DefaultQuery("limit", "3"))
	order := (*c).DefaultQuery("order", "percent")
	by := (*c).DefaultQuery("by", "asc")
	tanggal := (*c).DefaultQuery("tanggal", "null")
	loc, _ := time.LoadLocation("Asia/Jakarta")
	tgl, _ := time.ParseInLocation("2006-01", tanggal, loc)

	var rekap []JumlahRekapKelas

	paginator := pagination.Paging(&pagination.Param{
		DB:      database.DB.Table("absen_children").Select("SUM(absen_children.kode='h') / COUNT(absen_children.kode) * 100.0 AS percent,absens.kelas,absens.kelas_id,SUM(absen_children.kode='h') AS hadir,SUM(absen_children.kode='i') AS izin,SUM(absen_children.kode='s') AS sakit,SUM(absen_children.kode='t') AS terlambat,SUM(absen_children.kode='a') AS alpa").Joins("left join absens on absens.id = absen_children.absen_id").Where("YEAR(absens.tanggal) = ?", tgl.Year()).Where("MONTH(absens.tanggal) = ?", tgl.Month()).Group("absens.kelas_id,MONTH(absens.tanggal)"),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{order + " " + by},
		ShowSQL: false,
	}, &rekap)

	(*c).JSON(200, paginator)

}

func (a *Rekap) Guru(c *gin.Context) {

	page, _ := strconv.Atoi((*c).DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi((*c).DefaultQuery("limit", "3"))
	search := (*c).DefaultQuery("search", "%")
	order := (*c).DefaultQuery("order", "percent")
	by := (*c).DefaultQuery("by", "asc")
	tanggal := (*c).DefaultQuery("tanggal", "null")
	loc, _ := time.LoadLocation("Asia/Jakarta")
	tgl, _ := time.ParseInLocation("2006-01", tanggal, loc)

	var rekap []JumlahRekapGuru

	paginator := pagination.Paging(&pagination.Param{
		DB:      database.DB.Table("absen_children").Select("SUM(absen_children.kode='h') / COUNT(absen_children.kode) * 100.0 AS percent,absens.mapel_guru AS nama,absens.mapel,absens.user_id,SUM(absen_children.kode='h') AS hadir,SUM(absen_children.kode='i') AS izin,SUM(absen_children.kode='s') AS sakit,SUM(absen_children.kode='t') AS terlambat,SUM(absen_children.kode='a') AS alpa").Joins("left join absens on absens.id = absen_children.absen_id").Where("YEAR(absens.tanggal) = ?", tgl.Year()).Where("MONTH(absens.tanggal) = ?", tgl.Month()).Where("absens.mapel_guru LIKE ?", "%"+search+"%").Group("absens.user_id,MONTH(absens.tanggal)"),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{order + " " + by},
		ShowSQL: false,
	}, &rekap)

	(*c).JSON(200, paginator)

}

type JumlahRekapGuru struct {
	Nama      string `json:"nama,omitempty"`
	Mapel     string `json:"mapel,omitempty"`
	UserID    uint   `json:"user_id,omitempty"`
	Hadir     uint   `json:"hadir"`
	Izin      uint   `json:"izin"`
	Sakit     uint   `json:"sakit"`
	Terlambat uint   `json:"terlambat"`
	Alpa      uint   `json:"alpa"`
}

type JumlahRekapKelas struct {
	Kelas     string `json:"kelas,omitempty"`
	KelasID   uint   `json:"kelas_id,omitempty"`
	Hadir     uint   `json:"hadir"`
	Izin      uint   `json:"izin"`
	Sakit     uint   `json:"sakit"`
	Terlambat uint   `json:"terlambat"`
	Alpa      uint   `json:"alpa"`
}

type JumlahRekapSiswa struct {
	Nama      string `json:"nama,omitempty"`
	Jenis     string `json:"jenis,omitempty"`
	Foto      string `json:"foto,omitempty"`
	Kelas     string `json:"kelas,omitempty"`
	Hadir     uint   `json:"hadir"`
	Izin      uint   `json:"izin"`
	Sakit     uint   `json:"sakit"`
	Terlambat uint   `json:"terlambat"`
	Alpa      uint   `json:"alpa"`
}
