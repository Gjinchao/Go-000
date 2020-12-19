package api

import (
	"github.com/gin-gonic/gin"
)

type AccountInterface struct {
	*gin.Engine
}

func NewAccountInterface(accountService AccountService) AccountInterface {
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

	return AccountInterface{r}
}

type AccountService interface {
	Register() gin.HandlerFunc
	Unregister() gin.HandlerFunc
	UpdateProfile() gin.HandlerFunc
	GetProfile() gin.HandlerFunc
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
}
