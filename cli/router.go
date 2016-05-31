package cli

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type RouterFactory interface {
	CreateGET(settings *RouterSettings)
}

type RouterSettings struct {
	Uri         string
	ContentType string
	Status      int
	Body        string
	BodyFile    string
	Latency     int
}

type routerFactory struct {
	server *gin.Engine
}

func NewRouterFactory(s *gin.Engine) RouterFactory {
	return &routerFactory{server: s}
}

func (r *routerFactory) CreateGET(data *RouterSettings) {

	r.server.GET(data.Uri, func(c *gin.Context) {
		time.Sleep(time.Duration(data.Latency) * time.Millisecond)

		if data.ContentType == "application/json" {
			var result interface{}

			if len(data.BodyFile) > 0 {
				result = r.parseFile(data.BodyFile)
			} else {
				result = data.Body
			}

			c.JSON(data.Status, result)
			return
		}

		c.String(data.Status, data.Body)
	})

}

func (r *routerFactory) parseFile(f string) interface{} {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return ""
	}

	var d interface{}
	err = json.Unmarshal(content, &d)
	if err != nil {
		log.Println(err)
		return ""
	}

	return d
}
