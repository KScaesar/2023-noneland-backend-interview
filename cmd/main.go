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
	env = "template-dev" // 範例用途, 理想上應該是從環境變數讀到 local, dev, test, stage, prod
	cfg := configs.NewConfigFromFilename(env)
	server := di.NewServer(cfg)
	// fmt.Println(cfg)

	fmt.Printf("開始監聽 %v\n", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
