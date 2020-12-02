package service

import "github.com/Gjinchao/Go-000/Week02/dao"

func GetAddress(username string) (string, error) {
	user, err := dao.SelectByUsername(username)
	if err != nil {
		return "", err
	}
	return user.Address, nil
}
