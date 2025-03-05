package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/alextrufmanov/asyncCalculator/pkg/config"
	"github.com/alextrufmanov/asyncCalculator/pkg/models"
)

// Функция пытается получить от оркестратора очередную задачу
func getTask(addr string) (models.Task, bool) {
	var answerBody models.GETTaskAnswerBody
	response, err := http.Get(fmt.Sprintf("http://%s/internal/task", addr))
	if err == nil {
		if response.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(response.Body)
			if err == nil {
				if json.Unmarshal(bodyBytes, &answerBody) == nil {
					return answerBody.Task, true
				}
			}
		}
	}
	return models.Task{}, false
}

// Функция отправляет результат решения задачи оркестратору
func postTaskResul(addr string, id int, result float64, success bool) bool {
	requestBody, err := json.Marshal(models.POSTTaskResultRequestBody{Id: id, Result: result, Success: success})
	if err == nil {
		response, err := http.Post(fmt.Sprintf("http://%s/internal/task", addr), "application/json", bytes.NewBuffer(requestBody))
		if err == nil {
			return response.StatusCode == http.StatusOK
		}
	}
	return false
}

// Функция отправляет результат решения задачи оркестратору
func calculate(task *models.Task) bool {
	// имитируем "длительную" задачу
	time.Sleep(time.Duration(task.OperationTime) * time.Millisecond)
	// выполняем задачу
	switch task.Operation {
	case "+":
		task.Result = task.Arg1 + task.Arg2
	case "-":
		task.Result = task.Arg1 - task.Arg2
	case "*":
		task.Result = task.Arg1 * task.Arg2
	case "/":
		if task.Arg2 == 0 {
			log.Printf("Задача %d: (%f) %s (%f) => Деление на ноль", task.Id, task.Arg1, task.Operation, task.Arg2)
			return false
		}
		task.Result = task.Arg1 / task.Arg2
	default:
		log.Printf("Задача %d: (%f) %s (%f) => Неподдерживаемый оператор", task.Id, task.Arg1, task.Operation, task.Arg2)
		return false
	}
	log.Printf("Задача %d: (%f) %s (%f) => (%f)", task.Id, task.Arg1, task.Operation, task.Arg2, task.Result)
	return true
}

// Функция создания и запуска агента
func StartAgent(cfg config.Cfg) {
	log.Printf("Agent started (%s).", cfg.Addr)
	// запускаем указанное количество вычислитетлей в отдельных горутинах
	for range cfg.ComputingPower {
		go func() {
			for {
				// пытаемся получить от оркестратора очередную задачу
				task, r := getTask(cfg.Addr)
				if r {
					// если задача получена, то решаем ее
					success := calculate(&task)
					// отправляем результат решения задачи оркестратору c
					postTaskResul(cfg.Addr, task.Id, task.Result, success)
				}
				time.Sleep(time.Duration(500) * time.Millisecond)
			}
		}()
	}
	log.Printf("%d calculators was started", cfg.ComputingPower)
	select {}
}
