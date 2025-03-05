package main

import (
	"github.com/alextrufmanov/asyncCalculator/pkg/agent"
	"github.com/alextrufmanov/asyncCalculator/pkg/config"
)

func main() {
	config := config.NewCfg()
	agent.StartAgent(*config)
}
