package service

import (
	"crypto/md5"
	"fmt"
	"github.com/Gjinchao/Go-000/Week04/internal/biz"
	"github.com/gin-gonic/gin"
)

type accountDto struct {
	id       string `form:"id"`
	username string `form:"username"`
	phone    string `form:"phone"`
	password string `form:"password"`
}

type accountService struct {
	accountBiz AccountBiz
}

func NewAccountService(accountBiz AccountBiz) accountService {
	return accountService{accountBiz: accountBiz}
}

func (accountService accountService) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto accountDto
		err := ctx.ShouldBind(&dto)
		if err != nil {
			ctx.JSON(400, "bad request")
			return
		}
		do := biz.AccountDo{
			Id:       dto.id,
			Username: dto.username,
			Phone:    dto.phone,
			EncryPwd: fmt.Sprintf("%x", md5.Sum([]byte(dto.password))),
		}
		accountBiz := accountService.accountBiz
		if accountBiz.Find(&do) {
			ctx.JSON(403, "account exists")
			return
		}
		err = accountBiz.Insert(&do)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, "succeed")
	}
}

func (accountService accountService) Unregister() gin.HandlerFunc {
	//todo
	return nil
}

func (accountService accountService) UpdateProfile() gin.HandlerFunc {
	//todo
	return nil
}

func (accountService accountService) GetProfile() gin.HandlerFunc {
	//todo
	return nil
}

func (accountService accountService) Login() gin.HandlerFunc {
	//todo
	return nil
}

func (accountService accountService) Logout() gin.HandlerFunc {
	//todo
	return nil
}

type AccountBiz interface {
	Find(do *biz.AccountDo) bool
	Insert(do *biz.AccountDo) error
	//todo : other business
}
