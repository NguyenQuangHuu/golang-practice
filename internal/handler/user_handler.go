package handler

import (
	"awesomeProject/internal/helpers"
	"awesomeProject/internal/model"
	"awesomeProject/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type IUserHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type UserHandler struct {
	service service.IUserService
}

func NewUserHandle(userService service.IUserService) *UserHandler {
	return &UserHandler{service: userService}
}

func (u *UserHandler) Login(c *gin.Context) {
	var user model.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := u.service.GetUserByUsername(user.Username)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "username not exist"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Wrong username or password"})
		return
	}

	tokenResponse, err := helpers.GenerateToken(result, 5*60*1000*1000*1000)
	//c.SetCookie("token-jwt", tokenResponse, 60*60*30, "/", "localhost", false, true)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token-jwt",
		Value:    tokenResponse,
		MaxAge:   60 * 60 * 60 * 30,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
		Domain:   "localhost",
	})

	c.IndentedJSON(http.StatusOK, gin.H{"data": tokenResponse})
}

func (u *UserHandler) Register(c *gin.Context) {
	var register model.RegisterRequest
	if err := c.ShouldBindJSON(&register); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errs := helpers.ValidateStruct(register)
	//Get User By Username
	if len(errs) > 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}
	err := u.service.PreprocessBeforeSaveUser(&register)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "register successfully"})
}

func (u *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("token-jwt", "", -1, "/", "", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "logout successfully"})
}
