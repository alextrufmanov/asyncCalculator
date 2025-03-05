package orchestrator

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alextrufmanov/asyncCalculator/pkg/models"
)

// Функция возвращает обработчик GET запроса эндпоинта /api/v1/expressions,
// запрос у оркестратора информации обо всех арифметических выражениях
func GETExpressionsHandler(s *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET: /api/v1/expressions")
		json.NewEncoder(w).Encode(models.GETexpressionsAnswerBody{Expressions: s.GetAllExpressions()})
	}
}
