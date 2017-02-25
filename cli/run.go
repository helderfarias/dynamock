package cli

import (
	"fmt"

	"github.com/labstack/echo"
)

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

	// if config.Cors != nil {
	// 	server.Use(CorsMiddleware(config.Cors))
	// }

	server.Start(fmt.Sprintf(":%s", config.Port))
}
