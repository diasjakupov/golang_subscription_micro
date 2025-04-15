package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPServer HTTPServerConfig
	DB         DBConfig
}

type HTTPServerConfig struct {
	Host string
	Port int
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: не удалось загрузить .env файл(ы):", err)
	}

	conf := &Config{}

	// HTTP
	conf.HTTPServer.Host = os.Getenv("HTTP_SERVER_HOST")
	portStr := os.Getenv("HTTP_SERVER_PORT")
	if portStr == "" {
		return nil, fmt.Errorf("переменная окружения HTTP_SERVER_PORT не установлена")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("некорректное значение HTTP_SERVER_PORT: %w", err)
	}
	conf.HTTPServer.Port = port

	//DB
	conf.DB.Driver = os.Getenv("DB_DRIVER")
	if conf.DB.Driver == "" {
		return nil, fmt.Errorf("переменная окружения DB_DRIVER не установлена")
	}

	conf.DB.Host = os.Getenv("DB_HOST")
	if conf.DB.Host == "" {
		return nil, fmt.Errorf("переменная окружения DB_HOST не установлена")
	}

	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		return nil, fmt.Errorf("переменная окружения DB_PORT не установлена")
	}
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("некорректное значение DB_PORT: %w", err)
	}
	conf.DB.Port = dbPort

	conf.DB.User = os.Getenv("DB_USER")
	if conf.DB.User == "" {
		return nil, fmt.Errorf("переменная окружения DB_USER не установлена")
	}

	conf.DB.Password = os.Getenv("DB_PASSWORD")
	if conf.DB.Password == "" {
		return nil, fmt.Errorf("переменная окружения DB_PASSWORD не установлена")
	}

	conf.DB.DBName = os.Getenv("DB_NAME")
	if conf.DB.DBName == "" {
		return nil, fmt.Errorf("переменная окружения DB_NAME не установлена")
	}

	conf.DB.SSLMode = os.Getenv("DB_SSLMODE")
	if conf.DB.SSLMode == "" {
		conf.DB.SSLMode = "disable"
	}

	return conf, nil
}
