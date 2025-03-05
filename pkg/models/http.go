package models

// Структура тела POST запроса на вычисление арифметического выражения
type POSTCalculateRequestBody struct {
	Expression string `json:"expression"`
}

// Структура тела ответа на POST запрос на вычисление арифметического выражения
type POSTCalculateAnswerBody struct {
	Id int `json:"id"`
}

// Структура тела ответа на GET запрос списка арифметических выражений (обработанных и находящихся в обработке)
type GETexpressionsAnswerBody struct {
	Expressions []Expression `json:"expressions"`
}

// Структура тела ответа на GET запрос арифметического выражения по его Id
type GETexpressionByIDAnswerBody struct {
	Expression Expression `json:"expression"`
}

// Структура тела ответа на внутренний GET запрос задачи агентом
type GETTaskAnswerBody struct {
	Task Task `json:"task"`
}

// Структура тела внутреннего POST запроса передачи результатов вычислений агентом
type POSTTaskResultRequestBody struct {
	Id      int     `json:"id"`
	Result  float64 `json:"result"`
	Success bool    `json:"success"`
}
