package main

import (
	"ginvue/pkg/database"
	"ginvue/pkg/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := database.GetDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	r = router.CollectRouter(r)
	panic(r.Run())
}
