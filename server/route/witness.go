package route

import (
	"horizon/server/controller"

	"github.com/gin-gonic/gin"
)

func witnessRoute(r *gin.Engine) {
	user := r.Group("/witness")
	{
		user.POST("/requestTransaction", controller.WitnessGetTransaction)
	}
}
