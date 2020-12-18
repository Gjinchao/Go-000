package api

import (
	"github.com/gin-gonic/gin"
)

func NewAccountInterface(accountService AccountService) {
	r := gin.Default()
	account := r.Group("/accounts")
	{
		account.POST("", accountService.Register())
		account.DELETE("", accountService.Unregister())
		account.PUT("/:id", accountService.UpdateProfile())
		account.Group("/:id", accountService.GetProfile())
	}
	session := r.Group("/sessions")
	{
		session.POST("", accountService.Login())
		session.DELETE("", accountService.Logout())
	}
}

type AccountService interface {
	Register() gin.HandlerFunc
	Unregister() gin.HandlerFunc
	UpdateProfile() gin.HandlerFunc
	GetProfile() gin.HandlerFunc
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
}
