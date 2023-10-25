package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba.com/go_realworld/params/request"
	"github.com/vinoMamba.com/go_realworld/utils"
)

func UserHandler(r *gin.Engine) {
	ug := r.Group("/api/users")
	ug.POST("/login", userLogin)
	ug.POST("", userRegistration)
}

func userLogin(c *gin.Context) {
	body := &request.UserAuthenticationRequest{}

	if err := c.ShouldBindJSON(body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(utils.JsonMarshal(body))
}

func userRegistration(c *gin.Context) {
	body := &request.UserRegistrationRequest{}

	if err := c.ShouldBindJSON(body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(utils.JsonMarshal(body))
}
