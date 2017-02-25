package cli

import (
	"time"

	"github.com/labstack/echo"
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
	MockDir     string
}

type routerFactory struct {
	server *echo.Echo
}

func NewRouterFactory(s *echo.Echo) RouterFactory {
	return &routerFactory{server: s}
}

func (r *routerFactory) CreateGET(data *RouterSettings) {
	r.server.GET(data.Uri, func(c echo.Context) error {
		r.setLatency(c, data)
		r.setHeaders(c, data)
		r.execute(c, data)
		return nil
	})
}

func (r *routerFactory) CreatePOST(data *RouterSettings) {
	r.server.POST(data.Uri, func(c echo.Context) error {
		r.setLatency(c, data)
		r.setHeaders(c, data)
		r.execute(c, data)
		return nil
	})
}

func (r *routerFactory) CreatePUT(data *RouterSettings) {
	r.server.PUT(data.Uri, func(c echo.Context) error {
		r.setLatency(c, data)
		r.setHeaders(c, data)
		r.execute(c, data)
		return nil
	})
}

func (r *routerFactory) CreateDELETE(data *RouterSettings) {
	r.server.DELETE(data.Uri, func(c echo.Context) error {
		r.setLatency(c, data)
		r.setHeaders(c, data)
		r.execute(c, data)
		return nil
	})
}

func (r *routerFactory) execute(c echo.Context, data *RouterSettings) {
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

func (r *routerFactory) createDynamic(c echo.Context, data *RouterSettings) (int, interface{}) {
	for key, input := range data.Dynamic {
		if key == "random" {
			plugin := &RandomPlugin{
				MockDir: data.MockDir,
			}

			mapstructure.Decode(input, plugin)
			return plugin.Create()
		}

		if key == "switch" {
			plugin := &SwitchPlugin{
				Context: c,
				Input:   input,
				MockDir: data.MockDir,
			}

			return plugin.Create()
		}

		break
	}

	return createSingleResult(data)
}

func (r *routerFactory) setLatency(c echo.Context, data *RouterSettings) {
	time.Sleep(time.Duration(data.Latency) * time.Millisecond)
}

func (r *routerFactory) setHeaders(c echo.Context, data *RouterSettings) {
	for key, val := range data.Headers {
		c.Request().Header.Set(key, val)
	}
}
