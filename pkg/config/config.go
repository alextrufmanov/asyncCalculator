package config

import (
	"os"
	"strconv"
)

type Cfg struct {
	Host           string
	Port           int
	Addr           string
	AddTimeout     int
	SubTimeout     int
	MltTimeout     int
	DivTimeout     int
	ComputingPower int
}

func getIntEnv(key string, defValue int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err == nil {
		return value
	}
	return defValue
}

func getHostEnv(key string, defHost string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defHost
}

func NewCfg() *Cfg {
	return &Cfg{
		Host:           getHostEnv("ASYNC_CALCULATOR_HOST", "localhost"),
		Port:           getIntEnv("ASYNC_CALCULATOR_PORT", 8080),
		Addr:           getHostEnv("HOST", "localhost") + ":" + strconv.Itoa(getIntEnv("PORT", 8080)),
		AddTimeout:     getIntEnv("TIME_ADDITION_MS", 5000),
		SubTimeout:     getIntEnv("TIME_SUBTRACTION_MS", 5000),
		MltTimeout:     getIntEnv("TIME_MULTIPLICATIONS_MS", 5000),
		DivTimeout:     getIntEnv("TIME_DIVISIONS_MS", 5000),
		ComputingPower: getIntEnv("COMPUTING_POWER", 10),
	}
}
