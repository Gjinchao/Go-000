package main

import "log"

func main() {
	err := InitializeAllInstance().Run(":80")
	if err != nil {
		log.Fatal(err.Error())
	}
}
