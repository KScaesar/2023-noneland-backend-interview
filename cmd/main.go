package main

import (
	"fmt"
	"log"
	"os"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/di"
)

func main() {
	env := os.Getenv("ENV")
	env = "template-dev" // 理想上應該是從環境變數讀到 local, dev, test, stage, prod
	cfg := configs.NewConfig(env)
	// fmt.Println(cfg)
	server := di.NewServerV1(cfg)

	fmt.Printf("開始監聽 %v\n", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
