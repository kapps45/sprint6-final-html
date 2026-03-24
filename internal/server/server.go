package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// RunServer запускает HTTP-сервер на указанном логгере.
//
// Принимает:
//   - *log.Logger — логгер для ошибок и информационных сообщений.
//
// Возвращает:
//   - error — если сервер не удалось запустить.
func RunServer(logger *log.Logger) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	logger.Println("Сервер запущен на", srv.Addr)
	return srv.ListenAndServe()
}
