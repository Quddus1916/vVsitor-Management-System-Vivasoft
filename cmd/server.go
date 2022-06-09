package cmd

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/config"
	"visitor-management-system/database"
	"visitor-management-system/routes"
)

func Execute() {
	e := echo.New()
	config := config.GetConfig()
	database.GetDB("mytestdb")
	routes.Subcription(e)
	e.Start(":" + config.Port)
}
