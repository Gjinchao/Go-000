package controller

import (
	"fmt"
	"github.com/Gjinchao/Go-000/Week02/service"
	"net/http"
)

func GetAddress(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	address, err := service.GetAddress(query.Get("username"))
	if err != nil {
		w.WriteHeader(404)
		_, err = fmt.Fprintln(w, err)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	w.WriteHeader(200)
	_, err = fmt.Fprintln(w, address)
	if err != nil {
		fmt.Println(err)
	}
}
