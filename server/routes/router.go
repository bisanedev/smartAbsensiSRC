package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"gopkg.in/olahol/melody.v1"	
	"../controllers"
    // assetfs "github.com/elazarl/go-bindata-assetfs"      
    // "net/http"
    // "strings"
)

func InitRouter() *gin.Engine {	
	
	router := gin.New()	
	m := melody.New()
	//router.Use(gin.Logger())
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// relese mode
	gin.SetMode(gin.ReleaseMode)
	// index webapp
	//router.Use(static.Serve("/", BinaryFileSystem("public")))		
	router.Use(static.Serve("/", static.LocalFile("./public", true))) 	
	// auth	
	AuthJwt := new(controllers.AuthJwt)
	Users := new(controllers.Users)
	Info := new(controllers.Info)
	Sistem := new(controllers.Sistem)
	Kelas := new(controllers.Kelas)
	Siswa := new(controllers.Siswa)
	Semester := new(controllers.Semester)
	Jadwal := new(controllers.Jadwal)
	Absen := new(controllers.Absensi)
	Rekap := new(controllers.Rekap)
	Laporan := new(controllers.Laporan)
	Upload := new(controllers.Uploader)
	api := router.Group("/api")
	api.Use(AuthJwt.Auth)
	{		
		//sistem
		api.GET("/checkauth", Sistem.CheckAuth)
		api.GET("/sistem/password", Sistem.GetPassword)
		api.PATCH("/sistem/piket/password", Sistem.PostPasswordPiket)		
		api.POST("/sistem/piket/auth", Sistem.AuthPiket)
		api.PATCH("/sistem/clock/password", Sistem.PostPasswordClock)	
		api.GET("/sistem/backup/export", Sistem.DoBackupDatabase)		
		// info rekap siswa	dashboard & profile guru
		api.GET("/info/siswa/harian", Info.RekapHarian)
		api.GET("/info/siswa/semester", Info.RekapSemester)		
		api.GET("/info/guru/insight", Info.InsightGuru) 		
		//realtime all siswa info izin,sakit,alpa,terlambat
		api.GET("/info/absen/harian", Info.AbsenHarian)
		// user route
		//api.GET("/admin-users",auth ,adminUser.Index)		
		api.GET("/users", Users.Index)
		api.GET("/arsip/users", Users.Arsip)
		api.GET("/alluser", Users.All)
		api.POST("/users", Users.Store)
		api.POST("/usershapus", Users.MultiDelete)
		api.PATCH("/users/:id", Users.Update)
		api.PATCH("/users/:id/password", Users.Password)
		api.DELETE("/users/:id", Users.Destroy)
		api.GET("/users/:id", Users.Show)
		//upload router	
		api.POST("/upload/guru", Upload.StorageGuru)
		api.POST("/upload/siswa", Upload.StorageSiswa)
		//kelas
		api.GET("/allkelas", Kelas.All)
		api.GET("/kelas", Kelas.Index)
		api.POST("/kelas", Kelas.Store)
		api.PATCH("/kelas/:id", Kelas.Update)
		api.DELETE("/kelas/:id", Kelas.Destroy)
		//siswa
		api.GET("/allsiswa", Siswa.All)
		api.GET("/siswa", Siswa.Index)
		api.GET("/arsip/siswa", Siswa.Arsip)
		api.GET("/siswajumlah", Siswa.Jumlah)
		api.POST("/siswa", Siswa.Store)
		api.POST("/siswahapus", Siswa.MultiDelete)
		api.PATCH("/siswa/:id", Siswa.Update)
		api.PATCH("/multisiswa", Siswa.Multi)
		api.DELETE("/siswa/:id", Siswa.Destroy)	
		//semester
		api.GET("/semester/lists", Semester.ListAll)
		api.GET("/semester", Semester.Index)
		api.POST("/semester", Semester.Store)
		api.PATCH("/semester/:id", Semester.Update)
		api.DELETE("/semester/:id", Semester.Destroy)
		//jadwal
		api.GET("/jadwal", Jadwal.Index)
		api.GET("/jadwalabsen", Jadwal.JadwalAbsen)
		api.GET("/cekmapel", Jadwal.CekMapel)
		api.POST("/jadwal", Jadwal.Store)
		api.PATCH("/jadwal/:id", Jadwal.Update)
		api.DELETE("/jadwal/:id", Jadwal.Destroy)
		//absensi
		api.GET("/absensi", Absen.Index)
		api.GET("/absensi/harian", Absen.Harian)
		api.GET("/absensi/cekmapel", Absen.CekMapel)
		api.GET("/absensi/mobile", Absen.IndexMobile)
		api.GET("/absensi/listtanggal", Absen.ListTanggal)
		api.GET("/absensi/getsiswa", Absen.Siswa)
		api.GET("/absensi/getabsen", Absen.GetAbsensi)
		api.POST("/absensi/getabsen/store", Absen.PostAbsensi)
		api.POST("/absensi/mobile/store", Absen.StoreAbsen)
		api.POST("/absensi/mobile/updatejamkeluar", Absen.UpdateJamKeluar)
		api.POST("/absensi/store", Absen.Store)		
		api.PATCH("/absensi/:id", Absen.Update)
		api.DELETE("/absensi/:id", Absen.Destroy)
		api.POST("/absensihapus", Absen.MultiDelete)
		//rekap
		api.GET("/rekap/siswa", Rekap.Siswa)
		api.GET("/rekap/kelas", Rekap.Kelas)
		api.GET("/rekap/guru", Rekap.Guru)
		//laporan
		api.GET("/laporan/kelas/absensi", Laporan.AbsensiKelas)
		api.POST("/laporan/kelas/absensi", Laporan.AbsensiKelasExcel)
		api.GET("/laporan/clock", Laporan.AbsensiClock)
		api.GET("/laporan/clock/piket", Laporan.AbsensiClockPiket)
		api.POST("/laporan/clock", Laporan.AbsensiClockExcel)
	}	
	//foto storage
	//router.Static("/storage", "./storage")
	router.Use(static.Serve("/storage", static.LocalFile("./storage", true))) 
	//login	
	router.POST("/login", AuthJwt.LoginHandler)	
	//websocket
	router.GET("/websocket/:name/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})
	// end websocket
	//-----
	return router
}

// ---- asset to bytes

// type binaryFileSystem struct {
//     fs http.FileSystem
// }

// func (b *binaryFileSystem) Open(name string) (http.File, error) {
//     return b.fs.Open(name)
// }

// func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

//     if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
//         if _, err := b.fs.Open(p); err != nil {
//             return false
//         }
//         return true
//     }
//     return false
// }

// func BinaryFileSystem(root string) *binaryFileSystem {
//     fs := &assetfs.AssetFS{Asset, AssetDir, AssetInfo, root}
//     return &binaryFileSystem{
//         fs,
//     }
// }