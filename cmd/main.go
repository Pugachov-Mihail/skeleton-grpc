package main

import "sceleton/internal/config"

func main() {
	cfg := config.MustLoad()
	log := config.SetupLoger(cfg.Env)
	log.Debug("Init app")

}
