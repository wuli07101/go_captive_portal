package webserver

import (
	"github.com/Nrehearsal/go_captive_portal/config"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func Run(gwHttpConf config.GatewayHttp) {
	router := gin.Default()

	router.Use(GatewaySSLOn())

	router.GET("/auth", Auth)
	router.POST("/adduser", AddUser)
	router.GET("/onlinelist", OnlineList)
	router.GET("/kickoutuser", KickOutUser)

	/*
	router.GET("/onlinelist", OnlineListProxy)
	router.POST("/kickouter", KickOutUserProxy)
	router.POST("/adduser", AddUserProxy())
	*/

	router.NoRoute(NotFound404)

	if gwHttpConf.SSLOn {
		go func() {
			log.Fatal(router.RunTLS(":"+gwHttpConf.SSLPort, gwHttpConf.SSLCrt, gwHttpConf.SSLKey))
		}()
	}

	log.Fatal(router.Run(":" + gwHttpConf.Port))
}
