package controller

import (
	"horizon/server/model"

	"github.com/gin-gonic/gin"
)

func Adduser(c *gin.Context) {
	res := model.AExample{
		Message: "Add user test",
	}

	c.JSON(200, res)
}