package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vlasdash/dating_bot/config"
	"github.com/vlasdash/dating_bot/pkg/handler"
	"log"
	"net/http"
	"os"
)

const ConfigPath = "./config/"
const EnvPath = ".env.example"

func StartPasswordStorageBot() error {
	_ = godotenv.Load(EnvPath)

	err := config.LoadConfig(ConfigPath)
	if err != nil {
		return fmt.Errorf("read config failed: %v\n", err)
	}

	botToken := os.Getenv("BOT_TOKEN")
	h := handler.NewHandler(botToken)
	err = h.SetWebhook()
	if err != nil {
		return fmt.Errorf("set webhook error: %v\n", err)
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", config.C.App.Port), h)
	if err != nil {
		return fmt.Errorf("start server error: %v\n", err)
	}

	return nil
}

func main() {
	err := StartPasswordStorageBot()
	if err != nil {
		log.Fatal(err)
	}
}
