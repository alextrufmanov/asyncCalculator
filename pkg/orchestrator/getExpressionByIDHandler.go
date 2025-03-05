package orchestrator

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/alextrufmanov/asyncCalculator/pkg/models"
	"github.com/gorilla/mux"
)

// Функция возвращает обработчик GET запроса эндпоинта /api/v1/expressions/{id},
// запрос у оркестратора информации о арифметическом выражении с указанным id
func GETExpressionByIDHandler(s *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err == nil {
			log.Printf("GET: /api/v1/expressions/%d", id)
			expression, r := s.GetExpressionByID(id)
			if r {
				json.NewEncoder(w).Encode(models.GETexpressionByIDAnswerBody{Expression: expression})
			} else {
				SendNotFoundError404(w)
			}
		} else {
			log.Printf("GET: /api/v1/expressions/{???}}")
			SendInternalError500(w)
		}
	}
}
