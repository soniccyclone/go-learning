package main

import "log"

func main() {
	// Load Configurations from config.json using Viper
	config, configErr := LoadAppConfig()
	if configErr != nil {
		log.Fatal(configErr)
		return
	}
	app, appErr := NewApp(config.ConnectionString)
	if appErr != nil {
		log.Fatal(appErr)
		return
	}
	app.Start(config.Port)
}
