package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

// main запускает сервер и логирует ошибки.
// Если сервер не запускается, ошибка выводится в лог и выполнение завершается.
func main() {
	logger := log.New(os.Stdout, "morse: ", log.LstdFlags)
	err := server.RunServer(logger)
	if err != nil {
		logger.Fatalf("Ошибка сервера: %v", err)
	}
}
