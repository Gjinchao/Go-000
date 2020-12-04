package main

import (
	"github.com/Gjinchao/Go-000/Week02/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/address", controller.GetAddress)
	err := http.ListenAndServe("localhost:8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}
