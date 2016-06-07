package cli

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run(config *Configuration) {
	server := gin.Default()
	router := NewRouterFactory(server)

	for uri, settings := range config.Services {
		api := &EndpointFactory{
			uri:         uri,
			service:     &settings,
			latency:     config.Latency,
			contentType: config.ContentType,
			mockDir:     config.MockDir,
		}

		api.register(router)
	}

	server.Run(fmt.Sprintf(":%s", config.Port))
}
