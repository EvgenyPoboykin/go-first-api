package getRegister

import (
	registerTypes "firstapi/services/eqes/register-types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var itemGet = registerTypes.ItemSignature{
	SerialNumber: "SerialNumber-TEST-5590D262EE31BA07BCEB7AEF59EF9AF03A1ADA25-GET",
	Thumbprint:   "Thumbprint-TEST-5590D262EE31BA07BCEB7AEF59EF9AF03A1ADA25-GET",
}

func New(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, itemGet)
}
