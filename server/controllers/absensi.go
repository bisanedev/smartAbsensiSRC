package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"../database"
	"../models"
	"github.com/biezhi/gorm-paginator/pagination"
)

type Absensi struct {
	Basic
}

func (a *Absensi) Index(c *gin.Context) {

	tanggal := (*c).DefaultQuery("tanggal", "null")
	order := (*c).DefaultQuery("order", "id")
	by := (*c).DefaultQuery("by", "desc")
	search := (*c).DefaultQuery("search", "%")
	page, _ := strconv.Atoi((*c).DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi((*c).DefaultQuery("limit", "3"))

	var absen []models.Absen

	if tanggal == "null" {
		// tampilkan semuanya
		paginator := pagination.Paging(&pagination.Param{
			DB:      database.DB.Where("mapel_guru LIKE ?", "%"+search+"%"),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{order + " " + by + ", jam DESC"},
			ShowSQL: false,
		}, &absen)
		(*c).JSON(200, paginator)
	} else {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		tgl, _ := time.ParseInLocation("2006-01-02", tanggal, loc)
		paginator := pagination.Paging(&pagination.Param{
			DB:      database.DB.Where("DATE(tanggal) = ?", tgl).Where("mapel_guru LIKE ?", "%"+search+"%"),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{order + " " + by + ", jam DESC"},
			ShowSQL: false,
		}, &absen)
		(*c).JSON(200, paginator)
	}
}

type AbsenHarian struct {
	Tanggal  time.Time `json:"tanggal,omitempty"`
	Semester string    `json:"semester,omitempty"`
	SiswaID  uint      `json:"siswa_id,omitempty"`
	Nama     string    `json:"nama,omitempty"`
	Jenis    string    `json:"jenis,omitempty"`
	Kelas    string    `json:"kelas,omitempty"`
	Kode     string    `json:"kode,omitempty"`
}

func (a *Absensi) Harian(c *gin.Context) {

	tanggal := (*c).DefaultQuery("tanggal", "null")
	order := (*c).DefaultQuery("order", "absens.tanggal")
	by := (*c).DefaultQuery("by", "desc")
	search := (*c).DefaultQuery("search", "%")
	page, _ := strconv.Atoi((*c).DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi((*c).DefaultQuery("limit", "3"))

	var absen []AbsenHarian

	if tanggal == "null" {
		// tampilkan semuanya
		paginator := pagination.Paging(&pagination.Param{
			DB:      database.DB.Table("absen_children").Select("absens.semester,absens.tanggal,absen_children.siswa_id,absen_children.nama,absen_children.jenis,absens.kelas,GROUP_CONCAT(absen_children.kode ORDER BY absens.jam ASC) AS kode").Joins("left join absens on absens.id = absen_children.absen_id").Where("absen_children.nama LIKE ?", "%"+search+"%").Group("absen_children.siswa_id,date(absens.tanggal)"),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{order + " " + by},
			ShowSQL: false,
		}, &absen)
		(*c).JSON(200, paginator)
	} else {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		tgl, _ := time.ParseInLocation("2006-01-02", tanggal, loc)
		paginator := pagination.Paging(&pagination.Param{
			DB:      database.DB.Table("absen_children").Select("absens.semester,absens.tanggal,absen_children.siswa_id,absen_children.nama,absen_children.jenis,absens.kelas,GROUP_CONCAT(absen_children.kode ORDER BY absens.jam ASC) AS kode").Joins("left join absens on absens.id = absen_children.absen_id").Where("date(absens.tanggal) = ?", tgl).Where("absen_children.nama LIKE ?", "%"+search+"%").Group("absen_children.siswa_id,date(absens.tanggal)"),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{order + " " + by},
			ShowSQL: false,
		}, &absen)
		(*c).JSON(200, paginator)
	}
}

func (a *Absensi) IndexMobile(c *gin.Context) {

	guru := (*c).DefaultQuery("guru", "null")
	tanggal := (*c).DefaultQuery("tanggal", "null")

	if tanggal == "null" || guru == "null" {
		(*a).JsonFail(c, http.StatusBadRequest, "Parameter Query Undefined")
	} else {
		var absen []models.Absen
		loc, _ := time.LoadLocation("Asia/Jakarta")
		tgl, _ := time.ParseInLocation("2006-01-02", tanggal, loc)

		database.DB.Set("gorm:auto_preload", true).Where("date(tanggal) = ?", tgl).Where("user_id = ?", guru).First(&absen)
		(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": absen})
	}

}

func (a *Absensi) ListTanggal(c *gin.Context) {
	guru := (*c).DefaultQuery("guru", "null")

	if guru == "null" {
		(*a).JsonFail(c, http.StatusBadRequest, "Parameter Query Undefined")
	} else {

		type Result struct {
			Tgl string
		}

		var result []Result

		database.DB.Raw("SELECT DATE_FORMAT(tanggal, '%Y-%m-%d') AS tgl FROM absens WHERE tanggal > DATE_SUB(now(), INTERVAL 7 MONTH) AND user_id = ?", guru).Scan(&result)

		(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": result})
	}

}

func (a *Absensi) CekMapel(c *gin.Context) {
	kelasString := (*c).DefaultQuery("kelas", "null")
	tanggal := (*c).DefaultQuery("tanggal", "0")

	loc, _ := time.LoadLocation("Asia/Jakarta")
	tgl, _ := time.ParseInLocation("2006-01-02", tanggal, loc)

	var absen []models.Absen
	var cekmapel []int64

	if tanggal != "0" || kelasString != "null" {
		kelas, _ := strconv.Atoi(kelasString)

		if err := database.DB.Where("kelas_id = ?", kelas).Where("DATE(tanggal) = ?", tgl).Order("jam asc").Find(&absen).Pluck("jam", &cekmapel).Error; err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}

		(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": cekmapel})

	}
}

func CekSemesterID(tanggal string) uint {
	var semester models.Semester
	loc, _ := time.LoadLocation("Asia/Jakarta")
	tgl, _ := time.ParseInLocation("2006-01-02T15:04", tanggal, loc)
	database.DB.Where("waktustart <= ?", tgl).Where("waktuend >= ?", tgl).First(&semester)
	return semester.ID
}

func CekSemesterSTR(tanggal string) string {
	var semester models.Semester
	loc, _ := time.LoadLocation("Asia/Jakarta")
	tgl, _ := time.ParseInLocation("2006-01-02T15:04", tanggal, loc)
	database.DB.Where("waktustart <= ?", tgl).Where("waktuend >= ?", tgl).First(&semester)
	return semester.Semester
}

func CekSemesterMulti(tanggal string) (string, string) {
	var semester models.Semester
	loc, _ := time.LoadLocation("Asia/Jakarta")
	tgl, _ := time.ParseInLocation("2006-01-02T15:04", tanggal, loc)
	database.DB.Where("waktustart <= ?", tgl).Where("waktuend >= ?", tgl).First(&semester)
	return semester.Semester, semester.Nama
}

func inTimeSpan(start, check time.Time) bool {
	//return check.After(start) && check.Before(end)
	kurangsemenit := start.Add(-1 * time.Minute)
	return check.After(kurangsemenit)
}

func (a *Absensi) Store(c *gin.Context) {
	var request RequestAbsen
	if err := (*c).ShouldBind(&request); err == nil {
		var count int
		var siswa []models.AbsenChild
		loc, _ := time.LoadLocation("Asia/Jakarta")
		tgl, _ := time.ParseInLocation("2006-01-02T15:04", request.Tanggal, loc)
		keluar, _ := time.ParseInLocation("2006-01-02T15:04", request.Keluar, loc)
		mapelStart, _ := time.ParseInLocation("2006-01-02T15:04", request.MapelStart, loc)
		mapelEnd, _ := time.ParseInLocation("2006-01-02T15:04", request.MapelEnd, loc)

		if !inTimeSpan(mapelStart, tgl) {
			(*a).JsonFail(c, http.StatusBadRequest, "Dilarang Absen Sebelum Waktu-nya")
			return
		}
		now := time.Now()
		if tgl.Unix() > now.Unix() {
			(*a).JsonFail(c, http.StatusBadRequest, "Dilarang Absen Sebelum Waktu-nya")
			return
		}
		database.DB.Model(&models.Absen{}).Where("kelas_id = ?", request.KelasID).Where("date(tanggal) = ?", tgl.Format("2006-01-02")).Where("jam = ?", request.Jam).Count(&count)

		if count == 0 {
			//jika data tidak ada
			//getsiswa and save table default Kode Hadir
			database.DB.Table("siswas").Where("kelas_id = ?", request.KelasID).Where("status = ?", "a").Select("id as siswa_id,urut, nama, jenis, foto , 'h' as kode").Order("urut asc").Scan(&siswa)
			//store
			absen := &models.Absen{
				SemesterID: CekSemesterID(request.Tanggal),
				KelasID:    request.KelasID,
				UserID:     request.UserID,
				Tanggal:    tgl,
				Keluar:     keluar,
				Jam:        request.Jam,
				Semester:   CekSemesterSTR(request.Tanggal),
				Kelas:      request.Kelas,
				Mapel:      request.Mapel,
				MapelGuru:  request.MapelGuru,
				MapelStart: mapelStart,
				MapelEnd:   mapelEnd,
				PiketID:    request.PiketID,
				PiketGuru:  request.PiketGuru,
				Siswa:      siswa,
			}

			if err := database.DB.Create(&absen).Error; err != nil {
				(*a).JsonFail(c, http.StatusBadRequest, err.Error())
				return
			}
			//send message to broadcast a channel

			Absen := CommandWS{Name: "guru", Data: request.MapelGuru}
			a.SocketSend("absen", Absen)
			a.SocketSend("absenupdate", Absen)

			(*a).JsonSuccess(c, http.StatusCreated, gin.H{"message": "Created successfully"})

		}
	} else {
		//(*a).JsonFail(c, http.StatusBadRequest, "Silahkan Periksa Kembali Input Data !")
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
	}
}

func (a *Absensi) Update(c *gin.Context) {
	var request UpdateAbsen
	if err := (*c).ShouldBind(&request); err == nil {
		var absen models.Absen

		if database.DB.First(&absen, (*c).Param("id")).Error != nil {
			(*a).JsonFail(c, http.StatusNotFound, "Data does not exist")
			return
		}
		loc, _ := time.LoadLocation("Asia/Jakarta")
		tgl, _ := time.ParseInLocation("2006-01-02T15:04", request.Tanggal, loc)
		keluar, _ := time.ParseInLocation("2006-01-02T15:04", request.Keluar, loc)
		mapelStart, _ := time.ParseInLocation("2006-01-02T15:04", request.MapelStart, loc)
		mapelEnd, _ := time.ParseInLocation("2006-01-02T15:04", request.MapelEnd, loc)

		if !inTimeSpan(mapelStart, tgl) {
			(*a).JsonFail(c, http.StatusBadRequest, "Dilarang Absen Sebelum Waktu-nya")
			return
		}

		absen.SemesterID = CekSemesterID(request.Tanggal)

		absen.UserID = request.UserID
		absen.Tanggal = tgl
		absen.Keluar = keluar
		absen.Jam = request.Jam
		absen.Semester = CekSemesterSTR(request.Tanggal)
		absen.Mapel = request.Mapel
		absen.Kelas = request.Kelas
		absen.MapelGuru = request.MapelGuru
		absen.MapelStart = mapelStart
		absen.MapelEnd = mapelEnd
		absen.PiketID = request.PiketID
		absen.PiketGuru = request.PiketGuru

		if err := database.DB.Save(&absen).Error; err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}

		(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
	} else {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
	}
}

func (a *Absensi) UpdateJamKeluar(c *gin.Context) {
	var request RequestAbsenMobileJamKeluar
	if err := (*c).ShouldBind(&request); err == nil {
		//cek absen sudah apa tidak
		var count int
		loc, _ := time.LoadLocation("Asia/Jakarta")
		keluar, _ := time.ParseInLocation("2006-01-02T15:04", request.Keluar, loc)
		var absenUpdate models.Absen

		database.DB.Model(&models.Absen{}).Where("kelas_id = ?", request.KelasID).Where("date(tanggal) = ?", keluar.Format("2006-01-02")).Where("jam = ?", request.Jam).Where("keluar = '0000-00-00 00:00:00'").First(&absenUpdate).Count(&count)

		if count > 0 {
			//jika data sudah ada / update

			absenUpdate.Keluar = keluar

			if err := database.DB.Save(&absenUpdate).Error; err != nil {
				(*a).JsonFail(c, http.StatusBadRequest, err.Error())
				return
			}

			Absen := CommandWS{Name: "update", Data: "~"}
			a.SocketSend("absenupdate", Absen)

			(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
		} else {
			(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
		}
		// end cek attribut
	} else {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
	}
}

func (a *Absensi) StoreAbsen(c *gin.Context) {
	var request RequestAbsenMobile
	if err := (*c).ShouldBind(&request); err == nil {
		//cek absen sudah apa tidak
		var count int
		loc, _ := time.LoadLocation("Asia/Jakarta")
		tgl, _ := time.ParseInLocation("2006-01-02T15:04", request.Tanggal, loc)
		var absenUpdate models.Absen

		database.DB.Model(&models.Absen{}).Where("kelas_id = ?", request.KelasID).Where("date(tanggal) = ?", tgl.Format("2006-01-02")).Where("jam = ?", request.Jam).First(&absenUpdate).Count(&count)

		if count > 0 {
			//jika data sudah ada / update
			var siswa []models.AbsenChild
			json.Unmarshal([]byte(request.Siswa), &siswa)

			absenUpdate.Siswa = siswa

			if err := database.DB.Save(&absenUpdate).Error; err != nil {
				(*a).JsonFail(c, http.StatusBadRequest, err.Error())
				return
			}

			Absen := CommandWS{Name: "update", Data: "~"}
			a.SocketSend("absenupdate", Absen)

			(*a).JsonSuccess(c, http.StatusCreated, gin.H{"tanggal": tgl})
		}
		// end cek attribut
	} else {
		//(*a).JsonFail(c, http.StatusBadRequest, "Silahkan Periksa Kembali Input Data !")
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
	}
	//----
}

// get absensi (edit) berdasarkan siswa
func (a *Absensi) GetAbsensi(c *gin.Context) {
	tanggal := (*c).DefaultQuery("tanggal", "null")
	siswa := (*c).DefaultQuery("siswa", "null")

	if siswa == "null" || tanggal == "null" {
		(*a).JsonFail(c, http.StatusBadRequest, "Parameter Query Undefined")
	} else {
		var absensiswa []AbsenPerSiswa
		loc, _ := time.LoadLocation("Asia/Jakarta")
		tgl, _ := time.ParseInLocation("2006-01-02", tanggal, loc)

		database.DB.Table("absen_children").Select("absen_children.id,absens.jam,absens.mapel,absens.mapel_start,absens.mapel_end,absens.mapel_guru,absen_children.nama,absen_children.jenis,absen_children.foto,absen_children.urut,absen_children.kode,absen_children.kode_note").Joins("left join absens on absens.id = absen_children.absen_id").Where("DATE(absens.tanggal) = ?", tgl.Format("2006-01-02")).Where("absen_children.siswa_id = ?", siswa).Scan(&absensiswa)
		(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": absensiswa})
	}
}

func (a *Absensi) PostAbsensi(c *gin.Context) {
	var request RequestAbsenJsonPost

	if err := (*c).ShouldBind(&request); err == nil {
		var siswa []AbsenPerSiswa
		json.Unmarshal([]byte(request.Absen), &siswa)

		for _, elem := range siswa {
			database.DB.Exec("UPDATE absen_children SET kode=? , kode_note=? WHERE id=?", elem.Kode, elem.KodeNote, elem.ID)
		}

		(*a).JsonSuccess(c, http.StatusCreated, gin.H{})

	} else {
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
	}
}

// absensi edit / update berdasarkan kelas
func (a *Absensi) Siswa(c *gin.Context) {

	kelas := (*c).DefaultQuery("kelas", "null")
	tanggal := (*c).DefaultQuery("tanggal", "null")
	jam := (*c).DefaultQuery("jam", "null")

	if kelas == "null" || tanggal == "null" || jam == "null" {
		(*a).JsonFail(c, http.StatusBadRequest, "Parameter Query Undefined")
	} else {
		var absen models.Absen
		loc, _ := time.LoadLocation("Asia/Jakarta")
		tgl, _ := time.ParseInLocation("2006-01-02T15:04", tanggal, loc)

		if database.DB.Set("gorm:auto_preload", true).Where("kelas_id = ?", kelas).Where("date(tanggal) = ?", tgl.Format("2006-01-02")).Where("jam = ?", jam).First(&absen).RecordNotFound() {

			(*a).JsonFail(c, http.StatusBadRequest, "Data Tidak Di Temukan")

		} else {
			// ----------------
			(*a).JsonSuccess(c, http.StatusOK, gin.H{"siswa": absen.Siswa, "updated": "yes", "tanggal": tgl})
		}
	}

}

func (a *Absensi) Destroy(c *gin.Context) {
	//role akses
	if (*a).Role(c, "admin") {
		//----------
		var absen models.Absen

		if database.DB.First(&absen, (*c).Param("id")).Error != nil {
			(*a).JsonFail(c, http.StatusNotFound, "Data Not Found")
			return
		}

		// delete absen
		if err := database.DB.Unscoped().Delete(&absen).Error; err != nil {
			(*a).JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}
		(*a).JsonSuccess(c, http.StatusCreated, gin.H{})
		//else role
	} else {
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
}

func (a *Absensi) MultiDelete(c *gin.Context) {
	if (*a).Role(c, "admin") {
		// roleakses
		x := &MultiHapus{}

		if err := (*c).ShouldBind(&x); err == nil {

			database.DB.Exec("DELETE from absens WHERE id IN (?)", x.Selected)

			(*a).JsonSuccess(c, http.StatusCreated, gin.H{})

		} else {
			(*a).JsonFail(c, http.StatusBadRequest, "Silahkan Periksa Kembali Input Data !")
		}
		//else role
	} else {
		(*a).JsonFail(c, http.StatusUnauthorized, "Access Prohibited")
	}
}

type RequestAbsenMobileJamKeluar struct {
	Keluar  string `form:"keluar" json:"keluar" binding:"required"`
	KelasID uint   `form:"kelas_id" json:"kelas_id" binding:"required"`
	Jam     uint   `form:"jam" json:"jam" binding:"required"`
}

type RequestAbsenJsonPost struct {
	Absen string `form:"absen" json:"absen" binding:"required"`
}
type AbsenPerSiswa struct {
	ID         uint   `json:"id,omitempty"`
	Mapel      string `json:"mapel,omitempty"`
	MapelGuru  string `json:"mapel_guru,omitempty"`
	MapelStart string `json:"mapel_start,omitempty"`
	MapelEnd   string `json:"mapel_end,omitempty"`
	Urut       uint   `json:"urut,omitempty"`
	Jam        uint   `json:"jam,omitempty"`
	Nama       string `json:"nama,omitempty"`
	Jenis      string `json:"jenis,omitempty"`
	Foto       string `json:"foto,omitempty"`
	Kode       string `json:"kode,omitempty"`
	KodeNote   string `json:"kode_note,omitempty"`
}
type UpdateAbsen struct {
	UserID     uint   `form:"user_id" json:"user_id" binding:"required"`
	Tanggal    string `form:"tanggal" json:"tanggal" binding:"required"`
	Keluar     string `form:"keluar" json:"keluar" binding:"required"`
	Jam        uint   `form:"jam" json:"jam" binding:"required"`
	Kelas      string `form:"kelas" json:"kelas" binding:"required"`
	Mapel      string `form:"mapel" json:"mapel" binding:"required"`
	MapelGuru  string `form:"mapelguru" json:"mapelguru" binding:"required"`
	MapelStart string `form:"mapel_start" json:"mapel_start" binding:"required"`
	MapelEnd   string `form:"mapel_end" json:"mapel_end" binding:"required"`
	PiketID    uint   `form:"piketid" json:"piketid"`
	PiketGuru  string `form:"piketguru" json:"piketguru"`
}

type RequestAbsen struct {
	KelasID    uint   `form:"kelas_id" json:"kelas_id" binding:"required"`
	UserID     uint   `form:"user_id" json:"user_id" binding:"required"`
	Tanggal    string `form:"tanggal" json:"tanggal" binding:"required"`
	Keluar     string `form:"keluar" json:"keluar"`
	Jam        uint   `form:"jam" json:"jam" binding:"required"`
	Kelas      string `form:"kelas" json:"kelas" binding:"required"`
	Mapel      string `form:"mapel" json:"mapel" binding:"required"`
	MapelGuru  string `form:"mapelguru" json:"mapelguru" binding:"required"`
	MapelStart string `form:"mapel_start" json:"mapel_start" binding:"required"`
	MapelEnd   string `form:"mapel_end" json:"mapel_end" binding:"required"`
	PiketID    uint   `form:"piketid" json:"piketid"`
	PiketGuru  string `form:"piketguru" json:"piketguru"`
}

type RequestAbsenMobile struct {
	KelasID   uint   `form:"kelas_id" json:"kelas_id" binding:"required"`
	UserID    uint   `form:"user_id" json:"user_id" binding:"required"`
	Tanggal   string `form:"tanggal" json:"tanggal" binding:"required"`
	Jam       uint   `form:"jam" json:"jam" binding:"required"`
	Kelas     string `form:"kelas" json:"kelas" binding:"required"`
	Mapel     string `form:"mapel" json:"mapel" binding:"required"`
	MapelGuru string `form:"mapelguru" json:"mapelguru" binding:"required"`
	PiketID   uint   `form:"piketid" json:"piketid"`
	PiketGuru string `form:"piketguru" json:"piketguru"`
	Siswa     string `form:"siswa" json:"siswa" binding:"required"`
}

type AbsenSiswa struct {
	SiswaID  uint   `json:"siswa_id,omitempty"`
	Urut     uint   `json:"urut,omitempty"`
	Nama     string `json:"nama,omitempty"`
	Jenis    string `json:"jenis,omitempty"`
	Foto     string `json:"foto"`
	Kode     string `json:"kode"`
	KodeNote string `json:"kodenote"`
}
