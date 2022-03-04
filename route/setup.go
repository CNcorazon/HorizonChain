package route

import (
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {

	r := gin.Default()
	userRoute(r)
	witnessRoute(r)
	return r
}
