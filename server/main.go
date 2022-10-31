package main

import (
	"fmt"
	//"github.com/denisbrodbeck/machineid"
	"./config"
	"./database"
	"./routes"
)

func main() {
	if err := config.Load("config.yaml"); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

	// id, _ := machineid.ProtectedID("myAppName")

	// if (id != config.Get().License) {
	// 	fmt.Println("License Key Anda Tidak Terdaftar")
	// 	return
	// }

	db, err := database.InitDB()
	if err != nil {
		fmt.Println("err open databases")
		return
	}

	defer db.Close()

	router := routes.InitRouter()
	router.Run(config.Get().Addr)
}
