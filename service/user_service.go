package service

import (
	"cycling-tracker-server/models"
	"cycling-tracker-server/repositorie"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserService struct {
	userRepository repositorie.UserRepository
}

func (u *UserService) Init() {
	u.userRepository = repositorie.UserRepository{}
}

func (u *UserService) LoginUser(c *gin.Context) {
	var login models.Login

	err := c.ShouldBindJSON(&login)

	if err != nil {
		log.Printf("There was an error with parsing recived json: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("There was an error with parsing recived json: %v", err),
		})

		return
	}

	user, err := u.userRepository.LoginUser(login)
	if err != nil {
		log.Printf("There was an error getting the user: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("There was an error getting the user: %v", err),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": user.Token,
	})
}
