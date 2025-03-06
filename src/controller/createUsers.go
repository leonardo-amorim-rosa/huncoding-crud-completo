package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardo-amorim-rosa/huncoding-crud-completo/src/config/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Invalid JSON body")
	c.JSON(err.Code, err)
}
