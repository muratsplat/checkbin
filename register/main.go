package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/muratsplat/checkbin/register/service"
)

var (
	addr = flag.String("addr", ":8080", "Service registery addess")
	max  = flag.Int("maxService", 10, "maximum number of service")
)

func main() {
	flag.Parse()
	color.Blue("Register service is running..")
	serviceList := service.NewServiceRepository(*max)

	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		var json service.Register

		if err := c.ShouldBindJSON(&json); err == nil {
			var service service.Service
			service.Register = json
			service.LastUpdate = time.Now()
			serviceList.Append(service)
			c.JSON(200, gin.H{
				"ok": true,
			})

			color.Green("A new service is registered. the service name is %s", service.Register.Name)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

	})

	r.GET("/list", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"services": serviceList.List,
		})
	})

	color.Green("the web server is up running for new service. Listing to %s", *addr)
	r.Run(*addr) // listen and serve on 0.0.0.0:8080
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
