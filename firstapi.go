package main

import (
	"firstapi/services/eqes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	eqes.New(router)

	router.Run(":7878")
}
