package cli

import (
	"fmt"
	"net/http"

	"github.com/helderfarias/dynamock/cli/middleware"
	"github.com/labstack/echo"
)

// Run statup
func Run(config *Configuration) {
	server := echo.New()

	router := NewRouterFactory(server)
	for uri, settings := range config.Services {
		api := &EndpointFactory{
			uri:         uri,
			service:     &settings,
			latency:     config.Latency,
			contentType: config.ContentType,
			mockDir:     config.MockDir,
			cors:        config.Cors,
		}

		api.register(router)
	}

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.Cors.AllowOrigin},
		AllowHeaders:     []string{config.Cors.AllowHeaders},
		AllowMethods:     []string{config.Cors.AllowMethods},
		AllowCredentials: config.Cors.AllowCredentials,
		ExposeHeaders:    []string{config.Cors.ExposeHeaders},
	}))
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	server.Use(middleware.Recover())

	server.GET("/", welcome)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", config.Port)))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}
