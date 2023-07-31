package main

import (
	"ginvue/pkg/config"
	"ginvue/pkg/database"
	"ginvue/pkg/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitViper()
	database.InitDB()
	r := gin.Default()

	db := database.GetDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	r = router.CollectRouter(r)
	panic(r.Run())

}
