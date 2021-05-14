package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"room/server"
)

func main() {
	router := mux.NewRouter()
	go server.H.Run()
	router.HandleFunc("/ws", server.Myws)

	host := viper.GetString("web.host")
	port := viper.GetString("web.port")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "10000"
	}

	if err := http.ListenAndServe(host, router); err != nil {
		fmt.Println("err:", err)
	}
}