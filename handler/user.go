package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba.com/go_realworld/logger"
	"github.com/vinoMamba.com/go_realworld/params/request"
	"github.com/vinoMamba.com/go_realworld/params/response"
	"github.com/vinoMamba.com/go_realworld/utils"
)

func UserHandler(r *gin.Engine) {
	ug := r.Group("/api/users")
	ug.POST("/login", userLogin)
	ug.POST("", userRegistration)
}

func userLogin(c *gin.Context) {
	log := logger.New(c)
	body := &request.UserAuthenticationRequest{}

	if err := c.ShouldBindJSON(body); err != nil {
		log.WithError(err).Errorln("Bind Json Error")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateJWT(body.User.Email, "vino")
	if err != nil {
		log.WithError(err).Errorln("Generate JWT Error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    body.User.Email,
			Username: "vino",
			Bio:      "",
			Image:    "",
			Token:    token,
		},
	})
}

func userRegistration(c *gin.Context) {
	log := logger.New(c)
	body := &request.UserRegistrationRequest{}

	if err := c.ShouldBindJSON(body); err != nil {
		log.WithError(err).Errorln("Bind Json Error")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateJWT(body.User.Email, body.User.Username)
	if err != nil {
		log.WithError(err).Errorln("Generate JWT Error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    body.User.Email,
			Username: body.User.Username,
			Bio:      "",
			Image:    "",
			Token:    token,
		},
	})

}
