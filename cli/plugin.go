package cli

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type RandomPlugin struct {
	Status   []int
	Body     []string
	BodyFile []string
}

type SwitchPlugin struct {
	Input   interface{}
	Context *gin.Context
}

type Action struct {
	If   string
	When *Response
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (r *RandomPlugin) Create() (int, interface{}) {
	status := r.Status[rand.Intn(len(r.Status))]
	body := ""
	bodyFile := ""

	if len(r.Body) > 0 {
		body = r.Body[rand.Intn(len(r.Body))]
	} else if len(r.BodyFile) > 0 {
		bodyFile = r.BodyFile[rand.Intn(len(r.BodyFile))]
	}

	return createSingleResult(&RouterSettings{
		Status:   status,
		Body:     body,
		BodyFile: bodyFile,
	})
}

func (s *SwitchPlugin) Create() (int, interface{}) {
	mapper := make(map[string][]Action, 0)
	params := make(map[string]string, 0)

	for key, value := range s.Input.(map[string]interface{}) {
		var actions []Action

		for _, data := range value.([]interface{}) {
			var action Action
			mapstructure.Decode(data, &action)
			actions = append(actions, action)
		}

		mapper[key] = actions
	}

	for key := range mapper {
		query := s.Context.Query(key)
		if len(query) > 0 {
			params[key] = query
		}

		param := s.Context.Param(key)
		if len(param) > 0 {
			params[key] = param
		}
	}

	for key, value := range params {
		for _, action := range mapper[key] {
			if action.If == value {
				return createSingleResult(&RouterSettings{
					Status:   action.When.Status,
					Body:     action.When.Body,
					BodyFile: action.When.BodyFile,
				})
			}
		}
	}

	return 404, "action not found"
}
