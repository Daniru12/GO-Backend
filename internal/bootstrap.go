package internal

import (
	"project1/config"

	"project1/database"
	http "project1/transport"
)

func Init() {

	config.ParseAppConfig()
	//MySQL DB initialization
	database.Init()
	defer database.Close(database.Connections.Read)

	//HTTP transport initialization
	http.Listen()
}
