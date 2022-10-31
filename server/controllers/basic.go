package controllers

import (		
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"encoding/json"
	"../config"
	"log"
	"os"	
)

type Basic struct {

}

func (basic *Basic) JsonSuccess (c *gin.Context, status int, h gin.H) {
	h["status"] = "success"
	(*c).JSON(status , h)
	return
}

func (basic *Basic) JsonFail (c *gin.Context, status int, message string) {
	(*c).JSON(status , gin.H{
		"status" 	: "fail",
		"message"	: message,
	})
}

func (basic *Basic) Role (c *gin.Context, rules string) bool {
	role := (*c).MustGet("rules").(string)
	if role == rules || role == "superuser" {
		return true
	}else{
		return false
	}
}

func (basic *Basic) fileRem (filename string) {
	
	err := os.Remove(filename);if err !=nil {
    	log.Println(err)
    }
}

func (basic *Basic) SocketSend (channel string, arr CommandWS) {
	u := "ws://"+config.Get().Addr+"/websocket/"+channel+"/ws"
	c, _, err := websocket.DefaultDialer.Dial(u, nil)	
	if err != nil {
		log.Println("dial:", err)			
	}
	defer c.Close()	
	emp, _ := json.Marshal(arr)
	c.WriteMessage(websocket.TextMessage, []byte(emp))
	if err != nil {
		log.Println("write:", err)
		return
	}
}

type CommandWS struct {
    Name string    
    Data string
}