package main

import (
	"net/http"
	"github.com/heavenlwf/go-blog/routers"
	"fmt"
	"github.com/heavenlwf/go-blog/pkg/config"
	"log"
	"time"
	_ "github.com/heavenlwf/go-blog/models"
	)

func main() {
	router := routers.InitRouter()

	//router := setupRouter()
	endPoint := fmt.Sprintf(":%d", config.Conf.HttpPort)
	readTimeout := config.Conf.ReadTimeout * time.Second
	writeTimeout := config.Conf.WriteTimeout * time.Second
	s := &http.Server{
		Addr: 			endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	//s.ListenAndServe()
	s.ListenAndServe()
}
