// Структура проекта для расширенного сервиса сегментации на Go
// 7 файлов: main.go, store.go, segment.go, user.go, membership.go, http_handlers.go, csv_utils.go

// --- main.go ---
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	store := NewStore()
	httpServer := NewHTTPServer(store)

	log.Printf("Starting server on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, httpServer))
}
