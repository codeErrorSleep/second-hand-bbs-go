package main

import (
	"fmt"
	"net/http"
	"second-hand-bbs-go/internal/routers"
	"second-hand-bbs-go/utils"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", utils.HTTPPort),
		Handler:        router,
		ReadTimeout:    utils.ReadTimeout,
		WriteTimeout:   utils.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
