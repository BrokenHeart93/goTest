package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)

	host := viper.GetString("web.host")
	port := viper.GetString("web.port")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "10000"
	}

	if err := http.ListenAndServe(host+":"+port, router); err != nil {
		fmt.Println("err:", err)
	}
}