package biz

import (
	"github.com/Gjinchao/Go-000/Week04/internal/data"
	"time"
)

type AccountDo struct {
	Id       string
	Username string
	Phone    string
	EncryPwd string
}

type accountBiz struct {
	accountRepo AccountRepo
}

func NewAccountBiz(accountRepo AccountRepo) accountBiz {
	return accountBiz{accountRepo: accountRepo}
}

func (accountBiz accountBiz) Find(do *AccountDo) bool {
	repo := accountBiz.accountRepo
	_, err := repo.QueryByPhone(do.Phone)
	if err == nil {
		return true
	}
	_, err = repo.QueryByUsername(do.Username)
	if err == nil {
		return true
	}
	return false
}

func (accountBiz accountBiz) Insert(do *AccountDo) error {
	now := time.Now()
	po := data.AccountPo{
		Id:        do.Id,
		Username:  do.Username,
		Phone:     do.Phone,
		EncryPwd:  do.EncryPwd,
		CreatedAt: &now,
		DeletedAt: nil,
		UpdatedAt: &now,
	}
	return accountBiz.accountRepo.Insert(&po)
}

type AccountRepo interface {
	QueryById(id string) (data.AccountPo, error)
	QueryByUsername(username string) (data.AccountPo, error)
	QueryByPhone(phone string) (data.AccountPo, error)
	Insert(*data.AccountPo) error
	//todo
}
