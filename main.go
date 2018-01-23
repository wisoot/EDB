package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"edb/web/routes"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: ", err)
	}
}

func main() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))

	router.Use(sessions.Sessions("edb", store))

	routes.InitV1Routes(router)

	router.Run()
}
