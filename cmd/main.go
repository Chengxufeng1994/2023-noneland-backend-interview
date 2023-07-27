package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"noneland/backend/interview/internal/di"
	"noneland/backend/interview/internal/pkg"
)

func main() {
	config := di.NewConfig()
	h2 := pkg.InitHttpHandler()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Port),
		Handler:        h2,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("開始監聽 %v\n", config.Port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
