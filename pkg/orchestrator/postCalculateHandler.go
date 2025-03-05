package orchestrator

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/alextrufmanov/asyncCalculator/pkg/models"
)

// Функция возвращает обработчик POST запроса эндпоинта /api/v1/calculate
// запрос на асинхронное вычисление нового арифметического выражения
func POSTCalculateHandler(s *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody models.POSTCalculateRequestBody
		log.Printf("POST: /api/v1/calculate")
		bodyBytes, err := io.ReadAll(r.Body)
		if err == nil {
			if json.Unmarshal(bodyBytes, &requestBody) == nil {
				// log.Printf("  Body = %v", requestBody)
				id, res := s.AppendExpression(requestBody.Expression)
				if res {
					json.NewEncoder(w).Encode(models.POSTCalculateAnswerBody{Id: id})
					return
				}
			}
		}
		SendInvalidDataError422(w)
	}
}
