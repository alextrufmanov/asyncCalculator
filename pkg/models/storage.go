package models

// Константы возможных состояний арифметических выражений
const (
	ExpressionStatusReady     = "ready"
	ExpressionStatusCalculate = "calculate"
	ExpressionStatusFailed    = "failed"
	ExpressionStatusSuccess   = "success"
)

// Константы возможных состояний задач вычисления арифметических выражений
const (
	TaskStatusWait      = "wait"
	TaskStatusReady     = "ready"
	TaskStatusCalculate = "calculate"
	TaskStatusFailed    = "failed"
	TaskStatusSuccess   = "success"
)

// Структура арифметического выражения
type Expression struct {
	Expression string
	Id         int     `json:"id"`
	Status     string  `json:"status"`
	Result     float64 `json:"result"`
	Tasks      []*Task `json:"-"`
}

// Структура задач вычисления арифметических выражений
type Task struct {
	Owner         *Expression  `json:"-"`
	Id            int          `json:"id"`
	Arg1          float64      `json:"arg1"`
	Arg2          float64      `json:"arg2"`
	Operation     string       `json:"Operation"`
	OperationTime int          `json:"operation_time"`
	Status        string       `json:"-"`
	Result        float64      `json:"-"`
	ResultChan    chan float64 `json:"-"`
}
