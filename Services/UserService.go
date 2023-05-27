package Services

import (
	"cycling-tracker-server/Models"
	"cycling-tracker-server/Repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserService struct {
	userRepository Repositories.UserRepository
}

func (u *UserService) Init() {
	u.userRepository = Repositories.UserRepository{}
}

func (u *UserService) LoginUser(c *gin.Context) {
	var login Models.Login

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
