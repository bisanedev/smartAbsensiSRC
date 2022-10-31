package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"../config"
	"../database"
	"../models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var Key = "Rahasia Coy"

type AuthJwt struct {
	Basic
}

func (AuthJwt) Auth(c *gin.Context) {
	var users models.Users
	tokenString := (*c).Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(Key), nil
	})

	//cek
	if token != nil && err == nil {
		//cek token timeunix
		claims := token.Claims.(jwt.MapClaims)
		if errs := database.DB.Where("token = ?", claims["token"]).Where("status = ?", "a").Where("id = ?", claims["id"]).First(&users).Error; errs != nil {
			result := gin.H{
				"message": "not authorized",
				"error":   "Users Token Not Found",
			}
			(*c).JSON(http.StatusUnauthorized, result)
			(*c).Abort()
		} else {
			(*c).Set("rules", claims["role"])
			(*c).Set("id", claims["id"])
			fmt.Println("Token User:", users.Username, "Telah Mengakses")
		}

	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		(*c).JSON(http.StatusUnauthorized, result)
		(*c).Abort()
	}

}

func (a *AuthJwt) LoginHandler(c *gin.Context) {
	var users models.Users
	var request Credential
	var now = time.Now()

	if err := (*c).ShouldBind(&request); err != nil {
		(*a).JsonFail(c, http.StatusUnauthorized, "Silahkan Periksa Kembali Input Data !")
		return
	}

	password := []byte(request.Password)
	md5Ctx := md5.New()
	md5Ctx.Write(password)
	cipherStr := md5Ctx.Sum(nil)

	if database.DB.Where("username = ?", request.Username).Where("status = ?", "a").Where("password = ?", hex.EncodeToString(cipherStr)).First(&users).Error != nil {
		(*a).JsonFail(c, http.StatusUnauthorized, "Maaf Salah Username Atau Password")
		return
	} else {
		users.Token = now.Unix()

		if err := database.DB.Save(&users).Error; err != nil {
			(*a).JsonFail(c, http.StatusUnauthorized, err.Error())
			return
		}

		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    users.ID,
			"nama":  users.Nama,
			"jenis": users.Jenis,
			"foto":  users.Foto,
			"color": users.Color,
			"mapel": users.Mapel,
			"role":  users.Role,
			"token": now.Unix(),
		})

		token, err := (*sign).SignedString([]byte(Key))
		if err != nil {
			(*a).JsonFail(c, http.StatusUnauthorized, err.Error())
		}

		//get local ip address
		ip, err := externalIP()
		if err != nil {
			fmt.Println(err)
		}

		(*a).JsonSuccess(c, http.StatusOK, gin.H{"ipaddress": ip + config.Get().Addr, "token": token, "sekolah": config.Get().Sekolah, "warna": config.Get().Warna})
	}

}

type Credential struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
