package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/di"
	"noneland/backend/interview/internal/pkg"
)

func main() {
	env := os.Getenv("ENV")
	env = "template-dev" // 範例用途, 理想上應該是從環境變數讀到 local, dev, test, stage, prod
	config := configs.NewConfigFromFilename(env)
	engine := di.NewGin(config)

	var mux http.Handler
	switch env {
	case "stage":
		mux = engine
	case "prod":
		mux = pkg.SetupHttp2(engine)
	}

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Port),
		Handler:        mux,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("開始監聽 %v\n", config.Port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
