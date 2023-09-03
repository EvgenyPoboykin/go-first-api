package eqes

import (
	getRegister "evgenypoboykin/first-api/services/eqes/get-register"
	postRegister "evgenypoboykin/first-api/services/eqes/post-register"

	"github.com/gin-gonic/gin"
)

func New(c *gin.Engine) {
	c.GET("/@eqes-register/signatures", getRegister.New)
	c.POST("/@eqes-register/signatures", postRegister.New)
}
