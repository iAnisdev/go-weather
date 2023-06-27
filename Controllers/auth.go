package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupRequest struct {
	Email    string `json:"email" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id        int    `json:id`
	Email     string `json:email`
	Fullname  string `json:fullname`
	Password  string `json:password`
	CreatedAt string `json:createdAt`
	UpdatedAt string `json:updatedAt`
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func VerifyPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(ctx *gin.Context) {
	body := LoginRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid inputs, Please check your request body.",
		})
		return
	}

}

func Signup(ctx *gin.Context) {
	body := SignupRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid inputs, Please check your request body.",
		})
		return
	}
	hash := HashPassword(body.Password)
	body.Password = hash

}

func Profile(ctx *gin.Context) {

}
