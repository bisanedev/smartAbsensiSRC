package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"../database"
	"../models"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

type Laporan struct {
	Basic
}

func (a *Laporan) AbsensiKelas(c *gin.Context) {
	tanggal := (*c).DefaultQuery("tanggal", "null")
	kelas := (*c).DefaultQuery("kelas", "0")
	loc, _ := time.LoadLocation("Asia/Jakarta")
	tgl, _ := time.ParseInLocation("2006-01", tanggal, loc)

	var laporan []AbsenBulananKelas

	database.DB.Table("absen_children").Select("absen_children.siswa_id,absen_children.urut,absen_children.nama,DAY(absens.tanggal) AS tgl,GROUP_CONCAT(absen_children.kode) AS kode").Joins("left join absens on absens.id = absen_children.absen_id").Where("YEAR(absens.tanggal) = ?", tgl.Year()).Where("MONTH(absens.tanggal) = ?", tgl.Month()).Where("absens.kelas_id = ?", kelas).Group("absen_children.siswa_id,DAY(absens.tanggal)").Order("absen_children.urut").Scan(&laporan)

	(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": laporan})

}

func (a *Laporan) AbsensiKelasExcel(c *gin.Context) {
	var request AbsenKelasExcelRequest
	if err := (*c).ShouldBind(&request); err == nil {
		var absen []RequestAbsenBulananKelas
		json.Unmarshal([]byte(request.Data), &absen)

		f := excelize.NewFile()
		f.SetSheetName("Sheet1", request.Bulan)
		if err := f.SetPageLayout(request.Bulan, excelize.PageLayoutOrientation(excelize.OrientationLandscape)); err != nil {
			panic(err)
		}
		if err := f.SetPageLayout(request.Bulan, excelize.PageLayoutPaperSize(14)); err != nil {
			panic(err)
		}
		f.MergeCell(request.Bulan, "A2", "A4")
		f.MergeCell(request.Bulan, "B2", "D4")
		f.MergeCell(request.Bulan, "E2", "AI3")
		f.MergeCell(request.Bulan, "AJ2", "AO3")
		f.SetColWidth(request.Bulan, "A", "A", 3.00)
		f.SetColWidth(request.Bulan, "E", "AI", 2.70)
		f.SetColWidth(request.Bulan, "AJ", "AO", 3.90)
		//set value
		f.SetCellValue(request.Bulan, "A2", "No")
		f.SetCellValue(request.Bulan, "B2", "Nama Siswa")
		f.SetCellValue(request.Bulan, "E2", "HADIR[H], SAKIT[S], IZIN[I], ALPA[A], Tidak Diketahui[!]")
		f.SetCellValue(request.Bulan, "AJ2", "Jumlah")
		f.SetCellValue(request.Bulan, "AJ4", "S")
		f.SetCellValue(request.Bulan, "AK4", "I")
		f.SetCellValue(request.Bulan, "AL4", "A")
		f.SetCellValue(request.Bulan, "AM4", "H")
		f.SetCellValue(request.Bulan, "AN4", "!")
		f.SetCellValue(request.Bulan, "AO4", "%")
		tglHeader := map[string]int{"E4": 1, "F4": 2, "G4": 3, "H4": 4, "I4": 5, "J4": 6, "K4": 7, "L4": 8, "M4": 9, "N4": 10, "O4": 11, "P4": 12, "Q4": 13, "R4": 14, "S4": 15, "T4": 16, "U4": 17, "V4": 18, "W4": 19, "X4": 20, "Y4": 21, "Z4": 22, "AA4": 23, "AB4": 24, "AC4": 25, "AD4": 26, "AE4": 27, "AF4": 28, "AG4": 29, "AH4": 30, "AI4": 31}
		for k, v := range tglHeader {
			f.SetCellValue(request.Bulan, k, v)
		}
		//style
		styleHeader, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#4472c4"],"pattern":1},"font":{"bold":true,"family":"Calibri","size":11,"color":"FFFFFF"},"alignment":{"horizontal":"center","vertical":"center"}}`)
		styleHeader2, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#d6dce4"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"000000"},"alignment":{"horizontal":"center","vertical":"center"}}`)
		f.SetCellStyle(request.Bulan, "A2", "AO4", styleHeader)
		f.SetCellStyle(request.Bulan, "E4", "AO4", styleHeader2)
		// content
		col := []string{"A", "B", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO"}

		for k, v := range absen {
			urut := col[0] + strconv.Itoa(k+5)
			nama := col[1] + strconv.Itoa(k+5)
			//merge nama cel
			f.MergeCell(request.Bulan, col[1]+strconv.Itoa(k+5), col[2]+strconv.Itoa(k+5))
			//set value urut dan nama
			f.SetCellValue(request.Bulan, urut, v.Urut)
			f.SetCellValue(request.Bulan, nama, strings.Title(strings.ToLower(v.Nama)))
			//style urut & nama
			styleUrut, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#ffffff"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"000000"},"alignment":{"horizontal":"center","vertical":"center"}}`)
			styleNama, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#d6dce4"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"000000"},"alignment":{"horizontal":"left","vertical":"center"}}`)
			f.SetCellStyle(request.Bulan, urut, urut, styleUrut)
			f.SetCellStyle(request.Bulan, nama, col[2]+strconv.Itoa(k+5), styleNama)
			//style border tanggal
			borderTgl, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#ffffff"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"00000"},"alignment":{"horizontal":"center","vertical":"center"}}`)
			f.SetCellStyle(request.Bulan, col[3]+strconv.Itoa(k+5), "AI"+strconv.Itoa(k+5), borderTgl)
			//loop tgl
			for i := 1; i < 32; i++ {
				//absen tanggal
				if contains(v.Tgl, strconv.Itoa(i)) {
					tgl := col[i+2] + strconv.Itoa(k+5)
					//split kode
					absen := kodecheck(v.Tgl, v.Kode, strconv.Itoa(i))
					f.SetCellValue(request.Bulan, tgl, absen)
					//style tanggal dan kode
					KodeHadir, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#008000"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"FFFFF"},"alignment":{"horizontal":"center","vertical":"center"}}`)
					KodeSakit, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#07ade3"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"FFFFF"},"alignment":{"horizontal":"center","vertical":"center"}}`)
					KodeIzin, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#0716e3"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"FFFFF"},"alignment":{"horizontal":"center","vertical":"center"}}`)
					KodeAlpa, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#e30707"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"FFFFF"},"alignment":{"horizontal":"center","vertical":"center"}}`)
					KodeTidak, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#fc9403"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"FFFFF"},"alignment":{"horizontal":"center","vertical":"center"}}`)
					if absen == "H" {
						f.SetCellStyle(request.Bulan, tgl, tgl, KodeHadir)
					}
					if absen == "S" {
						f.SetCellStyle(request.Bulan, tgl, tgl, KodeSakit)
					}
					if absen == "I" {
						f.SetCellStyle(request.Bulan, tgl, tgl, KodeIzin)
					}
					if absen == "A" {
						f.SetCellStyle(request.Bulan, tgl, tgl, KodeAlpa)
					}
					if absen == "!" {
						f.SetCellStyle(request.Bulan, tgl, tgl, KodeTidak)
					}
				}
			}
			//jumlah absen
			sakit := "AJ" + strconv.Itoa(k+5)
			izin := "AK" + strconv.Itoa(k+5)
			alpa := "AL" + strconv.Itoa(k+5)
			hadir := "AM" + strconv.Itoa(k+5)
			seru := "AN" + strconv.Itoa(k+5)
			percen := "AO" + strconv.Itoa(k+5)
			rangestart := "E" + strconv.Itoa(k+5)
			rangeend := "AI" + strconv.Itoa(k+5)
			f.SetCellFormula(request.Bulan, sakit, "COUNTIF("+rangestart+":"+rangeend+",\"S\")")
			f.SetCellFormula(request.Bulan, izin, "COUNTIF("+rangestart+":"+rangeend+",\"I\")")
			f.SetCellFormula(request.Bulan, alpa, "COUNTIF("+rangestart+":"+rangeend+",\"A\")")
			f.SetCellFormula(request.Bulan, hadir, "COUNTIF("+rangestart+":"+rangeend+",\"H\")")
			f.SetCellFormula(request.Bulan, seru, "COUNTIF("+rangestart+":"+rangeend+",\"!\")")
			f.SetCellFormula(request.Bulan, percen, "("+hadir+"/COUNTIF("+rangestart+":"+rangeend+",\"*\"))*100")
			styleJumlah, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#a9d08e"],"pattern":1},"font":{"bold":true,"family":"Calibri","size":9,"color":"000000"},"alignment":{"horizontal":"center","vertical":"center"}}`)
			f.SetCellStyle(request.Bulan, sakit, percen, styleJumlah)
		}
		// Save xlsx file by the given path.
		loc, _ := time.LoadLocation("Asia/Jakarta")
		tgl, _ := time.ParseInLocation("2006-01", request.Tgl, loc)

		path := "storage/excel/kelas[" + request.Kelas + "]_" + request.Bulan + "_" + strconv.Itoa(tgl.Year()) + ".xlsx"
		if err := f.SaveAs(path); err == nil {
			(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": path})
		}

	} else {
		//(*a).JsonFail(c, http.StatusBadRequest, "Silahkan Periksa Kembali Input Data !")
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
	}

}

func kodecheck(tgl []string, kode []string, tanggal string) string {
	for k, v := range tgl {
		if v == tanggal {
			absen := strings.Split(kode[k], ",")
			jumlah := len(absen)
			if jumlah == countAbsen(absen, "h", "t") {
				return "H"
			}
			if jumlah == countAbsen(absen, "s", "s") {
				return "S"
			}
			if jumlah == countAbsen(absen, "i", "i") {
				return "I"
			}
			if jumlah == countAbsen(absen, "a", "a") {
				return "A"
			}

		}
	}
	return "!"
}

func countAbsen(arr []string, conditinal string, atau string) int {
	sum := 0
	for _, v := range arr {
		if v == conditinal || v == atau {
			sum += 1
		}
	}
	return sum
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

//laporan attend
func (a *Laporan) AbsensiClock(c *gin.Context) {
	tanggal := (*c).DefaultQuery("tanggal", "null")
	loc, _ := time.LoadLocation("Asia/Jakarta")
	tgl, _ := time.ParseInLocation("2006-01", tanggal, loc)

	var laporan []AbsenBulananClock

	database.DB.Table("absens").Select("user_id,mapel_guru,DAY(tanggal) AS tgl,GROUP_CONCAT(TIME_FORMAT(tanggal,'%H:%i') ORDER BY tanggal,jam DESC) AS mulai,GROUP_CONCAT(TIME_FORMAT(keluar,'%H:%i') ORDER BY tanggal,jam DESC) AS keluar").Where("YEAR(absens.tanggal) = ?", tgl.Year()).Where("MONTH(tanggal) = ?", tgl.Month()).Where("piket_id = 0").Group("user_id,DAY(tanggal)").Order("mapel_guru,tanggal desc").Scan(&laporan)

	(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": laporan})
}

func (a *Laporan) AbsensiClockPiket(c *gin.Context) {
	tanggal := (*c).DefaultQuery("tanggal", "null")
	loc, _ := time.LoadLocation("Asia/Jakarta")
	tgl, _ := time.ParseInLocation("2006-01", tanggal, loc)

	var laporan []AbsenBulananClock

	database.DB.Table("absens").Select("piket_id AS user_id,piket_guru AS mapel_guru,DAY(tanggal) AS tgl,GROUP_CONCAT(TIME_FORMAT(tanggal,'%H:%i') ORDER BY tanggal,jam DESC) AS mulai,GROUP_CONCAT(TIME_FORMAT(keluar,'%H:%i') ORDER BY tanggal,jam DESC) AS keluar").Where("YEAR(absens.tanggal) = ?", tgl.Year()).Where("MONTH(tanggal) = ?", tgl.Month()).Where("piket_id != 0").Group("piket_id,DAY(tanggal)").Order("piket_guru,tanggal desc").Scan(&laporan)

	(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": laporan})
}

func (a *Laporan) AbsensiClockExcel(c *gin.Context) {
	var request ClockExcelRequest
	if err := (*c).ShouldBind(&request); err == nil {

		loc, _ := time.LoadLocation("Asia/Jakarta")
		tgl, _ := time.ParseInLocation("2006-01", request.Tanggal, loc)
		//get password excel
		var sistem models.Sistem
		database.DB.First(&sistem)
		//fmt.Println(sistem)
		//excel
		f := excelize.NewFile()
		f.SetSheetName("Sheet1", request.Bulan)
		f.ProtectSheet(request.Bulan, &excelize.FormatSheetProtection{
			Password:      sistem.PasswordClock,
			EditScenarios: false,
		})
		if err := f.SetPageLayout(request.Bulan, excelize.PageLayoutOrientation(excelize.OrientationLandscape)); err != nil {
			panic(err)
		}
		if err := f.SetPageLayout(request.Bulan, excelize.PageLayoutPaperSize(14)); err != nil {
			panic(err)
		}
		//header
		f.MergeCell(request.Bulan, "A2", "Z4")
		f.MergeCell(request.Bulan, "AA2", "AE3")
		f.MergeCell(request.Bulan, "AA4", "AE4")
		f.SetColWidth(request.Bulan, "A", "AE", 4.70)

		if !request.Piket {
			f.SetCellValue(request.Bulan, "A2", "TABEL KEHADIRAN GURU")
		} else {
			f.SetCellValue(request.Bulan, "A2", "TABEL KEHADIRAN GURU PIKET")
		}

		f.SetCellValue(request.Bulan, "AA2", strings.ToUpper(request.Bulan))
		f.SetCellValue(request.Bulan, "AA4", "TAHUN "+strconv.Itoa(tgl.Year()))
		styleHeader, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#FFFFFF"],"pattern":1},"font":{"bold":true,"family":"Calibri","size":10,"color":"#4472c4"},"alignment":{"horizontal":"center","vertical":"center"}}`)
		f.SetCellStyle(request.Bulan, "A2", "AE4", styleHeader)
		styleHeaderTitle, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#FFFFFF"],"pattern":1},"font":{"bold":true,"family":"Calibri","size":16,"color":"#4472c4"},"alignment":{"horizontal":"center","vertical":"center"}}`)
		f.SetCellStyle(request.Bulan, "A2", "Z4", styleHeaderTitle)
		f.SetCellStyle(request.Bulan, "AA2", "AE3", styleHeaderTitle)
		//body
		//data
		var laporan []RequestAbsenBulananClock
		json.Unmarshal([]byte(request.Data), &laporan)
		//loop
		col := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO"}

		for k, v := range laporan {
			jumlah := 12
			kali := (k * jumlah) + 5
			f.MergeCell(request.Bulan, "A"+strconv.Itoa(kali), "AE"+strconv.Itoa(kali))
			f.SetRowHeight(request.Bulan, kali, 30)
			f.SetCellValue(request.Bulan, "A"+strconv.Itoa(kali), v.MapelGuru)
			styleNama, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#4472c4"],"pattern":1},"font":{"bold":true,"family":"Calibri","size":11,"color":"FFFFFF"},"alignment":{"horizontal":"center","vertical":"center"}}`)
			f.SetCellStyle(request.Bulan, "A"+strconv.Itoa(kali), "AE"+strconv.Itoa(kali), styleNama)
			//tgl
			for i := 1; i < 32; i++ {
				tgl := col[i-1] + strconv.Itoa((k*jumlah)+6)
				f.SetCellValue(request.Bulan, tgl, i)
				//style tgl header
				styleTgl, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#d6dce4"],"pattern":1},"font":{"bold":true,"family":"Calibri","size":9,"color":"000000"},"alignment":{"horizontal":"center","vertical":"center"}}`)
				f.SetCellStyle(request.Bulan, tgl, tgl, styleTgl)
				//style null tgl
				styleTglNull, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#FFFFFF"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"000000"},"alignment":{"horizontal":"center","vertical":"center"}}`)
				f.SetCellStyle(request.Bulan, col[i-1]+strconv.Itoa((k*jumlah)+7), col[i-1]+strconv.Itoa((k*jumlah)+16), styleTglNull)
				//tgl value
				if contains(v.Tgl, strconv.Itoa(i)) {
					index := findIndexOf(v.Tgl, strconv.Itoa(i))
					mulaiarr := strings.Split(v.Mulai[index], ",")
					akhirarr := strings.Split(v.Keluar[index], ",")
					for key, val := range mulaiarr {
						//mulai
						f.SetCellValue(request.Bulan, col[i-1]+strconv.Itoa((k*jumlah)+7+(key*2)), val)
						styleMulai, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#e7e6e6"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"000000"},"alignment":{"horizontal":"center","vertical":"center"}}`)
						f.SetCellStyle(request.Bulan, col[i-1]+strconv.Itoa((k*jumlah)+7+(key*2)), col[i-1]+strconv.Itoa((k*jumlah)+7+(key*2)), styleMulai)
						//akhir
						f.SetCellValue(request.Bulan, col[i-1]+strconv.Itoa((k*jumlah)+8+(key*2)), akhirarr[key])
						styleAkhir, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#ffc000"],"pattern":1},"font":{"bold":false,"family":"Calibri","size":9,"color":"000000"},"alignment":{"horizontal":"center","vertical":"center"}}`)
						f.SetCellStyle(request.Bulan, col[i-1]+strconv.Itoa((k*jumlah)+8+(key*2)), col[i-1]+strconv.Itoa((k*jumlah)+8+(key*2)), styleAkhir)
					}
				}
			}
		}
		//save
		if !request.Piket {
			path := "storage/excel/clock_" + request.Bulan + "_" + strconv.Itoa(tgl.Year()) + ".xlsx"
			if err := f.SaveAs(path); err == nil {
				(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": path})
			}
		} else {
			path := "storage/excel/clock_piket_" + request.Bulan + "_" + strconv.Itoa(tgl.Year()) + ".xlsx"
			if err := f.SaveAs(path); err == nil {
				(*a).JsonSuccess(c, http.StatusOK, gin.H{"data": path})
			}
		}

		//--------------------------
	} else {
		//(*a).JsonFail(c, http.StatusBadRequest, "Silahkan Periksa Kembali Input Data !")
		(*a).JsonFail(c, http.StatusBadRequest, err.Error())
	}
}
func findIndexOf(arr []string, str string) int {
	index := 0
	for k, a := range arr {
		if a == str {
			index = k
		}
	}
	return index
}

//type
type AbsenBulananClock struct {
	UserID    uint   `json:"user_id,omitempty"`
	MapelGuru string `json:"mapel_guru,omitempty"`
	Tgl       string `json:"tgl,omitempty"`
	Mulai     string `json:"mulai,omitempty"`
	Keluar    string `json:"keluar,omitempty"`
}
type RequestAbsenBulananClock struct {
	UserID    uint     `json:"user_id,omitempty"`
	MapelGuru string   `json:"mapel_guru,omitempty"`
	Tgl       []string `json:"tgl,omitempty"`
	Mulai     []string `json:"mulai,omitempty"`
	Keluar    []string `json:"keluar,omitempty"`
}
type RequestAbsenBulananKelas struct {
	Urut    uint     `json:"urut,omitempty"`
	SiswaID uint     `json:"siswa_id,omitempty"`
	Nama    string   `json:"nama,omitempty"`
	Tgl     []string `json:"tgl,omitempty"`
	Kode    []string `json:"kode,omitempty"`
}
type AbsenBulananKelas struct {
	Urut    uint   `json:"urut,omitempty"`
	SiswaID uint   `json:"siswa_id,omitempty"`
	Nama    string `json:"nama,omitempty"`
	Tgl     string `json:"tgl,omitempty"`
	Kode    string `json:"kode,omitempty"`
}

type AbsenKelasExcelRequest struct {
	Data  string `form:"data" json:"data" binding:"required"`
	Bulan string `form:"bulan" json:"bulan" binding:"required"`
	Kelas string `form:"kelas" json:"kelas" binding:"required"`
	Tgl   string `form:"tgl" json:"tgl" binding:"required"`
}

type ClockExcelRequest struct {
	Bulan   string `form:"bulan" json:"bulan" binding:"required"`
	Tanggal string `form:"tanggal" json:"tanggal" binding:"required"`
	Data    string `form:"data" json:"data" binding:"required"`
	Piket   bool   `form:"piket" json:"piket"`
}
