package cli

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type RouterFactory interface {
	CreateGET(settings *RouterSettings)
	CreatePOST(settings *RouterSettings)
	CreatePUT(settings *RouterSettings)
	CreateDELETE(settings *RouterSettings)
}

type RouterSettings struct {
	Uri         string
	ContentType string
	Status      int
	Body        string
	BodyFile    string
	Latency     int
	Dynamic     map[string]interface{}
	Headers     map[string]string
}

type routerFactory struct {
	server *gin.Engine
}

func NewRouterFactory(s *gin.Engine) RouterFactory {
	return &routerFactory{server: s}
}

func (r *routerFactory) CreateGET(data *RouterSettings) {
	r.server.GET(data.Uri, func(c *gin.Context) {
		r.setLatency(c, data)
		r.setHeaders(c, data)
		r.execute(c, data)
	})
}

func (r *routerFactory) CreatePOST(data *RouterSettings) {
	r.server.POST(data.Uri, func(c *gin.Context) {
		r.setLatency(c, data)
		r.setHeaders(c, data)
		r.execute(c, data)
	})
}

func (r *routerFactory) CreatePUT(data *RouterSettings) {
	r.server.PUT(data.Uri, func(c *gin.Context) {
		r.setLatency(c, data)
		r.setHeaders(c, data)
		r.execute(c, data)
	})
}

func (r *routerFactory) CreateDELETE(data *RouterSettings) {
	r.server.DELETE(data.Uri, func(c *gin.Context) {
		r.setLatency(c, data)
		r.setHeaders(c, data)
		r.execute(c, data)
	})
}

func (r *routerFactory) execute(c *gin.Context, data *RouterSettings) {
	if data.ContentType == "application/json" {
		if len(data.Dynamic) > 0 {
			c.JSON(r.createDynamic(c, data))
			return
		}

		c.JSON(createSingleResult(data))
		return
	}

	c.String(data.Status, data.Body)
}

func (r *routerFactory) createDynamic(c *gin.Context, data *RouterSettings) (int, interface{}) {
	for key, input := range data.Dynamic {
		if key == "random" {
			plugin := &RandomPlugin{}
			mapstructure.Decode(input, plugin)
			return plugin.Create()
		}

		if key == "switch" {
			plugin := &SwitchPlugin{Context: c, Input: input}
			return plugin.Create()
		}

		break
	}

	return createSingleResult(data)
}

func (r *routerFactory) setLatency(c *gin.Context, data *RouterSettings) {
	time.Sleep(time.Duration(data.Latency) * time.Millisecond)
}

func (r *routerFactory) setHeaders(c *gin.Context, data *RouterSettings) {
	for key, val := range data.Headers {
		c.Header(key, val)
	}
}
