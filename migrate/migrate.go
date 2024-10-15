package main

import (
	"edc-proxy/config"
	"edc-proxy/models"
)

func init() {
	config.LoadConfig()
	config.ConnectDB()

}

func main() {
	config.DB.AutoMigrate(&models.Connector{})
}
