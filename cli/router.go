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
	Uri            string
	ContentType    string
	Status         int
	Body           string
	BodyFile       string
	Latency        int
	Dynamic        map[string]interface{}
	Headers        map[string]string
	MockDir        string
	TemplateTokens map[string]string
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
	data.TemplateTokens = r.parseTemplateTokens(c)

	if data.ContentType == "application/json" {
		if len(data.Dynamic) > 0 {
			c.JSON(r.createDynamic(c, data))
			return
		}

		c.JSON(createSingleResult(data))
		return
	}

	if data.ContentType == "image/png" {
		if len(data.Dynamic) > 0 {
			status, content := r.createDynamic(c, data)
			c.Blob(status, "image/png", content.([]byte))
			return
		}
		return
	}

	c.String(data.Status, data.Body)
}

func (r *routerFactory) parseTemplateTokens(c echo.Context) map[string]string {
	tokens := map[string]string{}

	for _, name := range c.ParamNames() {
		tokens["@"+name] = c.Param(name)
	}

	for key, value := range c.QueryParams() {
		tokens["@"+key] = value[0]
	}

	return tokens
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

		if key == "qrcode" {
			plugin := &QrCodePlugin{
				Context:     c,
				MockDir:     data.MockDir,
				ContentType: data.ContentType,
			}
			mapstructure.Decode(input, plugin)
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
