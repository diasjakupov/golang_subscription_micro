package app

import (
	"fmt"
	"subscriptions/internal/app/config"
	"subscriptions/internal/app/start"
)

func Run(configFiles ...string) {
	cfg, err := config.NewConfig(configFiles...)
	if err != nil {
		fmt.Printf("error during config setup")
	}
	fmt.Printf("HTTP server config: %v\n", cfg.HTTPServer)
	fmt.Printf("DB server config: %v\n", cfg.DB)

	start.HTTP(cfg.HTTPServer)
}
