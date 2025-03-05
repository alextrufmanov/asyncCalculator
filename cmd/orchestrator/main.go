package main

import (
	"github.com/alextrufmanov/asyncCalculator/pkg/config"
	"github.com/alextrufmanov/asyncCalculator/pkg/orchestrator"
)

func main() {
	config := config.NewCfg()
	orchestrator.StartOrchestrator(*config)
}
