package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler обрабатывает HTTP-запрос и возвращает содержимое HTML-страницы.
//
// Принимает:
//   - http.ResponseWriter — используется для формирования и отправки HTTP-ответа клиенту;
//   - *http.Request — входящий HTTP-запрос.
//
// Возвращает:
//   - HTML-страницу (index.html) с кодом 200 при успешном выполнении;
//   - HTTP 500 и текст ошибки, если не удалось прочитать файл.
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Ошибка при чтении файла: "+err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

// UploadHandler обрабатывает загрузку файлов, конвертирует содержимое.
//
// Принимает:
//   - http.ResponseWriter — для записи ответа;
//   - *http.Request — HTTP-запрос с формой.
//
// Возвращает:
//   - HTTP 500 и текст ошибки, если не удалось прочитать или конвертировать файл.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Ошибка формы: "+err.Error(), 500)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка получения файла: "+err.Error(), 500)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла: "+err.Error(), 500)
		return
	}

	res, err := service.ConvertAuto(string(data))
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), 500)
		return
	}

	filename := time.Now().Format("20060102_150405") + filepath.Ext(header.Filename)
	out, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Ошибка создания файла: "+err.Error(), 500)
		return
	}
	defer out.Close()

	_, err = out.WriteString(res)
	if err != nil {
		http.Error(w, "Ошибка записи файла: "+err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(res))
}
