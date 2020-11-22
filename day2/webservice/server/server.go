package server

import "../config"

func Init() {
	conf := config.GetConfig()
	router := NewRouter()
	router.Run(":" + conf.GetString("server.port"))
}
