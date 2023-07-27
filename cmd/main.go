package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("等候 2 秒鐘")
	}
	log.Println("Server Exit...")
}
