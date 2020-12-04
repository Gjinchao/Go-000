package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	Username string
	Address  string
}

func init() {
	db, _ = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4", "root", "root", "localhost", "3306", "test"))
}

func SelectByUsername(username string) (user *User, err error) {
	user = &User{}
	err = db.QueryRow("select username, address from test.users where username = ?", username).Scan(&user.Username,
		&user.Address)
	if err != nil {
		return nil, fmt.Errorf("can not query user with username \"%s\" ----> %w", username, err)
	}
	return user, nil
}
