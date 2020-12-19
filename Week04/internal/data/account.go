package data

import (
	"fmt"
	"time"
)

type AccountPo struct {
	Id        string
	Username  string
	Phone     string
	EncryPwd  string
	CreatedAt *time.Time
	DeletedAt *time.Time
	UpdatedAt *time.Time
}

type accountRepo struct{}

func NewAccountRepo() accountRepo {
	return accountRepo{}
}

func (accountRepo) QueryById(id string) (AccountPo, error) {
	//todo
	return AccountPo{}, nil
}

func (accountRepo) QueryByUsername(username string) (AccountPo, error) {
	//todo
	return AccountPo{}, nil
}

func (accountRepo) QueryByPhone(phone string) (AccountPo, error) {
	//todo
	return AccountPo{}, nil
}

func (accountRepo) Insert(po *AccountPo) error {
	//todo
	fmt.Println(po.Username + " / " + po.Phone + " 持久化。")
	return nil
}
