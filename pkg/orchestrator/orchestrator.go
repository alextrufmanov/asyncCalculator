package orchestrator

import (
	"log"
	"net/http"

	"github.com/alextrufmanov/asyncCalculator/pkg/config"
	"github.com/gorilla/mux"
)

// Функция создания и запуска сервера оркестратора
func StartOrchestrator(cfg config.Cfg) {
	// создаем хранилище арифметических выражений и задач агента
	storage := NewStorage(cfg)
	// "настраиваем" мультиплексор
	router := mux.NewRouter()
	// router.HandleFunc("/", indexHandler)
	router.HandleFunc("/api/v1/calculate", POSTCalculateHandler(storage)).Methods("POST")
	router.HandleFunc("/api/v1/expressions", GETExpressionsHandler(storage)).Methods("GET")
	router.HandleFunc("/api/v1/expressions/{id}", GETExpressionByIDHandler(storage)).Methods("GET")
	router.HandleFunc("/internal/task", GETTaskHandler(storage)).Methods("GET")
	router.HandleFunc("/internal/task", POSTTaskResultHandler(storage)).Methods("POST")
	// запускаем сервер оркестратора
	log.Printf("Orchestrator started on %s", cfg.Addr)
	err := http.ListenAndServe(cfg.Addr, router)
	if err != nil {
		log.Fatal("... with error:", err)
	}
}
