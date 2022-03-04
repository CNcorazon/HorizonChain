package route

import (
	"horizon/controller"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/add", controller.Adduser)
	}
}
